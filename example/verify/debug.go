// Quick debug probe — show raw response body for the suspicious endpoints.
//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zy84338719/go-1panel/client"
)

func main() {
	c, _ := client.New(client.Config{
		BaseURL: "http://192.168.108.235:55555",
		APIKey:  "0HeE0VPfS2TY5N1ILGgE322McU7hRdmz",
	})

	endpoints := []struct{ method, path string }{
		{"GET", "/dashboard/base/os"},
		{"POST", "/toolbox/device/base"},
		{"POST", "/hosts/search"},
		{"GET", "/core/auth/current"},
		{"POST", "/groups/search"},
		{"POST", "/core/auth/api/search"},
		{"GET", "/core/auth/api/search"},
		{"POST", "/core/auth/api/generate"},
	}

	for _, e := range endpoints {
		fmt.Printf("\n========== %s %s ==========\n", e.method, e.path)
		full := c.Endpoint() + "/api/v2" + e.path
		req, _ := http.NewRequest(e.method, full, nil)
		req.Header.Set("User-Agent", "debug/1.0")
		req.Header.Set("Accept", "application/json")
		ts := fmt.Sprintf("%d", time.Now().Unix())
		req.Header.Set("1Panel-Timestamp", ts)
		req.Header.Set("1Panel-Token", client.SignToken("hmac-sha256", "0HeE0VPfS2TY5N1ILGgE322McU7hRdmz", ts))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		fmt.Printf("HTTP %s\n", resp.Status)
		fmt.Printf("body (%d bytes): %s\n", len(body), body)
	}
}
