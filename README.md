# 1Panel SDK for Go

A complete Go SDK for the [1Panel](https://github.com/1Panel-dev/1Panel) server
management panel. Covers **every public endpoint** exposed by the 1Panel
frontend (master panel + node agent): authentication, dashboard, host & node
management, firewall, SSH, monitor, containers, images, networks, volumes,
compose, the app store, websites, SSL/ACME/DNS, MySQL/PostgreSQL/MongoDB/Redis,
backups, scheduled tasks, file manager, settings, snapshot, alerts, AI tools,
agent management, and the toolbox.

The SDK is auto-generated from inspection of the
[1Panel](https://github.com/1Panel-dev/1Panel) repository (Go router layout
under `core/router/`, `agent/router/`, and the TypeScript DTO namespace under
`frontend/src/api/interface/`).

## Features

- **Cookie + CSRF + entrance** header management out of the box.
- **Multi-node** support via `sdk.OnNode("1")` — every sub-service targets
  the chosen node via the `CurrentNode` header.
- **Typed wrappers** for ~580 endpoints; the remaining endpoints (and future
  ones) are reachable through the `Call(ctx, method, path, body, out)` wildcard
  on every sub-service.
- Pure Go stdlib (no external HTTP client dependency).
- Go 1.22+, zero third-party modules.

## Install

```bash
go get github.com/zy84338719/1panel-sdk
```

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/zy84338719/1panel-sdk"
    "github.com/zy84338719/1panel-sdk/client"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()

    // Auto-login on construction.
    sdk, err := onepanel.New(onepanel.Options{
        BaseURL:  "https://1panel.example.com",
        Entrance: "1panel_entrance", // omit if the panel is at the root path
        Username: "admin",
        Password: "secret",
        OnLogin: func(r *client.LoginResult) {
            fmt.Printf("logged in as %s (role=%s)\n", r.Name, r.Role)
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    // Dashboard.
    info, _ := sdk.Dashboard.OSInfo(ctx)
    fmt.Println(info)

    // Containers.
    status, _ := sdk.Container.Status(ctx)
    fmt.Println(status)

    // Multi-node: target node 1.
    node := sdk.OnNode("1")
    containers, _ := node.Container.List(ctx)
    fmt.Println(containers)

    // Cleanup.
    _ = sdk.Auth.Logout(ctx)
}
```

## Manual session (no auto-login)

```go
c, _ := client.New(client.Config{
    BaseURL:  "https://1panel.example.com",
    Entrance: "1panel_entrance",
})
sdk := onepanel.NewFromClient(c)
_, err := sdk.Auth.Login(onepanel.LoginForm{Name: "admin", Password: "secret"})
// ...
```

## Sub-services

Every sub-service exposes typed helpers **and** a `Call()` wildcard:

```go
// Typed
hosts, _ := sdk.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 20})

// Wildcard
raw, _ := sdk.Dashboard.Call(ctx, "GET", "/dashboard/app/launcher", nil, &out)
```

| Service | Endpoint prefix | Description |
| --- | --- | --- |
| `Auth` | `/core/auth/*` | Login, MFA, passkey, OIDC, SAML2, API keys, profile |
| `Dashboard` | `/dashboard/*` | OS, CPU, memory, network, top processes |
| `Host` | `/hosts/*` | Host CRUD, disks, supervisor, terminal ws |
| `Container` | `/containers/*` | Container, image, network, volume, compose, repo, docker daemon |
| `App` | `/apps/*` | App store + installed apps |
| `Website` | `/websites/*` | Sites, domains, aliases, certificates |
| `WebsiteSSL` | `/websites/ssl/*` | SSL resource management |
| `WebsiteCA` | `/websites/ca/*` | Private CA |
| `WebsiteDNS` | `/websites/dns/*` | DNS provider accounts |
| `WebsiteAcme` | `/websites/acme/*` | ACME accounts |
| `WebsiteTpl` | `/websites/template/*` | Website templates |
| `Database` | `/databases/*` | MySQL, PostgreSQL, MongoDB, Redis |
| `Backup` | `/backups/*` | Backup destinations, records, restore |
| `Cronjob` | `/cronjobs/*` | Scheduled tasks |
| `File` | `/files/*` | File manager (browse, edit, upload, share, recycle) |
| `Settings` | `/settings/*` | Panel settings, SSL, port, upgrade, AI |
| `Snapshot` | `/settings/snapshot/*` | Panel snapshots |
| `Logs` | `/logs/*` | Login/operation logs, system logs, task list |
| `Groups` | `/groups/*` | Host/app/website groups |
| `Commands` | `/commands/*` | User command library |
| `Script` | `/script/*` | Script library |
| `Toolbox` | `/toolbox/*` | Device, fail2ban, FTP, ClamAV |
| `Alerts` | `/alert/*` | Alert rules, channels, logs |
| `AI` | `/ai/*` | Ollama, GPU, MCP, TensorRT-LLM, agent accounts |
| `Agents` | `/agents/*` | Managed agent instances |
| `SSH` | `/hosts/ssh/*` | SSH service config + root CA cert |
| `Monitor` | `/hosts/monitor/*` | CPU/memory/disk/net monitor |
| `Firewall` | `/hosts/firewall/*` | iptables / firewalld rules |
| `Nginx` | `/nginx/*` | Nginx config (master + server blocks) |
| `Process` | `/processes/*` | Process listing + kill |
| `Runtime` | `/hosts/diagnostics/*` | Runtime diagnostics (goroutines, profile) |
| `Favorite` | `/favorites/*` | User favorites |
| `Task` | `/tasks/*` | Async task progress |

## Endpoints not wrapped

Every typed method is a thin convenience around `s.Call(ctx, method, path, body, out)`.
When 1Panel adds a new endpoint that the SDK does not yet wrap, call it through
`Call`. The signature mirrors `*http.Client.Do`:

```go
err := sdk.SomeService.Call(ctx, "POST", "/the/new/endpoint", body, &out)
```

## MFA

If the panel returns `code != 0` after `Login()`, retry with the MFA code:

```go
login, err := sdk.Auth.Login(onepanel.LoginForm{Name: "admin", Password: "secret"})
if err != nil {
    if apiErr, ok := err.(*client.Error); ok && apiErr.APIError != nil {
        // The body has an "mfaSession" field; ask the user for a code, then:
        login, err = sdk.Auth.LoginByMFA(apiErr.APIError.Message, "123456")
    }
}
```

> Note: 1Panel encodes the MFA session id in the error message — read the
> exact format from your installed version. Many builds instead place the
> session id in `data` of the failed response.

## Custom HTTP client

Bring your own `*http.Client` (e.g. with custom CA, proxy, or instrumentation):

```go
hc := &http.Client{Timeout: 30 * time.Second, Transport: ...}
c, _ := client.New(client.Config{BaseURL: "...", HTTPClient: hc})
sdk := onepanel.NewFromClient(c)
```

## WebSocket / streaming endpoints

The SDK does not wrap streaming endpoints. Use the underlying `*http.Client` and
the path constants the SDK exposes:

```go
wsURL := sdk.C.BaseURL() + sdk.Host.LocalTerminalURL()
// dial wsURL with your favourite websocket library
```

## Example

A runnable end-to-end example lives in [`example/main.go`](./example/main.go).
It logs in, hits every major sub-service, and prints the responses.

```bash
export ONEPANEL_URL=https://1panel.example.com
export ONEPANEL_ENTRANCE=1panel_entrance
export ONEPANEL_USER=admin
export ONEPANEL_PASS=secret
go run ./example
```

## Compatibility

The SDK targets the public 1Panel v2 API (`/core/*` and `/*` resources). It
tracks 1Panel upstream and is regenerated when 1Panel ships new endpoints.
Enterprise / xpack-only endpoints are exposed through the same `Call()` wildcard
and do not require compile-time knowledge.

## License

MIT.
