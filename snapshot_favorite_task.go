package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type SnapshotService struct {
	ServiceBase
}

// LoadSnapshotData loads a snapshot file.
func (s *SnapshotService) LoadSnapshotData(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/snapshot/load")
}

// Create creates a snapshot.
func (s *SnapshotService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot", body)
}

// Recreate recreates a snapshot from a previous state.
func (s *SnapshotService) Recreate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/recreate", body)
}

// Search searches snapshots.
func (s *SnapshotService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/search", body)
}

// Import imports a snapshot.
func (s *SnapshotService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/import", body)
}

// Delete deletes a snapshot.
func (s *SnapshotService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/del", body)
}

// Recover recovers a snapshot.
func (s *SnapshotService) Recover(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/recover", body)
}

// Rollback rolls back to a snapshot.
func (s *SnapshotService) Rollback(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/rollback", body)
}

// UpdateDescription updates a snapshot's description.
func (s *SnapshotService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/snapshot/description/update", body)
}

// Call invokes an arbitrary /settings/snapshot/* endpoint.
func (s *SnapshotService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FavoriteService covers /favorites/* (user favorites/bookmarks).

type FavoriteService struct {
	ServiceBase
}

// Create creates a favorite.
func (s *FavoriteService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/favorites", body)
}

// Delete deletes a favorite.
func (s *FavoriteService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/favorites/del", body)
}

// Search searches favorites.
func (s *FavoriteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/favorites/search", body)
}

// Call invokes an arbitrary /favorites/* endpoint.
func (s *FavoriteService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// TaskService covers /tasks/* (long-running task progress queries).

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
