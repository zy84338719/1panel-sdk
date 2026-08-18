// Package onepanel is a complete Go SDK for the 1Panel server management
// panel (https://github.com/1Panel-dev/1Panel).
//
// # Overview
//
// Every public endpoint exposed by the 1Panel v2 frontend has a typed
// wrapper here, plus a generated per-endpoint helper that hits the same
// path. The SDK covers 731 unique swagger paths (100% of the public API).
//
// # Entry point
//
// Use New() to construct an authenticated SDK; the call performs a
// /core/auth/login synchronously and stores the session cookies + CSRF
// token in the underlying client.
//
//	sdk, err := onepanel.New(onepanel.Options{
//	    BaseURL:  "https://1panel.example.com",
//	    Entrance: "1panel_entrance",
//	    Username: "admin",
//	    Password: "secret",
//	})
//	if err != nil { log.Fatal(err) }
//	defer sdk.Auth.Logout(ctx)
//
// # Sub-services
//
// Sub-services live on the SDK struct as fields named after the 1Panel
// resource they wrap: sdk.Host, sdk.Container, sdk.Website, sdk.AI,
// sdk.OpenResty, sdk.Health, ... — see the SDK struct for the full list.
//
// Every sub-service embeds ServiceBase and therefore exposes Get / Post /
// Put / Delete / Do / Call. The Call wildcard is the escape hatch for
// endpoints that don't yet have a typed wrapper:
//
//	out, _ := sdk.SomeService.Call(ctx, "POST", "/the/new/endpoint", body, &out)
//
// # Paths
//
// Pass paths in the same form the 1Panel frontend uses (e.g.
// /core/auth/login, /containers/search, /websites/ssl). The client
// automatically prepends /api/v2 to match the swagger basePath.
//
// # Multi-node
//
// 1Panel fans out to multiple nodes; the SDK selects a node via the
// CurrentNode header. Use sdk.OnNode("1") to get a fresh SDK whose
// every sub-service targets the chosen node:
//
//	node := sdk.OnNode("1")
//	containers, _ := node.Container.List(ctx)
//
// # Errors
//
// All errors returned by typed methods are *client.Error. Inspect the
// 1Panel business code via IsCode() and unwrap with errors.As:
//
//	if apiErr, ok := err.(*client.Error); ok && apiErr.IsCode(401) {
//	    // re-auth
//	}
package onepanel
