package onepanel

import "context"

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
