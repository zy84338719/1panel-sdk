// Package client - shared DTOs used by both core and node APIs.
package client

import (
	"context"
	"encoding/json"
)

// Page is the standard 1Panel pagination request body used by /search endpoints.
type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// PageResp is the standard page envelope.
type PageResp[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

// PageOf is a small helper that pages a paginated endpoint into the result type.
// Caller provides a pointer to T-shaped result.
func PageOf[T any]() *PageResp[T] { return &PageResp[T]{} }

// IDBody is the standard "operate by id" body for /del, /update, /info style endpoints.
type IDBody struct {
	ID uint `json:"id"`
}

// NameBody is a "operate by name" body. Many agent endpoints take names instead of ids.
type NameBody struct {
	Name string `json:"name"`
}

// IDsBody accepts a list of ids for batch operations.
type IDsBody struct {
	IDs []uint `json:"ids"`
}

// PageInfo is intentionally not defined here. The SDK exposes
// onepanel.PageInfo at the top level; importing client.PageInfo would
// duplicate the type and confuse users. Callers that need a pagination
// request body should construct it via onepanel.PageInfo.

// RawData unmarshals a Result.Data into a map without losing precision.
func RawData(result *Result) (map[string]any, error) {
	if result == nil || len(result.Data) == 0 {
		return map[string]any{}, nil
	}
	out := map[string]any{}
	if err := json.Unmarshal(result.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// MustContext returns ctx or context.Background() when ctx is nil.
func MustContext(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}
