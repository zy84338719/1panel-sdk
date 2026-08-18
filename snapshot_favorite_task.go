package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type SnapshotService struct {
	ServiceBase
}

// LoadSnapshotData loads a snapshot file.
func (s *SnapshotService) LoadSnapshotData(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/snapshot/load", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a snapshot.
func (s *SnapshotService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Recreate recreates a snapshot from a previous state.
func (s *SnapshotService) Recreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/recreate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches snapshots.
func (s *SnapshotService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Import imports a snapshot.
func (s *SnapshotService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/import", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a snapshot.
func (s *SnapshotService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Recover recovers a snapshot.
func (s *SnapshotService) Recover(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/recover", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Rollback rolls back to a snapshot.
func (s *SnapshotService) Rollback(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/rollback", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDescription updates a snapshot's description.
func (s *SnapshotService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/description/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/favorites", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a favorite.
func (s *FavoriteService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/favorites/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches favorites.
func (s *FavoriteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/favorites/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/tasks", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListTasks lists tasks.
func (s *TaskService) ListTasks(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CancelTask cancels a running task.
func (s *TaskService) CancelTask(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/cancel", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TaskLog returns the log of a task.
func (s *TaskService) TaskLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /tasks/* endpoint.
func (s *TaskService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
