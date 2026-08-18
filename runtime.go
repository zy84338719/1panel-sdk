package onepanel

import "context"

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
