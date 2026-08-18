package onepanel

import "context"


type SSHService struct {
	ServiceBase
}

// GetSSHInfo returns the SSH config.
func (s *SSHService) GetSSHInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/search", body)
}

// UpdateSSH updates the SSH config.
func (s *SSHService) UpdateSSH(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/update", body)
}

// LoadSSHLogs loads SSH logs.
func (s *SSHService) LoadSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log", body)
}

// ExportSSHLogs exports SSH logs.
func (s *SSHService) ExportSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log/export", body)
}

// CleanSSHLogs cleans SSH logs.
func (s *SSHService) CleanSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log/clean", body)
}

// Operate runs a SSH op (start/stop/restart).
func (s *SSHService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/operate", body)
}

// LoadSSHFile loads the SSH config file.
func (s *SSHService) LoadSSHFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/file", body)
}

// UpdateSSHByFile replaces the SSH config file.
func (s *SSHService) UpdateSSHByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/file/update", body)
}

// CreateRootCert creates a root CA certificate.
func (s *SSHService) CreateRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert", body)
}

// EditRootCert edits a root CA certificate.
func (s *SSHService) EditRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/update", body)
}

// SyncRootCert syncs the root CA to nodes.
func (s *SSHService) SyncRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/sync", body)
}

// SearchRootCert searches root certificates.
func (s *SSHService) SearchRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/search", body)
}

// DeleteRootCert deletes a root certificate.
func (s *SSHService) DeleteRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/delete", body)
}

// Call invokes an arbitrary /hosts/ssh/* endpoint.
func (s *SSHService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// MonitorService covers /hosts/monitor/* (CPU / memory / disk / network monitor).

