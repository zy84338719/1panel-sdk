// Smoke tests for the onepanel SDK against an in-process 1Panel stand-in
// (httptest). No real panel is required.
//
// Run with:  go test -v ./...
package onepanel

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/zy84338719/go-1panel/client"
)

// fakePanel is a minimal 1Panel HTTP server for SDK smoke-testing.
type fakePanel struct {
	mu         sync.Mutex
	csrf       string
	apiKey     string
	loginCount int
	// lastRequest is captured per-handler to assert on headers and bodies.
}

func newFakePanel() *fakePanel {
	return &fakePanel{
		csrf:   "fake-csrf-token",
		apiKey: "1panel-test-api-key",
	}
}

func (f *fakePanel) handler() http.Handler {
	mux := http.NewServeMux()

	// Login — returns a session cookie and a "pcsrftoken" cookie.
	mux.HandleFunc("/api/v2/core/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "admin" || body["password"] != "secret" {
			writeErr(w, 401, "invalid credentials")
			return
		}
		ec := r.Header.Get("EntranceCode")
		if ec == "" {
			writeErr(w, 401, "missing EntranceCode")
			return
		}
		if decoded, err := base64.StdEncoding.DecodeString(ec); err != nil || string(decoded) != "1panel_entrance" {
			writeErr(w, 401, "bad EntranceCode")
			return
		}
		f.mu.Lock()
		f.loginCount++
		f.mu.Unlock()
		http.SetCookie(w, &http.Cookie{Name: "1p_session", Value: "fake-cookie", Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "pcsrftoken", Value: f.csrf, Path: "/"})
		writeOK(w, map[string]any{
			"name":      "admin",
			"role":      "super",
			"token":     f.apiKey,
			"mfaStatus": "off",
		})
	})

	// Logout — clears the cookie.
	mux.HandleFunc("/api/v2/core/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "1p_session", Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(w, &http.Cookie{Name: "pcsrftoken", Value: "", Path: "/", MaxAge: -1})
		writeOK(w, map[string]any{"message": "logged out"})
	})

	// Dashboard.
	mux.HandleFunc("/api/v2/dashboard/base/os", func(w http.ResponseWriter, r *http.Request) {
		requireAuth(w, r, f)
		writeOK(w, map[string]any{
			"os":              "linux",
			"platform":        "debian",
			"platformVersion": "12",
			"kernelVersion":   "6.1.0",
			"kernelArch":      "x86_64",
			"virtualization":  "kvm",
			"currentNode":     "local",
		})
	})

	// Containers list (POST + GET shape: array body, not map).
	mux.HandleFunc("/api/v2/containers/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		requireAuth(w, r, f)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(client.Result{
			Code:    client.CodeSuccess,
			Message: "success",
			Data:    json.RawMessage(`[{"name":"web","state":"running","image":"nginx:1.25"},{"name":"db","state":"running","image":"postgres:16"}]`),
		})
	})

	// Container status (GET, map).
	mux.HandleFunc("/api/v2/containers/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		requireAuth(w, r, f)
		writeOK(w, map[string]any{
			"running": 2, "stopped": 0, "total": 2,
		})
	})

	// Container create (POST + CSRF).
	mux.HandleFunc("/api/v2/containers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		requireCSRF(w, r, f)
		requireAuth(w, r, f)
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		writeOK(w, map[string]any{"id": 1, "name": body["name"]})
	})

	// Host list — exercises page-search request shape.
	mux.HandleFunc("/api/v2/hosts/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		requireAuth(w, r, f)
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		page, _ := body["page"].(float64)
		writeOK(w, map[string]any{
			"items": []any{
				map[string]any{"id": 1, "name": "node-a", "addr": "10.0.0.1"},
				map[string]any{"id": 2, "name": "node-b", "addr": "10.0.0.2"},
			},
			"total": 2,
			"page":  int(page),
		})
	})

	// Health check — unauthenticated.
	mux.HandleFunc("/api/v2/health/check", func(w http.ResponseWriter, r *http.Request) {
		writeOK(w, map[string]any{"status": "ok", "time": time.Now().Format(time.RFC3339)})
	})

	// Force a 500 to test the Error type.
	mux.HandleFunc("/api/v2/core/auth/welcome", func(w http.ResponseWriter, r *http.Request) {
		writeErr(w, 500, "boom")
	})

	return mux
}

// requireAuth checks the session cookie.
func requireAuth(w http.ResponseWriter, r *http.Request, f *fakePanel) {
	ck, err := r.Cookie("1p_session")
	if err != nil || ck.Value == "" {
		writeErr(w, 401, "unauthorized")
		return
	}
}

// requireCSRF checks the X-CSRF-Token header for non-safe methods.
func requireCSRF(w http.ResponseWriter, r *http.Request, f *fakePanel) {
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		return
	}
	got := r.Header.Get("X-CSRF-Token")
	if got != f.csrf {
		writeErr(w, 403, "invalid csrf token")
		return
	}
}

func writeOK(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(client.Result{
		Code:    client.CodeSuccess,
		Message: "success",
		Data:    mustJSON(data),
	})
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(client.Result{Code: code, Message: msg})
}

func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}

// === Tests ===

func TestEndToEnd(t *testing.T) {
	panel := newFakePanel()
	ts := httptest.NewServer(panel.handler())
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sdk, err := New(Options{
		BaseURL:  ts.URL,
		Entrance: "1panel_entrance",
		Username: "admin",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("new: %v", err)
	}
	if sdk.Auth == nil || sdk.Dashboard == nil {
		t.Fatal("sub-services not bound")
	}

	// Dashboard.
	osInfo, err := sdk.Dashboard.OSInfo(ctx)
	if err != nil {
		t.Fatalf("OSInfo: %v", err)
	}
	if osInfo["os"] != "linux" {
		t.Fatalf("OSInfo.os = %v, want linux", osInfo["os"])
	}
	if !strings.Contains(sdk.C.Endpoint(), "http") {
		t.Fatalf("endpoint = %q", sdk.C.Endpoint())
	}

	// Container list — server returns array, decode into a slice via Do.
	var containerList []map[string]any
	if err := sdk.Container.Do(ctx, "POST", "/containers/list", map[string]any{}, &containerList); err != nil {
		t.Fatalf("container.List: %v", err)
	}
	if len(containerList) != 2 {
		t.Fatalf("expected 2 containers, got %d", len(containerList))
	}

	// Container status (GET, map).
	status, err := sdk.Container.Status(ctx)
	if err != nil {
		t.Fatalf("container.Status: %v", err)
	}
	if status == nil {
		t.Fatal("expected status body")
	}

	// Hosts search (POST, page body).
	hosts, err := sdk.Host.SearchHosts(ctx, PageInfo{Page: 1, PageSize: 20})
	if err != nil {
		t.Fatalf("host.SearchHosts: %v", err)
	}
	if hosts == nil {
		t.Fatal("expected hosts body")
	}

	// Container create (POST + CSRF).
	created, err := sdk.Container.Create(ctx, map[string]any{"name": "test", "image": "nginx"})
	if err != nil {
		t.Fatalf("container.Create: %v", err)
	}
	if created == nil {
		t.Fatal("expected create response")
	}

	// Logout.
	if err := sdk.Auth.Logout(ctx); err != nil {
		t.Fatalf("logout: %v", err)
	}

	// Health (unauthenticated).
	var health map[string]any
	if err := sdk.Health.Call(ctx, "GET", "/health/check", nil, &health); err != nil {
		t.Fatalf("health: %v", err)
	}
	if health["status"] != "ok" {
		t.Fatalf("health.status = %v, want ok", health["status"])
	}
}

func TestLoginErrors(t *testing.T) {
	ts := httptest.NewServer(newFakePanel().handler())
	defer ts.Close()

	// Bad password.
	if _, err := New(Options{
		BaseURL:  ts.URL,
		Entrance: "1panel_entrance",
		Username: "admin",
		Password: "wrong",
	}); err == nil {
		t.Fatal("expected error for bad password")
	}

	// Missing entrance.
	if _, err := New(Options{
		BaseURL:  ts.URL,
		Username: "admin",
		Password: "secret",
	}); err == nil {
		t.Fatal("expected error for missing EntranceCode")
	}
}

func TestAPIV2PrefixAutoAdded(t *testing.T) {
	// Inspect: send a request and confirm the path the server saw includes /api/v2.
	var gotPath string
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v2/zzz/test", func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		writeOK(w, map[string]any{"ok": true})
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	c, err := client.New(client.Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Get(context.Background(), "/zzz/test", nil); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/api/v2/zzz/test" {
		t.Fatalf("path = %q, want /api/v2/zzz/test", gotPath)
	}
}

func TestErrorPropagation(t *testing.T) {
	panel := newFakePanel()
	ts := httptest.NewServer(panel.handler())
	defer ts.Close()

	c, _ := client.New(client.Config{BaseURL: ts.URL})
	sdk := NewFromClient(c)

	// Welcome endpoint always 500s.
	_, err := sdk.Auth.WelcomePage(context.Background())
	if err == nil {
		t.Fatal("expected error from /core/auth/welcome")
	}
	apiErr, ok := err.(*client.Error)
	if !ok {
		t.Fatalf("error type = %T, want *client.Error", err)
	}
	if apiErr.APIError == nil || apiErr.APIError.Code != 500 {
		t.Fatalf("error code = %v, want 500", apiErr)
	}
}

func TestOnNode(t *testing.T) {
	panel := newFakePanel()
	ts := httptest.NewServer(panel.handler())
	defer ts.Close()

	sdk, err := New(Options{BaseURL: ts.URL, Entrance: "1panel_entrance", Username: "admin", Password: "secret"})
	if err != nil {
		t.Fatal(err)
	}

	// OnNode returns a *different* SDK instance whose sub-services target a
	// specific node.
	nodeSDK := sdk.OnNode("node-7")
	if nodeSDK == nil {
		t.Fatal("OnNode returned nil")
	}
	if nodeSDK == sdk {
		t.Fatal("OnNode should return a different pointer")
	}
	if sdk.Container == nodeSDK.Container {
		t.Fatal("sub-service on the original SDK should differ from the per-node SDK")
	}
}

func TestClientCookieAndCSRF(t *testing.T) {
	ts := httptest.NewServer(newFakePanel().handler())
	defer ts.Close()

	c, _ := client.New(client.Config{BaseURL: ts.URL, Entrance: "1panel_entrance"})
	login, err := c.Login(context.Background(), client.LoginForm{Name: "admin", Password: "secret", AuthSource: "local"})
	if err != nil {
		t.Fatal(err)
	}
	if login == nil || login.Name != "admin" {
		t.Fatal("login failed")
	}
	if c.CSRFToken() == "" {
		t.Fatal("CSRF token should be populated after login")
	}
	// The cookies should be present in the jar.
	cookies := c.CookiesFor(ts.URL)
	var foundCSRF, foundSession bool
	for _, ck := range cookies {
		if ck.Name == "pcsrftoken" {
			foundCSRF = true
		}
		if ck.Name == "1p_session" {
			foundSession = true
		}
	}
	if !foundCSRF {
		t.Error("pcsrftoken cookie missing")
	}
	if !foundSession {
		t.Error("1p_session cookie missing")
	}
}

// TestFullCoverage uses the same swagger-driven path shape the production
// code uses, just against the fake server, to confirm path normalization.
func TestFullCoverage(t *testing.T) {
	ts := httptest.NewServer(newFakePanel().handler())
	defer ts.Close()

	sdk, err := New(Options{BaseURL: ts.URL, Entrance: "1panel_entrance", Username: "admin", Password: "secret"})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	var out map[string]any

	if err := sdk.Dashboard.Call(ctx, "GET", "/dashboard/base/os", nil, &out); err != nil {
		t.Fatalf("Call GET: %v", err)
	}
	if out["os"] != "linux" {
		t.Fatalf("wildcard GET os = %v", out["os"])
	}

	out = nil
	if err := sdk.Host.Call(ctx, "POST", "/hosts/search", PageInfo{Page: 1, PageSize: 10}, &out); err != nil {
		t.Fatalf("Call POST: %v", err)
	}
	if out["page"] == nil {
		t.Fatalf("Call POST page = %v, want present", out["page"])
	}
}

// TestServiceBase confirms every service exposes a usable Call() and
// has the ServiceBase embedded (compile-time proof). We can't use a strict
// interface because some services override Delete/Get/Post with typed
// variants, so we check each method by name via reflection.
func TestServiceBase(t *testing.T) {
	// Use NewFromClient to skip auto-login; we just want to enumerate methods.
	c, err := client.New(client.Config{BaseURL: "http://127.0.0.1:1", Entrance: "x"})
	if err != nil {
		t.Fatal(err)
	}
	sdk := NewFromClient(c)
	services := []any{
		sdk.Auth, sdk.Dashboard, sdk.Host, sdk.Container, sdk.App,
		sdk.Website, sdk.Database, sdk.Backup, sdk.Cronjob, sdk.File,
		sdk.Settings, sdk.Logs, sdk.Groups, sdk.Commands, sdk.Script,
		sdk.Toolbox, sdk.Alerts, sdk.AI, sdk.Agent, sdk.Agents,
		sdk.SSH, sdk.Monitor, sdk.Firewall, sdk.Nginx, sdk.Process,
		sdk.Runtime, sdk.Snapshot, sdk.Favorite, sdk.Task, sdk.Health,
		sdk.OpenResty, sdk.CoreAuth, sdk.CoreBackup, sdk.CoreCommand,
		sdk.CoreGroup, sdk.CoreLog, sdk.CoreScript, sdk.CoreSetting,
		sdk.AIAccount, sdk.AIAgent, sdk.AIDomain, sdk.AIMcp, sdk.AITensor,
	}
	if len(services) < 40 {
		t.Fatalf("only %d services exposed, want >= 40", len(services))
	}
	// Every service should expose at least Get/Post/Do from ServiceBase.
	for _, s := range services {
		rt := reflect.TypeOf(s)
		methodNames := map[string]bool{}
		for i := 0; i < rt.NumMethod(); i++ {
			methodNames[rt.Method(i).Name] = true
		}
		for _, want := range []string{"Get", "Post", "Do", "Call"} {
			if !methodNames[want] {
				t.Errorf("%s missing method %s", rt.String(), want)
			}
		}
	}
}
