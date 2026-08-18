// Verify the SDK against a real 1Panel server using API key auth.
//
// Usage:
//   go run ./example/verify \
//     -url http://192.168.108.235:55555 \
//     -key "0HeE0VPfS2TY5N1ILGgE322McU7hRdmz"
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zy84338719/go-1panel"
	"github.com/zy84338719/go-1panel/client"
)

func main() {
	var (
		baseURL  = flag.String("url", "http://192.168.108.235:55555", "1Panel base URL")
		key      = flag.String("key", "", "API key from the panel")
		signMth  = flag.String("sign", "hmac-sha256", "sign method: hmac-sha256 (default) or md5")
		nodeID   = flag.String("node", "", "node id for node-facing APIs (default: local)")
		skipNoAuth = flag.Bool("skip-noauth", false, "skip the unauthenticated baseline probe")
	)
	flag.Parse()

	if *key == "" {
		fmt.Fprintln(os.Stderr, "missing -key (API key from the panel)")
		os.Exit(2)
	}

	fmt.Printf("→ 1Panel at %s\n", *baseURL)
	fmt.Printf("→ sign method: %s\n", *signMth)
	fmt.Printf("→ node id:     %q (empty = local)\n", *nodeID)

	if !*skipNoAuth {
		// Sanity check: same call without API key — should be 401.
		c, _ := client.New(client.Config{BaseURL: *baseURL})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = c.Get(ctx, "/dashboard/base/os", nil)
		cancel()
	}

	// With API key.
	sdk, err := onepanel.New(onepanel.Options{
		BaseURL:       *baseURL,
		APIKey:        *key,
		APISignMethod: *signMth,
		NodeID:        *nodeID,
	})
	if err != nil {
		log.Fatalf("onepanel.New: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	probe := func(name string, fn func() (map[string]any, error)) {
		data, err := fn()
		if err != nil {
			fmt.Printf("\n[%s] ERR  %v\n", name, err)
			if perr, ok := err.(*client.Error); ok {
				if perr.APIError != nil {
					fmt.Printf("       code=%d message=%q trace_id=%q\n",
						perr.APIError.Code, perr.APIError.Message, perr.APIError.TraceID)
				} else {
					fmt.Printf("       http_status=%d (non-envelope body)\n", perr.StatusCode)
				}
			}
			return
		}
		pretty, _ := json.MarshalIndent(data, "", "  ")
		// Truncate huge responses.
		if len(pretty) > 1500 {
			pretty = append(pretty[:1500], []byte("\n  ... (truncated)")...)
		}
		fmt.Printf("\n[%s] OK\n%s\n", name, pretty)
	}

	// 1) Device base info — POST /toolbox/device/base (node-facing)
	probe("device/base", func() (map[string]any, error) {
		return sdk.Toolbox.POSTDeviceBase(ctx, nil)
	})

	// 2) Dashboard OS info — GET /dashboard/base/os
	probe("dashboard/base/os", func() (map[string]any, error) {
		return sdk.Dashboard.GETBaseOs(ctx)
	})

	// 3) Dashboard app launcher — returns a top-level array, so the map probe
	//    is the wrong shape. Tested below via GetList.

	// 4) Hosts — list all nodes
	probe("hosts/search", func() (map[string]any, error) {
		return sdk.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 10})
	})

	// 5) OnNode override — point to a non-existent node to confirm CurrentNode header works.
	if *nodeID == "" {
		probeSDK := sdk.OnNode("0")
		probe("hosts/search node=0", func() (map[string]any, error) {
			return probeSDK.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 5})
		})
	}

	// 6) Current user — /core/auth/* requires session cookies, not API key.
	//    We expect a 401 with "用户未登录" message — the SDK should surface it
	//    cleanly (no panic), and the error code should be client.CodeAuthFail.
	probe("auth/current (expect 401)", func() (map[string]any, error) {
		return sdk.Auth.CurrentUser(ctx)
	})

	// 7) Array-returning endpoints — use the public GetList/PostList helpers.
	type arrayResult struct {
		name string
		fn   func() ([]any, error)
	}
	arrays := []arrayResult{
		{"dashboard/app/launcher", func() ([]any, error) {
			return sdk.Dashboard.GetList(ctx, "/dashboard/app/launcher")
		}},
		{"groups/search", func() ([]any, error) {
			return sdk.Groups.PostList(ctx, "/groups/search", map[string]any{})
		}},
	}
	for _, a := range arrays {
		data, err := a.fn()
		if err != nil {
			fmt.Printf("\n[%s] ERR  %v\n", a.name, err)
			if perr, ok := err.(*client.Error); ok && perr.APIError != nil {
				fmt.Printf("       code=%d message=%q\n",
					perr.APIError.Code, perr.APIError.Message)
			}
			continue
		}
		pretty, _ := json.MarshalIndent(data, "", "  ")
		if len(pretty) > 800 {
			pretty = append(pretty[:800], []byte("\n  ... (truncated)")...)
		}
		fmt.Printf("\n[%s] OK (len=%d)\n%s\n", a.name, len(data), pretty)
	}
}
