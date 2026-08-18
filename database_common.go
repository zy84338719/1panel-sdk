package onepanel

import "context"

// === DB generic (per-database engine) ===

// CheckDatabase checks the database existence.
func (s *DatabaseService) CheckDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db/check", body)
}

// CreateDatabase creates a logical database.
func (s *DatabaseService) CreateDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db", body)
}

// GetDatabase returns a logical database by name.
func (s *DatabaseService) GetDatabase(ctx context.Context, name string) (map[string]any, error) {
	return s.getMap(ctx, "/databases/db/"+name)
}

// ListDatabases lists logical databases.
func (s *DatabaseService) ListDatabases(ctx context.Context, kind string) (map[string]any, error) {
	return s.getMap(ctx, "/databases/db/list/"+kind)
}

// LoadDatabaseItems loads the per-database dashboard items.
func (s *DatabaseService) LoadDatabaseItems(ctx context.Context, kind string) (map[string]any, error) {
	return s.getMap(ctx, "/databases/db/item/"+kind)
}

// UpdateDatabase updates a logical database.
func (s *DatabaseService) UpdateDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db/update", body)
}

// SearchDatabases searches logical databases.
func (s *DatabaseService) SearchDatabases(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db/search", body)
}

// DeleteCheckDatabase pre-flight check.
func (s *DatabaseService) DeleteCheckDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db/del/check", body)
}

// DeleteDatabase deletes a logical database.
func (s *DatabaseService) DeleteDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/db/del", body)
}
