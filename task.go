package onepanel

import "context"


type TaskService struct {
	ServiceBase
}

// SSLService is a thin alias for WebsiteSSLService kept for naming symmetry
// with the rest of the SDK. It exposes the same /websites/ssl/* endpoints.
type SSLService = WebsiteSSLService

// GetTask returns a task by id.

// GetTask returns a task by id.
func (s *TaskService) GetTask(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/tasks", body)
}

// ListTasks lists tasks.
func (s *TaskService) ListTasks(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/tasks/search", body)
}

// CancelTask cancels a running task.
func (s *TaskService) CancelTask(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/tasks/cancel", body)
}

// TaskLog returns the log of a task.
func (s *TaskService) TaskLog(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/tasks/log", body)
}

// Call invokes an arbitrary /tasks/* endpoint.
func (s *TaskService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
