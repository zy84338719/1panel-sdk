#!/usr/bin/env python3
"""Split a multi-type file into per-type files.

Each output file gets a type def + Call() method + the methods that
belong to that type.
"""
import re
import sys
from pathlib import Path


def split_types(src: Path, type_to_file: dict[str, str]):
    text = src.read_text()
    # Header: package + imports (everything before the first `type`).
    m = re.search(r'^(package onepanel\n+(import [^\n]+\n+(\nimport \([^)]+\)\n+)?)?)',
                  text, re.MULTILINE)
    if not m:
        print(f'no header found in {src}', file=sys.stderr)
        sys.exit(1)
    header = m.group(1)
    body = text[m.end():]

    # Find every `type X struct` block and the methods that follow.
    type_re = re.compile(r'^// (\w+Service) covers.*?\ntype \1 struct \{[^{}]*\}\n',
                         re.MULTILINE | re.DOTALL)
    # Find each type declaration position.
    type_matches = list(re.finditer(r'^type (\w+Service) struct', body, re.MULTILINE))
    # For each type, find its methods.
    type_blocks = {}
    for i, tm in enumerate(type_matches):
        name = tm.group(1)
        start = tm.start()
        end = type_matches[i + 1].start() if i + 1 < len(type_matches) else len(body)
        block = body[start:end]
        type_blocks[name] = block

    # Write each type to its target file. Each file gets the header.
    for type_name, fname in type_to_file.items():
        if type_name not in type_blocks:
            print(f'WARN: type {type_name} not found in {src}', file=sys.stderr)
            continue
        content = header + '\n' + type_blocks[type_name]
        Path(fname).write_text(content)
        print(f'wrote {fname}')


if __name__ == '__main__':
    target = sys.argv[1] if len(sys.argv) > 1 else None
    if target == 'ai_agents':
        split_types(
            Path('ai_agents.go'),
            {
                'AIService': 'ai.go',
                'AgentsService': 'agents.go',
            },
        )
    elif target == 'host_extra':
        split_types(
            Path('host_extra.go'),
            {
                'SSHService': 'host_ssh.go',
                'MonitorService': 'host_monitor.go',
                'FirewallService': 'host_firewall.go',
            },
        )
    elif target == 'nginx_process_runtime':
        split_types(
            Path('nginx_process_runtime.go'),
            {
                'NginxService': 'nginx.go',
                'ProcessService': 'process.go',
                'RuntimeService': 'runtime.go',
            },
        )
    elif target == 'snapshot_favorite_task':
        split_types(
            Path('snapshot_favorite_task.go'),
            {
                'SnapshotService': 'snapshot.go',
                'FavoriteService': 'favorite.go',
                'TaskService': 'task.go',
            },
        )
    elif target == 'toolbox_alerts':
        split_types(
            Path('toolbox_alerts.go'),
            {
                'ToolboxService': 'toolbox.go',
                'AlertsService': 'alerts.go',
            },
        )
    elif target == 'settings_logs':
        split_types(
            Path('settings_logs.go'),
            {
                'SettingsService': 'settings.go',
                'LogsService': 'logs.go',
            },
        )
    elif target == 'backup_cronjob':
        split_types(
            Path('backup_cronjob.go'),
            {
                'BackupService': 'backup.go',
                'CronjobService': 'cronjob.go',
            },
        )
    elif target == 'groups_commands_script':
        split_types(
            Path('groups_commands_script.go'),
            {
                'GroupsService': 'groups.go',
                'CommandsService': 'commands.go',
                'ScriptService': 'script.go',
            },
        )
    else:
        print('usage: split-types.py <ai_agents|host_extra|nginx_process_runtime|...>')
        sys.exit(1)
    Path(target + '.go').unlink()
    print(f'deleted {target}.go')
