package onepanel

import "context"
type DatabaseService struct {
	ServiceBase
}

// DBBaseInfo returns the runtime overview of a database instance.
func (s *DatabaseService) DBBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/common/info", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadDBFile loads a database config file.
func (s *DatabaseService) LoadDBFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/common/load/file", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDBConfByFile replaces the database config file.
func (s *DatabaseService) UpdateDBConfByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/common/update/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateMysql creates a MySQL instance.
func (s *DatabaseService) CreateMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListMysqlUsers lists users of a MySQL instance.
func (s *DatabaseService) ListMysqlUsers(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateMysqlUser creates a MySQL user.
func (s *DatabaseService) CreateMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMysqlUser updates a MySQL user.
func (s *DatabaseService) UpdateMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMysqlUserPassword changes a MySQL user's password.
func (s *DatabaseService) ChangeMysqlUserPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SaveMysqlUserPassword saves a MySQL user's password (local record).
func (s *DatabaseService) SaveMysqlUserPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users/password/save", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteMysqlUser deletes a MySQL user.
func (s *DatabaseService) DeleteMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/users/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListMysqlGrants lists grants.
func (s *DatabaseService) ListMysqlGrants(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/grants/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListMysqlGrantSummary lists the grant summary.
func (s *DatabaseService) ListMysqlGrantSummary(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/grants/summary", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GrantMysqlUser grants privileges to a MySQL user.
func (s *DatabaseService) GrantMysqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/grants", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RevokeMysqlGrant revokes a grant.
func (s *DatabaseService) RevokeMysqlGrant(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/grants/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadDBFromRemote imports a database from a remote server.
func (s *DatabaseService) LoadDBFromRemote(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMysqlAccess toggles remote/local access.
func (s *DatabaseService) ChangeMysqlAccess(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/change/access", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMysqlPassword changes the root password.
func (s *DatabaseService) ChangeMysqlPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/change/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheckMysql pre-flight check before deletion.
func (s *DatabaseService) DeleteCheckMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/del/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteMysql deletes a MySQL instance.
func (s *DatabaseService) DeleteMysql(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/del", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMysqlDescription updates a MySQL description.
func (s *DatabaseService) UpdateMysqlDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/description/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMysqlVariables updates MySQL variables.
func (s *DatabaseService) UpdateMysqlVariables(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/variables/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchMysql searches MySQL instances.
func (s *DatabaseService) SearchMysql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadMysqlVariables returns the current MySQL variables.
func (s *DatabaseService) LoadMysqlVariables(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/variables", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadMysqlStatus returns the MySQL status info.
func (s *DatabaseService) LoadMysqlStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRemoteAccess returns the remote access status.
func (s *DatabaseService) LoadRemoteAccess(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/remote", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListDBFormatCollationOptions lists available collation options.
func (s *DatabaseService) ListDBFormatCollationOptions(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/format/options", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Redis ===

// LoadRedisPersistenceConf loads the redis persistence configuration.
func (s *DatabaseService) LoadRedisPersistenceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/persistence/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRedisStatus returns the redis status info.
func (s *DatabaseService) LoadRedisStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRedisConf returns the redis config.
func (s *DatabaseService) LoadRedisConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckHasCli returns whether redis-cli is installed.
func (s *DatabaseService) CheckHasCli(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/databases/redis/check", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// InstallCli installs the redis-cli.
func (s *DatabaseService) InstallCli(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/install/cli", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeRedisPassword changes the redis password.
func (s *DatabaseService) ChangeRedisPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRedisConf updates the redis configuration.
func (s *DatabaseService) UpdateRedisConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/conf/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRedisPersistenceConf updates the redis persistence configuration.
func (s *DatabaseService) UpdateRedisPersistenceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/redis/persistence/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === DB generic (per-database engine) ===

// CheckDatabase checks the database existence.
func (s *DatabaseService) CheckDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateDatabase creates a logical database.
func (s *DatabaseService) CreateDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetDatabase returns a logical database by name.
func (s *DatabaseService) GetDatabase(ctx context.Context, name string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/databases/db/"+name, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListDatabases lists logical databases.
func (s *DatabaseService) ListDatabases(ctx context.Context, kind string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/databases/db/list/"+kind, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadDatabaseItems loads the per-database dashboard items.
func (s *DatabaseService) LoadDatabaseItems(ctx context.Context, kind string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/databases/db/item/"+kind, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDatabase updates a logical database.
func (s *DatabaseService) UpdateDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchDatabases searches logical databases.
func (s *DatabaseService) SearchDatabases(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheckDatabase pre-flight check.
func (s *DatabaseService) DeleteCheckDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db/del/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteDatabase deletes a logical database.
func (s *DatabaseService) DeleteDatabase(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/db/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === PostgreSQL ===

// CreatePostgresql creates a PostgreSQL instance.
func (s *DatabaseService) CreatePostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchPostgresql searches PostgreSQL instances.
func (s *DatabaseService) SearchPostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadPostgresqlDBFromRemote loads a PostgreSQL DB from a remote server.
func (s *DatabaseService) LoadPostgresqlDBFromRemote(ctx context.Context, database string, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/"+database+"/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BindPostgresqlUser binds a user to a database.
func (s *DatabaseService) BindPostgresqlUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/bind", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheckPostgresql pre-flight check.
func (s *DatabaseService) DeleteCheckPostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/del/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeletePostgresql deletes a PostgreSQL instance.
func (s *DatabaseService) DeletePostgresql(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangePostgresqlPrivileges toggles privileges.
func (s *DatabaseService) ChangePostgresqlPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/privileges", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangePostgresqlPassword changes the password.
func (s *DatabaseService) ChangePostgresqlPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatePostgresqlDescription updates the description.
func (s *DatabaseService) UpdatePostgresqlDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/pg/description", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === MongoDB ===

// CreateMongodb creates a MongoDB instance.
func (s *DatabaseService) CreateMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchMongodb searches MongoDB instances.
func (s *DatabaseService) SearchMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMongodbDescription updates the description.
func (s *DatabaseService) UpdateMongodbDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/description", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadMongodbFromRemote loads a MongoDB DB from a remote server.
func (s *DatabaseService) LoadMongodbFromRemote(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BindMongodbUser binds a user to a database.
func (s *DatabaseService) BindMongodbUser(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/bind", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMongodbPassword changes the user password.
func (s *DatabaseService) ChangeMongodbPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMongodbRootPassword changes the root password.
func (s *DatabaseService) ChangeMongodbRootPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/root/password", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadMongodbPrivileges loads current mongodb privileges.
func (s *DatabaseService) LoadMongodbPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/privileges", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeMongodbPrivileges updates mongodb privileges.
func (s *DatabaseService) ChangeMongodbPrivileges(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/privileges/change", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheckMongodb pre-flight check.
func (s *DatabaseService) DeleteCheckMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/del/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteMongodb deletes a MongoDB instance.
func (s *DatabaseService) DeleteMongodb(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/databases/mongodb/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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


