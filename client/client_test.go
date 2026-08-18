// Internal tests for the client package. Run with: go test ./client/...
package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
)

func TestNewRequiresBaseURLOrEndpoint(t *testing.T) {
	if _, err := New(Config{}); err == nil {
		t.Fatal("expected error when neither BaseURL nor Endpoint is set")
	}
}

func TestEndpointFallback(t *testing.T) {
	c, err := New(Config{Endpoint: "https://1panel.example.com:9999"})
	if err != nil {
		t.Fatal(err)
	}
	if c.Endpoint() != "https://1panel.example.com:9999" {
		t.Fatalf("Endpoint() = %q", c.Endpoint())
	}

	c2, _ := New(Config{BaseURL: "https://panel.example.com/1panel_entrance"})
	if c2.Endpoint() != "https://panel.example.com/1panel_entrance" {
		t.Fatalf("Endpoint() = %q", c2.Endpoint())
	}
}

func TestAPIV2PrefixAutoAdded(t *testing.T) {
	var gotPath string
	var gotMethod string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotMethod = r.Method
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Result{Code: CodeSuccess, Message: "ok", Data: json.RawMessage(`{"x":1}`)})
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	var out map[string]any
	if err := c.Get(context.Background(), "/foo/bar", &out); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/api/v2/foo/bar" {
		t.Fatalf("path = %q, want /api/v2/foo/bar", gotPath)
	}
	if gotMethod != http.MethodGet {
		t.Fatalf("method = %q, want GET", gotMethod)
	}
}

func TestAPIV2PrefixNotDuplicated(t *testing.T) {
	// If the BaseURL already contains /api/v2, the client should not add it again.
	var gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		_, _ = io.WriteString(w, `{"code":200,"message":"ok","data":{}}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL + "/api/v2"})
	var out map[string]any
	if err := c.Get(context.Background(), "/foo", &out); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/api/v2/foo" {
		t.Fatalf("path = %q, want /api/v2/foo", gotPath)
	}
}

func TestEntranceCodeHeader(t *testing.T) {
	var gotEntrance string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotEntrance = r.Header.Get("EntranceCode")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL, Entrance: "1panel_entrance"})
	if err := c.Post(context.Background(), "/core/auth/login", map[string]any{"name": "a", "password": "b"}, nil); err != nil {
		t.Fatal(err)
	}
	decoded, err := base64.StdEncoding.DecodeString(gotEntrance)
	if err != nil {
		t.Fatalf("EntranceCode not base64: %v", err)
	}
	if string(decoded) != "1panel_entrance" {
		t.Fatalf("decoded entrance = %q, want 1panel_entrance", string(decoded))
	}
}

func TestCSRFTokenOnNonSafeMethods(t *testing.T) {
	var csrfOnPost string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			csrfOnPost = r.Header.Get("X-CSRF-Token")
			// also send a csrf cookie so the client picks it up
			http.SetCookie(w, &http.Cookie{Name: "pcsrftoken", Value: "tok-123", Path: "/"})
		}
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	// First call sets the CSRF cookie via the response Set-Cookie.
	if err := c.Post(context.Background(), "/containers", map[string]any{"name": "x"}, nil); err != nil {
		t.Fatal(err)
	}
	if c.CSRFToken() == "" {
		t.Fatal("expected CSRF token after first POST")
	}
	// Second POST should send X-CSRF-Token.
	if err := c.Post(context.Background(), "/containers/del", map[string]any{"id": 1}, nil); err != nil {
		t.Fatal(err)
	}
	if csrfOnPost != "tok-123" {
		t.Fatalf("X-CSRF-Token = %q, want tok-123", csrfOnPost)
	}
}

func TestGETHasNoCSRF(t *testing.T) {
	var csrfOnGet string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfOnGet = r.Header.Get("X-CSRF-Token")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	// Pre-load a CSRF token; GET must not echo it.
	c.csrfToken = "should-not-be-sent"
	if err := c.Get(context.Background(), "/foo", nil); err != nil {
		t.Fatal(err)
	}
	if csrfOnGet != "" {
		t.Fatalf("X-CSRF-Token = %q on GET, want empty", csrfOnGet)
	}
}

func TestErrorTypeAndCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"code":433,"message":"node unbind","trace_id":"abc"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	err := c.Get(context.Background(), "/agents", nil)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *Error", err)
	}
	if !apiErr.IsCode(CodeNodeUnbind) {
		t.Fatalf("IsCode(433) = false; want true; got code=%d", apiErr.APIError.Code)
	}
	if apiErr.APIError.TraceID != "abc" {
		t.Fatalf("trace_id = %q, want abc", apiErr.APIError.TraceID)
	}
}

func TestNonJSONErrorStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = io.WriteString(w, "bad gateway")
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	err := c.Get(context.Background(), "/foo", nil)
	if err == nil {
		t.Fatal("expected error for 502")
	}
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *Error", err)
	}
	if apiErr.StatusCode != 502 {
		t.Fatalf("StatusCode = %d, want 502", apiErr.StatusCode)
	}
}

func TestCurrentNodeHeader(t *testing.T) {
	var gotNode string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotNode = r.Header.Get("CurrentNode")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL, NodeID: "node-1"})
	if err := c.Get(context.Background(), "/containers/status", nil); err != nil {
		t.Fatal(err)
	}
	decoded, _ := url.QueryUnescape(gotNode)
	if decoded != "node-1" {
		t.Fatalf("CurrentNode = %q (decoded %q), want node-1", gotNode, decoded)
	}
}

func TestOnNodeOverride(t *testing.T) {
	var gotNode string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotNode = r.Header.Get("CurrentNode")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	nc := c.OnNode("worker-7")
	if err := nc.Get(context.Background(), "/containers/list", nil); err != nil {
		t.Fatal(err)
	}
	decoded, _ := url.QueryUnescape(gotNode)
	if decoded != "worker-7" {
		t.Fatalf("CurrentNode = %q (decoded %q), want worker-7", gotNode, decoded)
	}
}

func TestAcceptLanguageHeader(t *testing.T) {
	var gotLang string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotLang = r.Header.Get("Accept-Language")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL, Language: "ja-JP"})
	if err := c.Get(context.Background(), "/foo", nil); err != nil {
		t.Fatal(err)
	}
	if gotLang != "ja-JP" {
		t.Fatalf("Accept-Language = %q, want ja-JP", gotLang)
	}
}

func TestUserAgent(t *testing.T) {
	var gotUA string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.Header.Get("User-Agent")
		_, _ = io.WriteString(w, `{"code":200,"message":"ok"}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	if err := c.Get(context.Background(), "/foo", nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(gotUA, "1panel-sdk") {
		t.Fatalf("User-Agent = %q, want contains 1panel-sdk", gotUA)
	}
}

func TestDoAndBaseDo(t *testing.T) {
	var hits int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		_, _ = io.WriteString(w, `{"code":200,"message":"ok","data":{}}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	var out map[string]any
	// Do.
	if err := c.Do(context.Background(), http.MethodGet, "/x", nil, &out); err != nil {
		t.Fatal(err)
	}
	// Get.
	if err := c.Get(context.Background(), "/x", &out); err != nil {
		t.Fatal(err)
	}
	if atomic.LoadInt32(&hits) != 2 {
		t.Fatalf("hits = %d, want 2", hits)
	}
}

func TestCookiesAndCSRFAfterLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "1p_session", Value: "sess-1", Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "pcsrftoken", Value: "csrf-1", Path: "/"})
		_, _ = io.WriteString(w, `{"code":200,"message":"ok","data":{"name":"a","role":"r","token":"t","mfaStatus":"off"}}`)
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	login, err := c.Login(context.Background(), LoginForm{Name: "a", Password: "b"})
	if err != nil {
		t.Fatal(err)
	}
	if login == nil || login.Name != "a" {
		t.Fatal("login failed")
	}
	if c.CSRFToken() != "csrf-1" {
		t.Fatalf("CSRF = %q, want csrf-1", c.CSRFToken())
	}
	cookies := c.CookiesFor(ts.URL)
	var found bool
	for _, ck := range cookies {
		if ck.Name == "1p_session" {
			found = true
		}
	}
	if !found {
		t.Fatal("1p_session cookie missing from jar")
	}
}

func TestAsInt(t *testing.T) {
	cases := []struct {
		in   any
		want int64
		ok   bool
	}{
		{int(7), 7, true},
		{int64(8), 8, true},
		{float64(9), 9, true},
		{"10", 10, true},
		{"x", 0, false},
		{nil, 0, false},
	}
	for _, c := range cases {
		got, ok := AsInt(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("AsInt(%v) = (%d, %v), want (%d, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}
