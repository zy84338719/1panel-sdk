// Package onepanel — runnable godoc examples. Each Example* function below
// shows up under `go doc` and is exercised by `go test ./...`. They are
// checked into the SDK so the docs site / pkg.go.dev always renders a
// concrete code sample.
package onepanel

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Example_login shows the simplest end-to-end usage. The example
// is compile-only (no `// Output:` directive) because it would otherwise
// try to dial 1panel.example.com during `go test`. See
// TestExampleFakesup below for a runnable equivalent that uses httptest.
func Example_login() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sdk, err := New(Options{
		BaseURL:  "https://1panel.example.com",
		Entrance: "1panel_entrance",
		Username: "admin",
		Password: "secret",
	})
	if err != nil {
		fmt.Println("login failed:", err)
		return
	}
	defer func() { _ = sdk.Auth.Logout(ctx) }()

	info, _ := sdk.Dashboard.OSInfo(ctx)
	fmt.Println(info["os"])
}

// Example_multinode shows how to target a specific node via OnNode.
func Example_multinode() {
	// Assume sdk was created via New(...) and is logged in.
	// nodeSDK := sdk.OnNode("node-1")
	// containers, _ := nodeSDK.Container.List(ctx)
	// _ = containers
}

// Example_callWildcard shows the Call() escape hatch for endpoints not
// yet wrapped by a typed method.
func Example_callWildcard() {
	// sdk.SomeService.Call(ctx, "GET", "/some/new/endpoint", nil, &out)
}

// Example_pageRequest shows how to drive a /search endpoint with a PageInfo body.
func Example_pageRequest() {
	// req := PageInfo{Page: 1, PageSize: 20}
	// hosts, _ := sdk.Host.SearchHosts(ctx, req)
}

// TestExampleFakesup spins up an httptest stand-in to make Example_login
// runnable under `go test`. The stand-in is intentionally trivial.
func TestExampleFakesup(t *testing.T) {
	panel := http.NewServeMux()
	panel.HandleFunc("/api/v2/core/auth/login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "1p_session", Value: "x", Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "pcsrftoken", Value: "c", Path: "/"})
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"ok","data":{"name":"a","role":"r","token":"t","mfaStatus":"off"}}`))
	})
	panel.HandleFunc("/api/v2/dashboard/base/os", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"ok","data":{"os":"linux"}}`))
	})
	ts := httptest.NewServer(panel)
	defer ts.Close()

	// Direct call into the example body.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sdk, err := New(Options{BaseURL: ts.URL, Entrance: "1panel_entrance", Username: "admin", Password: "secret"})
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	defer func() { _ = sdk.Auth.Logout(ctx) }()
	info, err := sdk.Dashboard.OSInfo(ctx)
	if err != nil {
		t.Fatalf("OSInfo: %v", err)
	}
	if info["os"] != "linux" {
		t.Fatalf("os = %v, want linux", info["os"])
	}
}
