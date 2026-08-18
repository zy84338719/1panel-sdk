// Command example demonstrates the 1Panel SDK end-to-end.
//
// Usage:
//
//	export ONEPANEL_URL=https://1panel.example.com
//	export ONEPANEL_ENTRANCE=1panel_entrance
//	export ONEPANEL_USER=admin
//	export ONEPANEL_PASS=secret
//	go run ./example
//
// The SDK automatically prepends /api/v2 to every request path
// (matching 1Panel's swagger basePath), so user-supplied paths are
// always the post-strip form used by the 1Panel frontend
// (e.g. /core/auth/login, /containers/search, /websites/ssl).
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zy84338719/1panel-sdk"
	"github.com/zy84338719/1panel-sdk/client"
)

func main() {
	url := envOr("ONEPANEL_URL", "https://1panel.example.com")
	entrance := envOr("ONEPANEL_ENTRANCE", "")
	user := envOr("ONEPANEL_USER", "admin")
	pass := envOr("ONEPANEL_PASS", "")

	if pass == "" {
		log.Fatal("ONEPANEL_PASS is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// === Auto-login ===
	sdk, err := onepanel.New(onepanel.Options{
		BaseURL:  url,
		Entrance: entrance,
		Username: user,
		Password: pass,
		OnLogin: func(r *client.LoginResult) {
			fmt.Printf("logged in as %s, role=%s\n", r.Name, r.Role)
		},
	})
	if err != nil {
		log.Fatalf("login: %v", err)
	}

	// === Dashboard ===
	osInfo, err := sdk.Dashboard.OSInfo(ctx)
	check("dashboard.OSInfo", err)
	printJSON("OSInfo", osInfo)

	// === Hosts ===
	hosts, err := sdk.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 20})
	check("host.SearchHosts", err)
	printJSON("hosts", hosts)

	// === Containers ===
	status, err := sdk.Container.Status(ctx)
	check("container.Status", err)
	printJSON("container.status", status)

	containers, err := sdk.Container.List(ctx)
	check("container.List", err)
	printJSON("containers", containers)

	// === Apps ===
	apps, err := sdk.App.Search(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("app.Search", err)
	printJSON("apps", apps)

	// === Websites ===
	websites, err := sdk.Website.Search(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("website.Search", err)
	printJSON("websites", websites)

	// === Databases ===
	dbs, err := sdk.Database.SearchMysql(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("database.SearchMysql", err)
	printJSON("databases", dbs)

	// === Backups ===
	backups, err := sdk.Backup.Search(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("backup.Search", err)
	printJSON("backups", backups)

	// === Cronjobs ===
	jobs, err := sdk.Cronjob.Search(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("cronjob.Search", err)
	printJSON("cronjobs", jobs)

	// === Settings ===
	setting, err := sdk.Settings.GetSettingInfo(ctx, map[string]any{})
	check("settings.GetSettingInfo", err)
	printJSON("settings", setting)

	// === Logs ===
	logs, err := sdk.Logs.GetOperationLogs(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("logs.GetOperationLogs", err)
	printJSON("logs", logs)

	// === AI / OpenClaw agents (new in 1Panel v2) ===
	ollamaModels, err := sdk.AI.OllamaSearch(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("ai.ollama/model/search", err)
	printJSON("ollama models", ollamaModels)

	// AI agents (note: actually mounted at /ai/agents/...)
	agents, err := sdk.Agent.Search(ctx, map[string]any{"page": 1, "pageSize": 5})
	check("agents.Search", err)
	printJSON("agents", agents)

	// === New sub-services that weren't in the original code review ===
	// Health check (does not require auth).
	var health map[string]any
	err = sdk.Health.Call(ctx, "GET", "/health/check", nil, &health)
	check("health/check", err)
	printJSON("health", health)

	// OpenResty (Nginx) status.
	var openrestyStatus map[string]any
	err = sdk.OpenResty.Call(ctx, "GET", "/openresty/status", nil, &openrestyStatus)
	check("openresty/status", err)
	printJSON("openresty status", openrestyStatus)

	// Runtime list (PHP, Node, etc.).
	var runtimes map[string]any
	err = sdk.Runtime.POSTSearch(ctx, map[string]any{"page": 1, "pageSize": 5}, &runtimes)
	check("runtimes/search", err)
	printJSON("runtimes", runtimes)

	// === Multi-node: target node 1 ===
	if false {
		nodeSDK := sdk.OnNode("1")
		nodeContainers, err := nodeSDK.Container.List(ctx)
		check("node[1].container.List", err)
		printJSON("node[1].containers", nodeContainers)
	}

	// === Wildcard: reach an endpoint not yet typed ===
	// (every service exposes Call(ctx, method, path, body, out) as a backdoor.)
	// rawResp, err := map[string]any{}, error(nil)
	// err = sdk.Dashboard.Call(ctx, "GET", "/dashboard/app/launcher", nil, &rawResp)

	// === Logout ===
	if err := sdk.Auth.Logout(ctx); err != nil {
		log.Printf("logout: %v", err)
	}
}

// === helpers ===

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func check(label string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", label, err)
	}
}

func printJSON(label string, v any) {
	if v == nil {
		fmt.Printf("%s: <nil>\n", label)
		return
	}
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("=== %s ===\n%s\n", label, string(b))
}
