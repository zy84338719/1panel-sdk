#!/usr/bin/env python3
"""Split services.go (43KB) into per-domain files."""
import re
from pathlib import Path

SRC = Path('services.go')
BACKUP = Path('/tmp/services.go.bak')
text = SRC.read_text()
BACKUP.write_text(text)

# Service name (without "Service" suffix) → output file.
groups = {
    'app.go': ['App'],
    'website.go': ['Website', 'WebsiteSSL', 'WebsiteCA', 'WebsiteDNSAccount', 'WebsiteAcmeAccount', 'WebsiteTemplate'],
    'database.go': ['Database'],
}

# Find each service's start (the "type XService struct" line) and end.
service_re = re.compile(r'^type (\w+Service) struct\s*\{', re.MULTILINE)
matches = list(service_re.finditer(text))
service_blocks = {}
for i, m in enumerate(matches):
    name = m.group(1)
    start = m.start()
    end = matches[i + 1].start() if i + 1 < len(matches) else len(text)
    service_blocks[name] = text[start:end]

# Verify all services in the file are mapped (by name prefix).
unmapped = [n for n in service_blocks if not any(n.startswith(name) for name in [g for names in groups.values() for g in names])]
if unmapped:
    print(f'WARNING: unmapped services: {unmapped}')

# Find the header (package + imports).
header_match = re.search(r'^(//.*\n)*package onepanel\n(\nimport [^\n]+\n(\nimport \([^)]+\))?)?', text, re.MULTILINE)
header = header_match.group(0) if header_match else 'package onepanel\n'

# Write each group file.
for filename, names in groups.items():
    out = Path(filename)
    matching = [n for n in service_blocks if any(n.startswith(prefix) for prefix in names)]
    body = ''.join(service_blocks[n] + '\n' for n in matching)
    out.write_text(header + body + '\n')
    print(f'wrote {out} ({sum(len(service_blocks[n]) for n in matching)} bytes)')

# Write a stub.
SRC.write_text(f'''// Package onepanel - this file used to host the seven service types below,
// but was split into per-domain files for maintainability. The split keeps
// every exported name in the same package, so callers do not need to
// change their imports.
//
//   - AppService                 -> app.go
//   - WebsiteService             -> website.go
//   - WebsiteSSLService          -> website.go
//   - WebsiteCAService           -> website.go
//   - WebsiteDNSAccountService   -> website.go
//   - WebsiteAcmeAccountService  -> website.go
//   - WebsiteTemplateService     -> website.go
//   - DatabaseService            -> database.go
//
// Re-running scripts/gen-from-swagger.py will only re-emit the codegen file
// (services_swagger.go); these hand-written service files are the source of
// truth.
package onepanel
''')
print(f'wrote stub {SRC}')
print(f'backup at {BACKUP}')
