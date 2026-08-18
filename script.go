package onepanel

import "context"


type ScriptService struct {
	ServiceBase
}

// Create creates a script.
func (s *ScriptService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/script", body)
}

// Search searches scripts.
func (s *ScriptService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/script/search", body)
}

// Delete deletes a script.
func (s *ScriptService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/script/del", body)
}

// Update updates a script.
func (s *ScriptService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/script/update", body)
}

// Sync syncs scripts to the connected nodes.
func (s *ScriptService) Sync(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/script/sync", body)
}

// Run runs a script.
func (s *ScriptService) Run(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.getMap(ctx, "/script/run")
}

// Call invokes an arbitrary /script/* endpoint.
func (s *ScriptService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// ToolboxService covers /toolbox/* (device, fail2ban, ftp, clamav).
