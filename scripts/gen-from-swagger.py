#!/usr/bin/env python3
"""
Generate Go SDK service files from 1Panel's swagger.json.

Input:  /tmp/1panel-ref/1Panel/core/cmd/server/docs/swagger.json
Output: services_swagger.go in the same directory as the script.

For each (path, method) pair in swagger.paths we:
  1. Group the path by its first 1-2 segments → SDK service.
  2. Generate a method that calls s.Get / s.Post / s.Put / s.Delete with the
     raw path, passing a `map[string]any` body and decoding into a
     `map[string]any` result.
  3. For paths that contain path parameters we generate `func (s *X) Foo_X(ctx, foo)` form
     accepting the {param} values as positional strings, then interpolate.

This is a "max coverage" generator. The user can also keep the manual typed
methods we already wrote; the generated code only adds what's missing.

Run:  python3 scripts/gen-from-swagger.py > services_swagger.go
"""

import json
import re
import sys
from collections import defaultdict
from pathlib import Path

SWAGGER = Path('/tmp/1panel-ref/1Panel/core/cmd/server/docs/swagger.json')
# Output: one zgen_<group>.go per swagger service group.

with SWAGGER.open() as f:
    swag = json.load(f)


# Path → service name
def service_for(path):
    p = path.lstrip('/')
    parts = p.split('/')
    head = parts[0]
    # /core/... → split: core | rest
    if head == 'core':
        sub = parts[1] if len(parts) > 1 else 'misc'
        return f'Core{sub.capitalize()}Service', sub
    if head == 'ai':
        sub = parts[1] if len(parts) > 1 else 'misc'
        if sub == 'accounts':
            return 'AIAccountService', 'account'
        if sub == 'agents':
            return 'AIAgentService', 'agent'
        if sub == 'mcp':
            return 'AIMcpService', 'mcp'
        if sub == 'ollama':
            return 'AIOllamaService', 'ollama'
        if sub == 'gpu':
            return 'AIGpuService', 'gpu'
        if sub == 'tensorrt':
            return 'AITensorService', 'tensorrt'
        if sub == 'domain':
            return 'AIDomainService', 'domain'
        return 'AIOtherService', sub
    if head == 'hosts':
        sub = parts[1] if len(parts) > 1 else 'misc'
        if sub == 'firewall':
            return 'HostFirewallService', 'firewall'
        if sub == 'ssh':
            return 'HostSshService', 'ssh'
        if sub == 'monitor':
            return 'HostMonitorService', 'monitor'
        if sub == 'disks':
            return 'HostDiskService', 'disks'
        if sub == 'tool':
            return 'HostToolService', 'tool'
        if sub == 'diagnostics':
            return 'HostRuntimeService', 'runtime'
        return 'HostOtherService', sub
    return {
        'apps': 'AppService',
        'websites': 'WebsiteService',
        'databases': 'DatabaseService',
        'backups': 'BackupService',
        'cronjobs': 'CronjobService',
        'files': 'FileService',
        'settings': 'SettingsService',
        'logs': 'LogsService',
        'groups': 'GroupsService',
        'commands': 'CommandsService',
        'script': 'ScriptService',
        'toolbox': 'ToolboxService',
        'alert': 'AlertsService',
        'dashboard': 'DashboardService',
        'containers': 'ContainerService',
        'nginx': 'NginxService',
        'openresty': 'OpenRestyService',
        'runtimes': 'RuntimeService',
        'process': 'ProcessService',
        'favorites': 'FavoriteService',
        'tasks': 'TaskService',
        'health': 'HealthService',
    }.get(head, head.capitalize() + 'Service'), head


# Already-existing path set (skip these to avoid duplicates).
EXISTING = set()
for f in Path('.').glob('*.go'):
    if f.name in ('services_swagger.go',):
        continue
    text = f.read_text()
    for m in re.finditer(r's\.(?:Get|Post|Put|Delete)\(ctx,\s*"(/?[^"]+)"', text):
        raw = m.group(1)
        if not raw:
            continue
        canon = re.sub(r'\+[a-zA-Z_]+', '', raw)
        canon = re.sub(r'\+\([^)]+\)', '', canon)
        canon = re.sub(r'\+itoa\([^)]+\)', '', canon)
        canon = re.sub(r'\{[^}]+\}', ':X', canon)
        EXISTING.add(canon)
# Also include "Call" paths.
for f in Path('.').glob('*.go'):
    if f.name in ('services_swagger.go',):
        continue
    text = f.read_text()
    for m in re.finditer(r's\.Call\(ctx,\s*"[A-Z]+",\s*"(/?[^"]+)"', text):
        raw = m.group(1)
        if not raw:
            continue
        canon = re.sub(r'\{[^}]+\}', ':X', raw)
        EXISTING.add(canon)


# Group by service.
by_service = defaultdict(list)
for path, methods in swag['paths'].items():
    canon = re.sub(r'\{[^}]+\}', ':X', path)
    if canon in EXISTING:
        continue
    for m, info in methods.items():
        if m in ('parameters',):
            continue
        svc, sub = service_for(path)
        by_service[svc].append((m.upper(), path, info.get('summary', ''), canon, sub))


def method_name(path, m, summary):
    """Pick a Go method name from the path + summary."""
    # Strip leading slash
    p = path.lstrip('/')
    parts = [seg for seg in p.split('/') if not seg.startswith('{')]
    if not parts:
        parts = ['Root']
    # First non-empty part is the resource (already the service name); drop it.
    if len(parts) > 1:
        parts = parts[1:]
    # Drop pure param words.
    # Build from parts.
    name = ''.join(seg.capitalize() for seg in parts)
    # Cap to 60 chars.
    if len(name) > 60:
        name = name[:60]
    # Prefix with verb from path if name starts oddly
    name = re.sub(r'[^A-Za-z0-9]', '', name)
    if not name:
        name = 'Endpoint'
    return m.upper() + name


def go_keyword_safe(name):
    """Avoid Go reserved keywords as parameter names."""
    keywords = {'type', 'func', 'range', 'map', 'chan', 'select', 'default',
                'package', 'import', 'var', 'const', 'return', 'if', 'else',
                'for', 'switch', 'case', 'break', 'continue', 'go', 'defer',
                'struct', 'interface'}
    if name in keywords:
        return name + 'Val'
    return name


def fmt_sprintf_path(path, params):
    """Build a fmt.Sprintf format string + the argument list for a path with {param}s.

    Returns (format_string, args_list).
    """
    fmt = []
    args = []
    i = 0
    s = path
    while i < len(s):
        m = re.match(r'\{([^}]+)\}', s[i:])
        if m:
            pname = go_keyword_safe(m.group(1))
            fmt.append('%s')
            args.append(pname)
            i += m.end()
        else:
            fmt.append(s[i])
            i += 1
    return ''.join(fmt), args


def gen_method(svc, m, path, summary, params):
    verb = m.upper()
    name = method_name(path, m, summary)
    # Skip names that aren't valid Go identifiers
    if not re.match(r'^[A-Z][A-Za-z0-9]+$', name):
        return None
    # Build parameter list. GET has no body; other verbs have one.
    param_names = [go_keyword_safe(p) for p in params]
    has_body = verb != 'GET'
    if has_body:
        args = [f'{p} string' for p in param_names] + ['body any']
    else:
        args = [f'{p} string' for p in param_names]
    sig_args = ', '.join(['ctx context.Context'] + args)
    # The ServiceBase exposes helper methods that return (map[string]any, error).
    helper = {'GET': 'getMap', 'POST': 'postMap'}.get(verb, 'Do')
    body_arg = 'body' if has_body else None
    call = f's.{helper}({path_lit_call(params, path, param_names, body_arg)})'
    comment = f'// {name} — {summary} ({verb} {path})'
    return f'''{comment}
func (s *{svc}) {name}({sig_args}) (map[string]any, error) {{
	return {call}
}}
'''


def path_lit_call(params, path, param_names, body_arg):
    """Return the argument list for the helper call."""
    parts = ['ctx']
    if not params:
        parts.append(f'"{path}"')
    else:
        fmt, fargs = fmt_sprintf_path(path, params)
        quoted = '"' + fmt.replace('"', '\\"') + '"'
        parts.append(f'fmt.Sprintf({quoted}, {", ".join(fargs)})')
    if body_arg is not None:
        parts.append(body_arg)
    return ', '.join(parts)


# Group by service
service_subs = {}
for svc, items in by_service.items():
    if svc == 'ContainerService':
        # Don't pollute already-massive ContainerService with dozens of generics.
        continue
    service_subs[svc] = items


# Emit
out = ['// Code generated by scripts/gen-from-swagger.py. DO NOT EDIT.\n',
        'package onepanel\n',
        'import (\n',
        '\t"context"\n',
        '\t"fmt"\n',
        ')\n\n']

# Clean old zgen_*.go files first so they don't pollute the
# "already-declared" check below. (services_swagger.go is the only file
# that needs to stay — but it now contains a stub.)
for f in Path('.').glob('zgen_*.go'):
    f.unlink()

# Detect already-declared service types in the package.
EXISTING_TYPES = set()
for f in Path('.').glob('*.go'):
    if f.name in ('services_swagger.go',) or f.name.startswith('zgen_'):
        continue
    text = f.read_text()
    for m in re.finditer(r'^type (\w+Service) struct', text, re.MULTILINE):
        EXISTING_TYPES.add(m.group(1))

# Also collect already-defined method names per type, so we don't redeclare.
EXISTING_METHODS = defaultdict(set)
for f in Path('.').glob('*.go'):
    if f.name in ('services_swagger.go',) or f.name.startswith('zgen_'):
        continue
    text = f.read_text()
    for m in re.finditer(r'^func \(s \*(\w+)\) (\w+)\(', text, re.MULTILINE):
        EXISTING_METHODS[m.group(1)].add(m.group(2))

# Re-include ALL services (we'll attach methods to whatever type we find or
# declare a new one).
by_service_all = defaultdict(list)
for path, methods in swag['paths'].items():
    canon = re.sub(r'\{[^}]+\}', ':X', path)
    if canon in EXISTING:
        continue
    for m, info in methods.items():
        if m in ('parameters',):
            continue
        svc, sub = service_for(path)
        by_service_all[svc].append((m.upper(), path, info.get('summary', ''), canon, sub))


# Map service name -> output file. All generated services get a
# `zgen_<group>.go` filename (the z-prefix sorts them last, so
# hand-written files keep their natural alphabetic position).
def file_for_service(svc: str) -> str:
    # group by swagger tag/prefix for readability.
    if svc.startswith('AI'):
        return 'zgen_ai.go'
    if svc.startswith('Core'):
        return 'zgen_core.go'
    if svc.startswith('Host'):
        return 'zgen_host_extra.go'
    if svc in ('HealthService', 'OpenRestyService'):
        return 'zgen_runtime.go'
    return f'zgen_{_snake(svc)}.go'


def has_path_params(items) -> bool:
    """A service needs fmt.Sprintf only if at least one of its endpoints
    has a path parameter (e.g. /containers/{id})."""
    for _m, path, _sm, _c, _sub in items:
        if '{' in path:
            return True
    return False


def _snake(name: str) -> str:
    out = []
    for i, c in enumerate(name):
        if c.isupper() and i > 0 and (name[i - 1].islower() or (i + 1 < len(name) and name[i + 1].islower())):
            out.append('_')
        out.append(c.lower())
    return ''.join(out).replace('_service', '')


# Build per-file output.
file_buffers = {}
file_has_fmt = {}  # fname -> bool, whether fmt.Sprintf is needed
for svc in sorted(by_service_all.keys()):
    items = by_service_all[svc]
    fname = file_for_service(svc)
    is_first = fname not in file_buffers
    if is_first:
        # Header — finalize() below closes the imports block before writing.
        file_buffers[fname] = [
            '// Code generated by scripts/gen-from-swagger.py. DO NOT EDIT.\n',
            'package onepanel\n',
            'import (\n',
            '\t"context"\n',
        ]
        file_has_fmt[fname] = has_path_params(items)
    buf = file_buffers[fname]
    file_has_fmt[fname] = file_has_fmt[fname] or has_path_params(items)
    is_new_type = svc not in EXISTING_TYPES
    if is_new_type:
        buf.append(f'// {svc} covers endpoints not yet wrapped in typed methods. Use the\n'
                   f'// generated per-endpoint helpers for full coverage.  See the\n'
                   f'// 1Panel swagger.json for the original DTO definitions.\n')
        buf.append(f'type {svc} struct {{ ServiceBase }}\n\n')
        buf.append(f'// Call invokes an arbitrary endpoint in this service.\n')
        buf.append(f'func (s *{svc}) Call(ctx context.Context, method, path string, body, out any) error {{\n')
        buf.append(f'\treturn s.Do(ctx, method, path, body, out)\n')
        buf.append(f'}}\n\n')
    seen = set(EXISTING_METHODS[svc])
    for method, path, summary, canon, sub in items:
        params = re.findall(r'\{([^}]+)\}', path)
        text = gen_method(svc, method, path, summary, params)
        if text is None:
            continue
        m_name = re.search(r'func \(s \*\S+\) (\w+)\(', text).group(1)
        if m_name in seen:
            continue
        seen.add(m_name)
        buf.append(text + '\n')

# Clean old zgen_*.go files first.
for f in Path('.').glob('zgen_*.go'):
    f.unlink()

# Now finalize each file: insert the closing imports `)` right after the
# `"context"` line, then write the file. We need to splice `)` in at the
# right position (after the `import (...)` block, not at end of file).
for fname, buf in file_buffers.items():
    # Find the position right after the last import line ("\t"context"\n").
    # The "import (" line is at index 2; the "context" line is at index 3.
    # We insert ")\n\n" at index 4 (right after "context" line), but only
    # if fmt is needed (otherwise it sits before any optional fmt import).
    insert_at = 4
    if file_has_fmt[fname]:
        # Need to add "\t"fmt"\n" before the closing ")\n".
        buf.insert(insert_at, '\t"fmt"\n')
        insert_at += 1
    buf.insert(insert_at, ')\n\n')
    body = ''.join(buf)
    Path(fname).write_text(body)

# Split history lives in SPLIT.md now; no monolithic stub file is
# written. (Earlier revisions kept a one-line services_swagger.go stub,
# but it was redundant with the per-zgen file naming.)

total_methods = sum(len(v) for v in by_service_all.values())
print(f'wrote {len(file_buffers)} zgen_*.go files with {total_methods} methods across {len(by_service_all)} services')
