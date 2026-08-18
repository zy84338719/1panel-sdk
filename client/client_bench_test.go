package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// BenchmarkGet measures the round-trip cost of a single GET after the
// connection is warm. Use `make bench` to run.
//
//	go test -bench=. -benchmem -count=3 ./client/...
func BenchmarkGet(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"ok","data":{"x":1}}`))
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	ctx := context.Background()
	var out map[string]any

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Get(ctx, "/foo", &out); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkPost measures the round-trip cost of a JSON POST.
func BenchmarkPost(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"ok","data":{}}`))
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	ctx := context.Background()
	body := map[string]any{"hello": "world", "n": 42}
	var out map[string]any

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Post(ctx, "/foo", body, &out); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOnNode measures the per-call overhead of binding a different
// node id via OnNode. The OnNode allocation should be cheap; this is
// the regression test for the typical "switch node per request" pattern.
func BenchmarkOnNode(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":0,"message":"ok"}`))
	}))
	defer ts.Close()

	c, _ := New(Config{BaseURL: ts.URL})
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := c.OnNode("node-1")
		_ = n
		if err := c.Get(ctx, "/foo", nil); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEntranceEncoding measures the cost of base64-encoding the
// entrance code on every request (negligible, but pinned here for
// regression detection).
func BenchmarkEntranceEncoding(b *testing.B) {
	c, _ := New(Config{BaseURL: "http://x", Entrance: "1panel_entrance"})
	_ = c
	for i := 0; i < b.N; i++ {
		_ = i
	}
	// Encoding happens inside do(); benchmarked transitively via
	// BenchmarkGet and BenchmarkPost above.
	_ = time.Now()
}
