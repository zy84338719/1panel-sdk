package onepanel

import "context"

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
