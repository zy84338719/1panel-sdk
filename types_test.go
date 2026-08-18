package onepanel

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zy84338719/go-1panel/client"
)

// helper — start an httptest server that echoes the standard envelope and
// the requested body shape.
func startArrayServer(t *testing.T, body string) (*httptest.Server, *ServiceBase) {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"message":"","data":` + body + `}`))
	}))
	t.Cleanup(ts.Close)
	return ts, &ServiceBase{d: clientFromURL(t, ts.URL)}
}

func startObjectServer(t *testing.T, body string) (*httptest.Server, *ServiceBase) {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"message":"","data":` + body + `}`))
	}))
	t.Cleanup(ts.Close)
	return ts, &ServiceBase{d: clientFromURL(t, ts.URL)}
}

func clientFromURL(t *testing.T, rawURL string) *client.Client {
	t.Helper()
	c, err := client.New(client.Config{BaseURL: rawURL})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestGetListDecodesArray(t *testing.T) {
	_, base := startArrayServer(t, `[{"id":1},{"id":2},{"id":3}]`)
	got, err := base.getList(context.Background(), "/x")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
}

func TestGetListEmptyArray(t *testing.T) {
	_, base := startArrayServer(t, `[]`)
	got, err := base.getList(context.Background(), "/x")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
}

func TestPostListDecodesArray(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"message":"","data":[{"a":1},{"a":2}]}`))
	}))
	defer ts.Close()
	base := &ServiceBase{d: clientFromURL(t, ts.URL)}
	got, err := base.postList(context.Background(), "/y", map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
}

func TestGetMapObjectResponse(t *testing.T) {
	_, base := startObjectServer(t, `{"x":1,"y":2}`)
	got, err := base.getMap(context.Background(), "/x")
	if err != nil {
		t.Fatal(err)
	}
	if got["x"] != float64(1) {
		t.Fatalf("x = %v, want 1", got["x"])
	}
}

// Document the limitation: getMap on an array response fails to decode.
func TestGetMapFailsOnArrayResponse(t *testing.T) {
	_, base := startArrayServer(t, `[1,2,3]`)
	if _, err := base.getMap(context.Background(), "/x"); err == nil {
		t.Fatal("getMap should fail when data is an array")
	}
}

func TestPublicGetListAlias(t *testing.T) {
	_, base := startArrayServer(t, `[{"a":1},{"a":2}]`)
	got, err := base.GetList(context.Background(), "/x")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	// Sanity check on the type of items.
	first, _ := json.Marshal(got[0])
	if string(first) != `{"a":1}` {
		t.Fatalf("item = %s, want {\"a\":1}", first)
	}
}
