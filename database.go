package onepanel

import "context"
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

// === Redis ===

// LoadRedisPersistenceConf loads the redis persistence configuration.
func (s *DatabaseService) LoadRedisPersistenceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/persistence/conf", body)
}

// LoadRedisStatus returns the redis status info.
func (s *DatabaseService) LoadRedisStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/status", body)
}

// LoadRedisConf returns the redis config.
func (s *DatabaseService) LoadRedisConf(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/conf", body)
}

// CheckHasCli returns whether redis-cli is installed.
func (s *DatabaseService) CheckHasCli(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/databases/redis/check")
}

// InstallCli installs the redis-cli.
func (s *DatabaseService) InstallCli(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/install/cli", body)
}

// ChangeRedisPassword changes the redis password.
func (s *DatabaseService) ChangeRedisPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/password", body)
}

// UpdateRedisConf updates the redis configuration.
func (s *DatabaseService) UpdateRedisConf(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/conf/update", body)
}

// UpdateRedisPersistenceConf updates the redis persistence configuration.
func (s *DatabaseService) UpdateRedisPersistenceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/redis/persistence/update", body)
}

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
func (s *DatabaseService) LoadPostgresqlDBFromRemote(ctx context.Context, database string, body map[string]any) (map[string]any, error) {
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
func (s *DatabaseService) UpdatePostgresqlDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/pg/description", body)
}

// === MongoDB ===

// CreateMongodb creates a MongoDB instance.
func (s *DatabaseService) CreateMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb", body)
}

// SearchMongodb searches MongoDB instances.
func (s *DatabaseService) SearchMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/search", body)
}

// UpdateMongodbDescription updates the description.
func (s *DatabaseService) UpdateMongodbDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/description", body)
}

// LoadMongodbFromRemote loads a MongoDB DB from a remote server.
func (s *DatabaseService) LoadMongodbFromRemote(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/load", body)
}

// BindMongodbUser binds a user to a database.
func (s *DatabaseService) BindMongodbUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/bind", body)
}

// ChangeMongodbPassword changes the user password.
func (s *DatabaseService) ChangeMongodbPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/password", body)
}

// ChangeMongodbRootPassword changes the root password.
func (s *DatabaseService) ChangeMongodbRootPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/root/password", body)
}

// LoadMongodbPrivileges loads current mongodb privileges.
func (s *DatabaseService) LoadMongodbPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/privileges", body)
}

// ChangeMongodbPrivileges updates mongodb privileges.
func (s *DatabaseService) ChangeMongodbPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/privileges/change", body)
}

// DeleteCheckMongodb pre-flight check.
func (s *DatabaseService) DeleteCheckMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/del/check", body)
}

// DeleteMongodb deletes a MongoDB instance.
func (s *DatabaseService) DeleteMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/databases/mongodb/del", body)
}

// Call invokes an arbitrary /databases/* endpoint.
func (s *DatabaseService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// itoa converts a uint to a string (avoids importing strconv at every call site).
func itoa(u uint) string {
	if u == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for u > 0 {
		i--
		buf[i] = byte('0' + u%10)
		u /= 10
	}
	return string(buf[i:])
}


