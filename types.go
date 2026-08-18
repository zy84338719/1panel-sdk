package onepanel

import "context"

// PageInfo is the standard pagination request body used by 1Panel /search
// endpoints. Page is 1-based; PageSize caps the response.
type PageInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
}

// Page[T] is the standard pagination response envelope. Decode into a
// Page[T] when an endpoint returns {items, total}. Most SDK helpers do
// this for you and yield a map[string]any — use Page[T] only when you
// want typed decoding of the items slice.
type Page[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

// Req is the most common 1Panel request envelope used by "search"-style
// endpoints. The Info/Name/Filters/OrderBy/Order fields cover the
// majority of paginated list endpoints in the 1Panel core/agent API.
type Req struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Info     string `json:"info,omitempty"`
	Name     string `json:"name,omitempty"`
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
	Filters  string `json:"filters,omitempty"`
}

// Operate wraps a generic operation request body. Many endpoints accept
// "operation" + an id/name payload (e.g. start/stop/restart/...).
type Operate struct {
	Operation string `json:"operation"`
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}

// IDReq is a request body that carries only an id.
type IDReq struct {
	ID uint `json:"id"`
}

// IDResult is the standard "id-only" response (used for uploads, creates, etc).
type IDResult struct {
	ID uint `json:"id"`
}

// TaskIDResult is returned for async operations. Track progress via TaskService.
type TaskIDResult struct {
	TaskID string `json:"taskID"`
}

// MessageResult is returned for endpoints that only need to convey success.
type MessageResult struct {
	Message string `json:"message"`
}

// IDSlice is used by batch /del or /update endpoints.
type IDSlice struct {
	IDs []uint `json:"ids"`
}

// ServiceBase is the shared helper every sub-service embeds. It exposes the
// underlying HTTP verbs so that sub-service methods stay short and consistent.
type ServiceBase struct{ d doer }

// Get / Post / Put / Delete / Do proxy the doer.
func (b *ServiceBase) Get(ctx context.Context, path string, out any) error {
	return b.d.Get(ctx, path, out)
}
func (b *ServiceBase) Post(ctx context.Context, path string, body, out any) error {
	return b.d.Post(ctx, path, body, out)
}
func (b *ServiceBase) Put(ctx context.Context, path string, body, out any) error {
	return b.d.Put(ctx, path, body, out)
}
func (b *ServiceBase) Delete(ctx context.Context, path string, body, out any) error {
	return b.d.Delete(ctx, path, body, out)
}
func (b *ServiceBase) Do(ctx context.Context, method, path string, body, out any) error {
	return b.d.Do(ctx, method, path, body, out)
}

// Call is a wildcard for endpoints the typed methods don't wrap. Every service
// inherits this; the path is in the same form the 1Panel frontend uses
// (e.g. "/containers/search") — the client will prepend "/api/v2".
func (b *ServiceBase) Call(ctx context.Context, method, path string, body, out any) error {
	return b.d.Do(ctx, method, path, body, out)
}

// getMap issues a GET and decodes the response into a *map[string]any.

// getMap issues a GET and decodes the response into a *map[string]any.
// Every typed helper in the SDK funnels through this or postMap to keep
// call sites uniform:
//
//	return b.getMap(ctx, "/containers/status")
//
// is equivalent to:
//
//	var out map[string]any
//	if err := b.Get(ctx, "/containers/status", &out); err != nil {
//	    return nil, err
//	}
//	return out, nil
//
// but reads as a single line. Both helpers are used heavily by both the
// hand-written typed methods and the codegen-generated per-endpoint
// wrappers in the zgen_*.go files.
func (b *ServiceBase) getMap(ctx context.Context, path string) (map[string]any, error) {
	var out map[string]any
	if err := b.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// postMap is the POST counterpart of getMap.
func (b *ServiceBase) postMap(ctx context.Context, path string, body any) (map[string]any, error) {
	var out map[string]any
	if err := b.Post(ctx, path, body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// getList decodes a JSON array response. Some 1Panel endpoints (e.g.
// /dashboard/app/launcher, /groups/search) return their data field as a
// top-level array rather than an object, so getMap cannot be used. Use
// getList for those, or call Get/Post directly with a typed slice via Do.
func (b *ServiceBase) getList(ctx context.Context, path string) ([]any, error) {
	var out []any
	if err := b.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// postList is the POST counterpart of getList.
func (b *ServiceBase) postList(ctx context.Context, path string, body any) ([]any, error) {
	var out []any
	if err := b.Post(ctx, path, body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetList is the public, exported alias of getList. Use it for endpoints
// whose data field is a top-level array rather than an object.
func (b *ServiceBase) GetList(ctx context.Context, path string) ([]any, error) {
	return b.getList(ctx, path)
}

// PostList is the public, exported alias of postList.
func (b *ServiceBase) PostList(ctx context.Context, path string, body any) ([]any, error) {
	return b.postList(ctx, path, body)
}
