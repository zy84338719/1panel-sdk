package onepanel

import "context"
// === MySQL ===
type DatabaseService struct {
	ServiceBase
}

// DBBaseInfo returns the runtime overview of a database instance.
func (s *DatabaseService) DBBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/common/info", body)
}

// LoadDBFile loads a database config file.
func (s *DatabaseService) LoadDBFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/common/load/file", body)
}

// UpdateDBConfByFile replaces the database config file.
func (s *DatabaseService) UpdateDBConfByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/common/update/conf", body)
}

// CreateMysql creates a MySQL instance.
func (s *DatabaseService) CreateMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases", body)
}

// ListMysqlUsers lists users of a MySQL instance.
func (s *DatabaseService) ListMysqlUsers(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users/search", body)
}

// CreateMysqlUser creates a MySQL user.
func (s *DatabaseService) CreateMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users", body)
}

// UpdateMysqlUser updates a MySQL user.
func (s *DatabaseService) UpdateMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users/update", body)
}

// ChangeMysqlUserPassword changes a MySQL user's password.
func (s *DatabaseService) ChangeMysqlUserPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users/password", body)
}

// SaveMysqlUserPassword saves a MySQL user's password (local record).
func (s *DatabaseService) SaveMysqlUserPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users/password/save", body)
}

// DeleteMysqlUser deletes a MySQL user.
func (s *DatabaseService) DeleteMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/users/del", body)
}

// ListMysqlGrants lists grants.
func (s *DatabaseService) ListMysqlGrants(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/grants/search", body)
}

// ListMysqlGrantSummary lists the grant summary.
func (s *DatabaseService) ListMysqlGrantSummary(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/grants/summary", body)
}

// GrantMysqlUser grants privileges to a MySQL user.
func (s *DatabaseService) GrantMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/grants", body)
}

// RevokeMysqlGrant revokes a grant.
func (s *DatabaseService) RevokeMysqlGrant(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/grants/del", body)
}

// LoadDBFromRemote imports a database from a remote server.
func (s *DatabaseService) LoadDBFromRemote(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/load", body)
}

// ChangeMysqlAccess toggles remote/local access.
func (s *DatabaseService) ChangeMysqlAccess(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/change/access", body)
}

// ChangeMysqlPassword changes the root password.
func (s *DatabaseService) ChangeMysqlPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/change/password", body)
}

// DeleteCheckMysql pre-flight check before deletion.
func (s *DatabaseService) DeleteCheckMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/del/check", body)
}

// DeleteMysql deletes a MySQL instance.
func (s *DatabaseService) DeleteMysql(ctx context.Context, id uint) (map[string]any, error) {
return s.postMap(ctx, "/databases/del", IDReq{ID: id})
}

// UpdateMysqlDescription updates a MySQL description.
func (s *DatabaseService) UpdateMysqlDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/description/update", body)
}

// UpdateMysqlVariables updates MySQL variables.
func (s *DatabaseService) UpdateMysqlVariables(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/variables/update", body)
}

// SearchMysql searches MySQL instances.
func (s *DatabaseService) SearchMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/search", body)
}

// LoadMysqlVariables returns the current MySQL variables.
func (s *DatabaseService) LoadMysqlVariables(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/variables", body)
}

// LoadMysqlStatus returns the MySQL status info.
func (s *DatabaseService) LoadMysqlStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/status", body)
}

// LoadRemoteAccess returns the remote access status.
func (s *DatabaseService) LoadRemoteAccess(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/remote", body)
}

// ListDBFormatCollationOptions lists available collation options.
func (s *DatabaseService) ListDBFormatCollationOptions(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/format/options", body)
}

