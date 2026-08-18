package onepanel

import "context"

// PageInfo is the standard pagination request body used by 1Panel /search endpoints.
type PageInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
}

// Page[T] is the standard pagination response envelope.
type Page[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

// Req is the most common 1Panel request envelope used by "search" style endpoints.
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
// "operation" + an id/name payload.
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
func (b *ServiceBase) Get(ctx context.Context, path string, out any) error    { return b.d.Get(ctx, path, out) }
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
