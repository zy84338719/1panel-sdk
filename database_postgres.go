package onepanel

import "context"

// === PostgreSQL ===

// CreatePostgresql creates a PostgreSQL instance.
func (s *DatabaseService) CreatePostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg", body)
}

// SearchPostgresql searches PostgreSQL instances.
func (s *DatabaseService) SearchPostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/search", body)
}

// LoadPostgresqlDBFromRemote loads a PostgreSQL DB from a remote server.
func (s *DatabaseService) LoadPostgresqlDBFromRemote(
	ctx context.Context,
	database string,
	body map[string]any,
) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/"+database+"/load", body)
}

// BindPostgresqlUser binds a user to a database.
func (s *DatabaseService) BindPostgresqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/bind", body)
}

// DeleteCheckPostgresql pre-flight check.
func (s *DatabaseService) DeleteCheckPostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/del/check", body)
}

// DeletePostgresql deletes a PostgreSQL instance.
func (s *DatabaseService) DeletePostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/del", body)
}

// ChangePostgresqlPrivileges toggles privileges.
func (s *DatabaseService) ChangePostgresqlPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/privileges", body)
}

// ChangePostgresqlPassword changes the password.
func (s *DatabaseService) ChangePostgresqlPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/password", body)
}

// UpdatePostgresqlDescription updates the description.
func (s *DatabaseService) UpdatePostgresqlDescription(
	ctx context.Context,
	body map[string]any,
) (map[string]any, error) {
	return s.postMap(ctx, "/databases/pg/description", body)
}
