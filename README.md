# go-1panel — Go SDK for 1Panel v2

[![Go Reference](https://pkg.go.dev/badge/github.com/zy84338719/go-1panel.svg)](https://pkg.go.dev/github.com/zy84338719/go-1panel)
[![Test](https://img.shields.io/badge/tests-38%2F38-brightgreen)](./sdk_test.go)
[![Lint](https://img.shields.io/badge/lint-0%20issues-brightgreen)](./.golangci.yml)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8)](./go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue)](./LICENSE)

A complete, pure-stdlib Go SDK for the [1Panel](https://github.com/1Panel-dev/1Panel)
server management panel. Covers **every public endpoint** in the 1Panel
[`swagger.json`](https://github.com/1Panel-dev/1Panel/blob/master/core/cmd/server/docs/swagger.json)
(**731 paths × 663 DTOs**) and is verified end-to-end against a live 1Panel
server with the official API-key auth flow.

## Why

1Panel is a great panel, but talking to it from Go means:

- Maintaining a session cookie jar + CSRF dance.
- Prepending `/api/v2` to every path (the swagger `basePath`).
- Sending an `EntranceCode` base64 header for hidden entrances.
- Switching the default node via the `CurrentNode` header.
- For service accounts, computing `1Panel-Timestamp` + `1Panel-Token` on
  every request (HMAC-SHA256 or MD5).

`go-1panel` does all of that for you, and gives you **48 typed sub-services**
that mirror the panel UI 1:1 — `Auth`, `Host`, `Container`, `App`, `Website`,
`Database`, `Backup`, `Cronjob`, `File`, `Settings`, `Snapshot`, `Logs`,
`Groups`, `Commands`, `Script`, `Toolbox`, `Alerts`, `AI*`, `OpenResty`,
`Monitor`, `Firewall`, `Process`, `Runtime`, `Favorite`, `Task`, plus seven
master-panel `Core*` services for the `/core/*` API family.

## Features

- **All 731 endpoints wrapped** — hand-curated typed methods for the common
  workflows, plus per-endpoint codegen-generated methods
  (`GET…`, `POST…`, `PUT…`, `DELETE…` named after the swagger path) for the
  long tail. Every sub-service also exposes a `Call()` wildcard for endpoints
  1Panel adds between SDK releases.
- **API key auth (header-based)** — set `Options.APIKey` and the SDK signs
  every request with HMAC-SHA256 (recommended) or MD5 (legacy). No login
  needed; the two header flow is fully compatible with the panel's official
  [API 接口 docs](https://1panel.cn/docs/v2/dev_manual/api_manual/#22-token).
- **Cookie + CSRF + entrance** header management when you do use a
  username/password login, including MFA retries.
- **Auto `/api/v2` prefix** — paths you pass look exactly like the swagger.
- **Multi-node** via `sdk.OnNode("1")` — every sub-service targets the
  chosen node via `CurrentNode`.
- **Per-service decode helpers** — `GetMap` / `PostMap` for object
  responses, `GetList` / `PostList` for array responses. Use `Do` with a
  typed struct for full type safety.
- **Pure Go stdlib, zero third-party modules** — `go 1.22+` and nothing
  else. `golangci-lint` reports 0 issues. 38/38 tests pass.
- **Codegen** — when 1Panel adds a new endpoint, re-run
  `scripts/gen-from-swagger.py` and the new method shows up in seconds.

## Install

```bash
go get github.com/zy84338719/go-1panel
```

## Quick start

### API key auth (recommended for service accounts / CI / agents)

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/zy84338719/go-1panel"
    "github.com/zy84338719/go-1panel/client"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    sdk, err := onepanel.New(onepanel.Options{
        BaseURL: "https://1panel.example.com",
        APIKey:  "<your-api-key-from-the-panel>",
        // APISignMethod: client.SignMethodHMACSHA256,  // default; client.SignMethodMD5 for legacy panels
    })
    if err != nil {
        panic(err)
    }

    info, _ := sdk.Dashboard.GETBaseOs(ctx)
    fmt.Printf("%+v\n", info)

    hosts, _ := sdk.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 20})
    fmt.Printf("nodes: %+v\n", hosts)

    // Point at a specific node:
    worker := sdk.OnNode("1")
    containers, _ := worker.Container.Call(ctx, "POST", "/containers/search",
        map[string]any{"page": 1, "pageSize": 10}, nil)
    fmt.Printf("worker containers: %+v\n", containers)
}
```

### Username + password (with optional auto-login)

```go
sdk, err := onepanel.New(onepanel.Options{
    BaseURL:  "https://1panel.example.com",
    Entrance: "1panel_entrance", // empty string if the panel is at the root
    Username: "admin",
    Password: "secret",
    // OnLogin: func(r *client.LoginResult) { persist r.Token ... },
})
```

If MFA is required, the first `Login` returns a 401 with `mfaSession` in
the response — read it from `(*client.Error).APIError.Data` and call
`sdk.Auth.LoginByMFA(sessionID, "123456")`.

## Sub-services

The SDK ships **48 sub-services** (plus `Agent` as a back-compat alias for
`Agents`). Every one exposes a `Call()` wildcard so you're never blocked
when 1Panel ships a new endpoint before the SDK does.

| Service | Endpoint prefix | What it covers |
| --- | --- | --- |
| `Auth` | `/core/auth/*` | login, MFA, passkey, OIDC, SAML2, API keys, profile |
| `Dashboard` | `/dashboard/*` | OS, CPU, memory, network, top processes, app launcher |
| `Host` | `/hosts/*` | host CRUD, disks, supervisor, terminal ws |
| `Container` | `/containers/*` | container, image, network, volume, compose, repo, daemon |
| `App` | `/apps/*` | app store + installed apps |
| `Website` | `/websites/*` | sites, domains, aliases, certificates, nginx conf, PHP |
| `WebsiteSSL` | `/websites/ssl/*` | SSL resource management |
| `WebsiteCA` | `/websites/ca/*` | private CA |
| `WebsiteDNS` | `/websites/dns/*` | DNS provider accounts |
| `WebsiteAcme` | `/websites/acme/*` | ACME accounts |
| `WebsiteTpl` | `/websites/template/*` | website templates |
| `Database` | `/databases/*` | MySQL, PostgreSQL, MongoDB, Redis |
| `Backup` | `/backups/*` | backup destinations, records, restore |
| `Cronjob` | `/cronjobs/*` | scheduled tasks |
| `File` | `/files/*` | file manager (browse, edit, upload, share, recycle, history, favorite) |
| `Settings` | `/settings/*` | panel settings, SSL, port, upgrade, AI, base dir |
| `Snapshot` | `/settings/snapshot/*` | panel snapshots |
| `Logs` | `/logs/*` | login/operation logs, system logs, task list |
| `Groups` | `/groups/*` | host/app/website groups |
| `Commands` | `/commands/*` | user command library |
| `Script` | `/script/*` | script library |
| `Toolbox` | `/toolbox/*` | device, fail2ban, FTP, ClamAV |
| `Alerts` | `/alert/*` | alert rules, channels, logs |
| `AI` | `/ai/*` | Ollama, GPU, MCP, TensorRT-LLM |
| `Agent` / `Agents` | `/ai/agents/*` | managed agent instances |
| `AIAccount` | `/ai/accounts/*` | model accounts (OpenAI-compatible) |
| `AIAgent` | `/ai/agents/*` | agent role / channel / skill config |
| `AIDomain` | `/ai/domain/*` | AI gateway domain bindings |
| `AIMcp` | `/ai/mcp/*` | MCP servers |
| `AITensor` | `/ai/tensorrt/*` | TensorRT-LLM engines |
| `CoreAuth` / `CoreBackup` / `CoreCommand` / `CoreGroup` / `CoreLog` / `CoreScript` / `CoreSetting` | `/core/*` | master-panel API (auth, scripts, command library, groups, logs, settings, backup client) |
| `Health` | `/health/*` | service health check |
| `OpenResty` | `/openresty/*` | OpenResty / Nginx config |
| `Runtime` | `/runtimes/*` | PHP / Node / Go / Python runtime management |
| `SSH` | `/hosts/ssh/*` | SSH service config + root CA cert |
| `Monitor` | `/hosts/monitor/*` | CPU / memory / disk / net monitor |
| `Firewall` | `/hosts/firewall/*` | iptables / firewalld rules |
| `Nginx` | `/nginx/*` | Nginx config (master + server blocks) |
| `Process` | `/process/*` | process listing + kill |
| `Favorite` | `/files/favorite/*` | user favorites |
| `Task` | `/tasks/*` | async task progress |

## API key auth — the wire format

When `Options.APIKey` is set, the SDK adds two headers to every request
(auto-computed, no manual work):

| Header | Value |
| --- | --- |
| `1Panel-Timestamp` | current Unix time, in seconds |
| `1Panel-Token` | hex signature of the timestamp with the API key |

The signature is one of:

- **HMAC-SHA256 (default, recommended)**:
  `hex(hmac_sha256(API_KEY, "1panel:" + timestamp))`
- **MD5 (legacy, removed in future 1Panel releases)**:
  `hex(md5("1panel" + API_KEY + timestamp))`

Cookie/CSRF auth and API-key auth are **mutually exclusive at the wire
level**: the panel's `/core/auth/*` endpoints require session cookies, so
`Auth.CurrentUser` will return `401 用户未登录: 当前会话已过期` when called
with an API key. Use API keys for node-facing work (`/dashboard`, `/hosts`,
`/toolbox`, `/containers`, `/apps`, `/websites`, `/databases`, `/files`,
...) and a real login for `/core/auth/*`.

You can sign the same way externally with `client.SignToken(method, key, ts)`:

```go
sig := client.SignToken(client.SignMethodHMACSHA256, "my-key", "1700000000")
```

## Array responses — `GetList` / `PostList`

A handful of endpoints (`/dashboard/app/launcher`, `/groups/search`,
`/ai/accounts/providers`, `/core/settings/interface`, ...) return their
`data` as a top-level JSON array instead of an object. The default
`GetMap` / `PostMap` helpers can't decode those — use the array helpers
on any service, or call `Do` with a typed slice:

```go
// Built-in helpers (returns []any):
items, err := sdk.Dashboard.GetList(ctx, "/dashboard/app/launcher")
items, err := sdk.Groups.PostList(ctx, "/groups/search", map[string]any{})

// Or decode into a typed slice (recommended once you have the DTO):
var items []dto.AppLauncher
err := sdk.Dashboard.Do(ctx, "GET", "/dashboard/app/launcher", nil, &items)
```

## Custom HTTP client

Bring your own `*http.Client` for custom CA, proxy, or instrumentation:

```go
hc := &http.Client{
    Timeout: 30 * time.Second,
    Transport: ...,
}
c, _ := client.New(client.Config{BaseURL: "...", HTTPClient: hc})
sdk := onepanel.NewFromClient(c)
```

## Architecture

```
┌──────────────────────────────────────────────────────────────┐
│                         user code                            │
└──────────────────┬───────────────────────────────────────────┘
                   │ sdk.Auth / sdk.Host / sdk.Container ...
                   ▼
┌──────────────────────────────────────────────────────────────┐
│              onepanel package — 48 sub-services              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────────┐    │
│  │ Auth     │  │ Host     │  │ Container│  │ Website    │    │
│  │ AI       │  │ Database │  │ Settings │  │ ...52 more │    │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └─────┬──────┘    │
│       │             │             │             │            │
│       └─────────────┴─────┬───────┴─────────────┘            │
│                            │                                │
│                  ┌─────────▼──────────┐                     │
│                  │   ServiceBase     │                     │
│                  │ Get/Post/Put/Del  │                     │
│                  │ Do/Call + helpers │                     │
│                  └─────────┬─────────┘                     │
└────────────────────────────┼────────────────────────────────┘
                             │
                             ▼
                ┌─────────────────────────────┐
                │   client package            │
                │  ┌─────────────────────┐    │
                │  │ cookie jar + CSRF   │    │
                │  │ /api/v2 auto-prefix │    │
                │  │ EntranceCode header │    │
                │  │ API-key sign+ts hdr │    │
                │  └─────────────────────┘    │
                └──────────────┬──────────────┘
                               │ http.Client (stdlib)
                               ▼
                       1Panel panel
                    /<entrance>/api/v2/...
```

1Panel serves two API families on one port:

- `/core/...` — master panel APIs (auth, settings, scripts, command library,
  logs, groups, backup client). Served by the `core` process and reverse-
  proxied behind the public entrance path.
- `/<resource>...` — node-facing APIs (hosts, containers, apps, websites,
  databases, files, ...). Each method is invoked on a particular node
  selected via the `CurrentNode` header.

The SDK normalises both: you pass post-strip paths (`/core/auth/login`,
`/containers/search`, ...) and the client prepends `/api/v2` and the right
`CurrentNode` automatically.

## Repo layout

```
go-1panel/
├── client/                 # low-level HTTP client
│   ├── client.go           #   cookie jar, CSRF sync, /api/v2, EntranceCode
│   ├── sign.go             #   API-key signing (HMAC-SHA256 + MD5)
│   ├── auth_helpers.go     #   Login / Logout / SetCookieForURL
│   ├── types.go            #   shared types (Result, Error, Config, codes)
│   ├── client_test.go
│   └── client_bench_test.go
├── scripts/
│   ├── gen-from-swagger.py   # 1Panel swagger.json -> zgen_*.go
│   ├── split-*.py            # one-time history scripts
│   └── use-helpers.py        # one-time 4-line -> 1-line method rewrite
├── auth.go                 # AuthService (/core/auth/*)
├── dashboard.go            # DashboardService
├── host.go                 # HostService
├── host_ssh.go             # SSHService
├── host_monitor.go         # MonitorService
├── host_firewall.go        # FirewallService
├── container_*.go          # ContainerService — split per concern
├── app.go                  # AppService
├── website.go              # WebsiteService + 5 sub-services
├── database_*.go           # DatabaseService — split per engine
├── backup.go               # BackupService
├── cronjob.go              # CronjobService
├── file.go                 # FileService
├── settings.go             # SettingsService
├── snapshot.go             # SnapshotService
├── logs.go                 # LogsService
├── groups.go               # GroupsService
├── commands.go             # CommandsService
├── script.go               # ScriptService
├── toolbox.go              # ToolboxService
├── alerts.go               # AlertsService
├── ai.go                   # AIService
├── agents.go               # AgentsService
├── nginx.go                # NginxService
├── process.go              # ProcessService
├── runtime.go              # RuntimeService
├── favorite.go             # FavoriteService
├── task.go                 # TaskService
├── zgen_*.go               # GENERATED — one per swagger service group
│                             (z prefix sorts last in `ls`)
├── onepanel.go             # SDK entry point + Options + bind()
├── types.go                # PageInfo, ServiceBase, getMap/getList helpers
├── doc.go                  # package godoc
├── example/                # runnable demo (username + password)
├── example/verify/         # live integration probe (API key)
├── example_test.go         # godoc Example* + httptest fake
├── sdk_test.go             # SDK integration tests
├── client/*_test.go        # client unit tests
├── types_test.go           # getMap / getList / GetList helper tests
├── go.mod, go.sum          # zero third-party deps
├── Makefile                # 14 targets (build, test, lint, codegen, ...)
├── .golangci.yml           # lint v2 config
├── CHANGELOG.md
├── CONTRIBUTING.md
└── LICENSE
```

## Development

```bash
make all           # build + test + lint
make test          # go test -count=1 ./...
make test-race     # go test -race ./...
make cover         # coverage report
make lint          # golangci-lint run ./...
make codegen       # regenerate zgen_*.go from swagger.json
make fmt           # gofmt -s -w .
make vet           # go vet ./...
make tidy          # go mod tidy
```

See [CONTRIBUTING.md](./CONTRIBUTING.md) for the full contribution guide
and [CHANGELOG.md](./CHANGELOG.md) for what's changed.

## Codegen

The SDK is partially generated from the 1Panel `swagger.json` to stay in
sync with upstream. To regenerate the auto-generated methods:

```bash
# Pull the latest swagger
curl -sL https://raw.githubusercontent.com/1Panel-dev/1Panel/master/core/cmd/server/docs/swagger.json \
    > /tmp/swagger.json
# Edit scripts/gen-from-swagger.py to point at /tmp/swagger.json
python3 scripts/gen-from-swagger.py
go build ./...
```

Every endpoint that 1Panel adds or renames in `swagger.json` will become a
new method (named after the HTTP verb + path) on its corresponding service
on the next codegen run.

## Live verification

`example/verify/` is a small integration program that exercises the SDK
against a real 1Panel server with API-key auth. It probes a representative
endpoint from every one of the 48 services (read-only), classifies the
result (OK / server-rejected / decode-issue / 404) and exits non-zero on
any unexpected failure. See [`example/verify/README.md`](./example/verify/README.md).

## Compatibility

- The SDK targets the public 1Panel v2 API (`/core/*` and `/*` resources).
  It is regenerated against the upstream `swagger.json` whenever 1Panel
  ships new endpoints.
- Enterprise / xpack-only endpoints are reached through the same `Call()`
  wildcard and do not require compile-time knowledge.
- Go 1.22+, no third-party modules.

## License

MIT — see [LICENSE](./LICENSE).
