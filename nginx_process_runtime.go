package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type NginxService struct {
	ServiceBase
}

// LoadConf loads the Nginx configuration.
func (s *NginxService) LoadConf(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/nginx/conf")
}

// UpdateConf updates the Nginx configuration.
func (s *NginxService) UpdateConf(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/nginx/conf", body)
}

// LoadServerInfo loads a server block.
func (s *NginxService) LoadServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/nginx/server", body)
}

// UpdateServerInfo updates a server block.
func (s *NginxService) UpdateServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/nginx/server/update", body)
}

// Operate runs an Nginx op (start/stop/reload/...).
func (s *NginxService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/nginx/operate", body)
}

// Call invokes an arbitrary /nginx/* endpoint.
func (s *NginxService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// ProcessService covers /processes/* (process listing, killing).

type ProcessService struct {
	ServiceBase
}

// List lists processes.
func (s *ProcessService) List(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/processes", body)
}

// Kill terminates a process.
func (s *ProcessService) Kill(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/processes/kill", body)
}

// Call invokes an arbitrary /processes/* endpoint.
func (s *ProcessService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// RuntimeService covers /hosts/diagnostics/* (runtime diagnostics).

type RuntimeService struct {
	ServiceBase
}

// Summary returns the diagnostics summary.
func (s *RuntimeService) Summary(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/hosts/diagnostics/summary")
}

// Goroutines returns the goroutine stack dump.
func (s *RuntimeService) Goroutines(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/hosts/diagnostics/goroutines")
}

// CreateProfile starts a profile capture.
func (s *RuntimeService) CreateProfile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/diagnostics/profiles", body)
}

// Call invokes an arbitrary /hosts/diagnostics/* endpoint.
func (s *RuntimeService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SnapshotService covers /settings/snapshot/* (panel snapshots).
