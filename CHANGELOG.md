# Changelog

All notable changes to this project are documented in this file. The format
follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Changed
- `client.PageInfo` removed; the canonical pagination body is `onepanel.PageInfo`
  at the top-level package. Imports do not change.

### Added
- Per-domain service files: `app.go`, `website.go`, `database.go`,
  `backup_cronjob.go`, `file.go`, `settings_logs.go`,
  `groups_commands_script.go`, `toolbox_alerts.go`, `ai_agents.go`,
  `host_extra.go`, `nginx_process_runtime.go`, `snapshot_favorite_task.go`.
  `services.go` and `services2.go` are now 1 KB stubs that document the split.

## [0.1.0] — 2026-08-18

### Added
- Initial release: 731 unique 1Panel v2 API endpoints covered (100% of the
  public `core/cmd/server/docs/swagger.json`).
- 48 sub-services organised by resource: auth, dashboard, host, container,
  app, website (+ SSL / CA / DNS / ACME / template sub-resources), database,
  backup, cronjob, file, settings, logs, groups, commands, script, toolbox,
  alerts, AI, agent, SSH, monitor, firewall, nginx, process, runtime,
  snapshot, favorite, task, health, openresty, plus seven `core-*` helpers
  for the master-panel API.
- Cookie jar + CSRF token auto-sync + entrance-code base64 header.
- Per-node routing via `sdk.OnNode("1")` — every sub-service targets the
  chosen node via the `CurrentNode` header.
- Auto-prepend of `/api/v2` to every request path so callers use the
  same form as the 1Panel frontend.
- Codegen-driven `services_swagger.go` produced by
  `scripts/gen-from-swagger.py`; re-run to absorb upstream API additions.
- `Makefile` (build / test / test-race / cover / lint / fmt / vet / codegen
  / tidy / clean / all) and `.golangci.yml` (v2) — `make lint` is 0 issues.
- Smoke tests in `sdk_test.go` (8 tests) and `client/client_test.go`
  (16 tests) using `httptest`, no real panel required.
- `example/main.go` — runnable end-to-end demo.
