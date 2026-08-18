#!/usr/bin/env python3
"""Split services2.go into per-domain files for better maintainability."""
import re
from pathlib import Path

SRC = Path('services2.go')
BACKUP = Path('/tmp/services2.go.bak')

text = SRC.read_text()
BACKUP.write_text(text)

# Service name → output file mapping.
groups = {
    'backup_cronjob.go': ['Backup', 'Cronjob'],
    'file.go': ['File'],
    'settings_logs.go': ['Settings', 'Logs'],
    'groups_commands_script.go': ['Groups', 'Commands', 'Script'],
    'toolbox_alerts.go': ['Toolbox', 'Alerts'],
    'ai_agents.go': ['AI', 'Agents'],
    'host_extra.go': ['SSH', 'Monitor', 'Firewall'],
    'nginx_process_runtime.go': ['Nginx', 'Process', 'Runtime'],
    'snapshot_favorite_task.go': ['Snapshot', 'Favorite', 'Task'],
}

# Find each service's start (the "type XService struct" line) and end (next
# "type YService struct" line, or end of file).
service_re = re.compile(r'^type (\w+)Service struct', re.MULTILINE)
matches = list(service_re.finditer(text))
service_blocks = {}
for i, m in enumerate(matches):
    name = m.group(1)
    start = m.start()
    end = matches[i + 1].start() if i + 1 < len(matches) else len(text)
    service_blocks[name] = text[start:end]

# Verify all services in the file are mapped.
unmapped = [n for n in service_blocks if not any(n in names for names in groups.values())]
if unmapped:
    print(f'WARNING: unmapped services: {unmapped}')

# Find the header (the package decl + imports at the top of services2.go).
header_match = re.search(r'^(//.*\n)*package onepanel\n\nimport \([^)]+\)\n', text, re.MULTILINE)
if not header_match:
    # Simpler: just take the first "type" line as the boundary.
    header = text[:matches[0].start()]
else:
    header = header_match.group(0)

# Write each group file.
for filename, names in groups.items():
    out = Path(filename)
    body = ''.join(service_blocks[n] + '\n' for n in names if n in service_blocks)
    out.write_text(header + body + '\n')
    print(f'wrote {out} ({sum(len(service_blocks[n]) for n in names if n in service_blocks)} bytes)')

# Write a stub services2.go that re-exports everything for back-compat.
# This keeps `import "github.com/zy84338719/1panel-sdk"` callers working
# without forcing them to update.
exports = ''.join(f'// (split into {filename}) was {n}Service\n' for filename, names in groups.items() for n in names if n in service_blocks)
SRC.write_text(f'''// Package onepanel - this file used to hold the 21 services listed below,
// but the file was split into per-domain files to make maintenance easier.
// All the type definitions remain in this package, so callers do not need
// to change their imports.
//
//   - BackupService  -> backup_cronjob.go
//   - CronjobService -> backup_cronjob.go
//   - FileService    -> file.go
//   - SettingsService, LogsService          -> settings_logs.go
//   - GroupsService, CommandsService, ScriptService -> groups_commands_script.go
//   - ToolboxService, AlertsService         -> toolbox_alerts.go
//   - AIService, AgentsService              -> ai_agents.go
//   - SSHService, MonitorService, FirewallService -> host_extra.go
//   - NginxService, ProcessService, RuntimeService -> nginx_process_runtime.go
//   - SnapshotService, FavoriteService, TaskService -> snapshot_favorite_task.go
//
// The split preserves all exported names. Re-running the codegen
// (scripts/gen-from-swagger.py) will re-emit them into the new files.
package onepanel
''')

# Also keep the RefreshToken type alias.
SRC.write_text(SRC.read_text() + '''
// RefreshToken is exposed for the rare case where the caller wants to refresh a
// destination's OAuth token without touching the BackupService helpers.
// Kept here for back-compat with the pre-split file.
type RefreshToken struct {
	ServiceBase
}

// SSLService is a thin alias for WebsiteSSLService kept for naming symmetry
// with the rest of the SDK. It exposes the same /websites/ssl/* endpoints.
type SSLService = WebsiteSSLService
''')

print(f'wrote stub {SRC} ({len(SRC.read_text())} bytes)')
print(f'backup at {BACKUP}')
