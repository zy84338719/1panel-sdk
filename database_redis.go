package onepanel

import "context"

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
