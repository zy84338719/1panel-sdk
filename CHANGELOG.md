# Changelog

All notable changes to this project are documented in this file. The format
follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added
- **API key auth** (header-based, the official machine-to-machine flow):
  `client.Config.APIKey` + `APISignMethod` (`hmac-sha256` default, `md5`
  legacy). When set, every request carries `1Panel-Timestamp` and
  `1Panel-Token` headers, and `X-CSRF-Token` is skipped (the two flows
  are independent in 1Panel). Exposed as `onepanel.Options.APIKey` for
  one-shot construction. Matches the spec at
  https://1panel.cn/docs/v2/dev_manual/api_manual/#22-token.
- `ServiceBase.GetList` / `PostList` helpers (plus unexported `getList` /
  `postList` for the codegen) for endpoints whose `data` field is a
  top-level array (e.g. `/dashboard/app/launcher`, `/groups/search`).
- `client.SignToken` exported helper for callers that need to compute the
  same signature themselves (proxies, custom transports, verification).
- `example/verify/` integration program that probes a live 1Panel server
  with API key auth and reports each endpoint's status.

### Changed
- **`CodeSuccess` is now `http.StatusOK` (200)**, matching the real 1Panel
  v2 protocol. The panel puts the HTTP status into the envelope's `code`
  field — `200` for success, `400/401/500` for the usual HTTP errors. The
  old `0` was a wrong guess carried from an early read. Added
  `CodeBadRequest` and `CodeInternalError` as aliases for the new
  standard codes. `CodeAuthFail` (401), `CodeForbidden` (403), and the
  business codes 313/408/410/433/434 are preserved.
- `client.New` now accepts `Config.APIKey` and `Config.APISignMethod`.
  The HTTP client adds the signing headers in `doWithNode` so the
  middleware path is identical for all four verbs.
- Test fixtures updated to return `code: 200` on success.

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
