# File split history

This repo has been split several times to keep individual Go files small
and focused. The two commit-driven splits (June 2026) replaced three big
files with dozens of small per-domain files:

| Original file | Size | Replaced by |
| --- | --- | --- |
| `services.go` (1st split) | 43 KB, 7 service types (App / Website + 5 sub / Database) | `app.go`, `website.go`, `database_mysql.go`, `database_postgres.go`, `database_mongodb.go`, `database_redis.go`, `database_common.go` |
| `services2.go` (1st split) | 82 KB, 21 service types (Backup / Cronjob / File / Settings / Logs / Groups / Commands / Script / Toolbox / Alerts / AI / Agents / SSH / Monitor / Firewall / Nginx / Process / Runtime / Snapshot / Favorite / Task) | `backup.go`, `cronjob.go`, `file.go`, `settings.go`, `logs.go`, `groups.go`, `commands.go`, `script.go`, `toolbox.go`, `alerts.go`, `ai.go`, `agents.go`, `host_ssh.go`, `host_monitor.go`, `host_firewall.go`, `nginx.go`, `process.go`, `runtime.go`, `snapshot.go`, `favorite.go`, `task.go` |
| `container.go` (2nd split) | 16 KB, 1 service 74 methods (ContainerService) | `container_crud.go`, `container_files.go`, `container_image.go`, `container_repo.go`, `container_compose.go`, `container_compose_template.go`, `container_network.go`, `container_volume.go`, `container_daemon.go` (split at `// === Section ===` markers) |
| `database.go` (2nd split) | 15 KB, 1 service 64 methods (DatabaseService) | same as above under `database_*.go` |
| Multi-type mixes (2nd split) | `ai_agents.go` (6.7 KB), `host_extra.go` (8.1 KB), `nginx_process_runtime.go` (2.9 KB), `snapshot_favorite_task.go` (4.3 KB), `toolbox_alerts.go` (11.3 KB), `settings_logs.go` (10.5 KB), `backup_cronjob.go` (8.3 KB), `groups_commands_script.go` (4.4 KB) | per-type files (one file per service) |
| `services_swagger.go` (2nd split) | 179 KB, 23 generated service types, 764 methods | `zgen_ai.go`, `zgen_app.go`, `zgen_alerts.go`, `zgen_backup.go`, `zgen_container.go`, `zgen_core.go`, `zgen_cronjob.go`, `zgen_dashboard.go`, `zgen_database.go`, `zgen_file.go`, `zgen_groups.go`, `zgen_host_extra.go`, `zgen_logs.go`, `zgen_process.go`, `zgen_runtime.go`, `zgen_settings.go`, `zgen_toolbox.go`, `zgen_website.go` |

The one-time split scripts live in `scripts/` for reproducibility:

- `scripts/split-services.py` / `split-services2.py` — first-pass splits.
- `scripts/split-by-section.py` — splits a file at `// === Section ===`
  markers (used for `container.go` and `database.go`).
- `scripts/split-types.py` — splits a multi-type file by receiver type
  (used for `ai_agents.go`, `host_extra.go`, etc.).
- `scripts/use-helpers.py` — rewrites the 4-line `var out; if err := ...`
  boilerplate to a one-line `s.getMap` / `s.postMap` call.

After the second pass all three of the original stub files
(`services.go`, `services2.go`, `services_swagger.go`) were removed —
this file is the single place that records the history.

## Why split?

A single 80 KB file with 21 service types is hard to navigate and easy
to break (a 10-line edit becomes a 1000-line diff). Per-domain files
give the IDE something small to load, the linter something small to
parse, and the reviewer something small to read.

## What about the codegen?

`make codegen` re-emits the `zgen_*.go` files from the latest 1Panel
`swagger.json`. The hand-written service files (`auth.go`, `host.go`,
`container_crud.go`, etc.) are not touched by codegen — they are the
source of truth for typed wrappers; new endpoints that need typed
wrappers go through the regular Go file edit cycle.
