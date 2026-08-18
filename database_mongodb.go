package onepanel

import "context"
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


