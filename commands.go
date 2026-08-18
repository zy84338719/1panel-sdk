package onepanel

import "context"

type CommandsService struct {
	ServiceBase
}

// List lists commands.
func (s *CommandsService) List(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/list", body)
}

// Create creates a command.
func (s *CommandsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands", body)
}

// Delete deletes a command.
func (s *CommandsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/del", body)
}

// Search searches commands.
func (s *CommandsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/search", body)
}

// SearchTree returns the command tree.
func (s *CommandsService) SearchTree(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/tree", body)
}

// Update updates a command.
func (s *CommandsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/update", body)
}

// Export exports commands to CSV.
func (s *CommandsService) Export(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/export", body)
}

// UploadCsv uploads a CSV of commands.
func (s *CommandsService) UploadCsv(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/upload", body)
}

// Import imports commands.
func (s *CommandsService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/commands/import", body)
}

// Call invokes an arbitrary /commands/* endpoint.
func (s *CommandsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// ScriptService covers /script/* (script library: shared by panel and nodes).
