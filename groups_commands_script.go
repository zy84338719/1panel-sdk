package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type GroupsService struct {
	ServiceBase
}

// Create creates a group.
func (s *GroupsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/groups", body)
}

// Delete deletes a group.
func (s *GroupsService) Delete(ctx context.Context, id uint) (map[string]any, error) {
return s.postMap(ctx, "/groups/del", IDReq{ID: id})
}

// Update updates a group.
func (s *GroupsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/groups/update", body)
}

// List lists groups.
func (s *GroupsService) List(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/groups/search", nil)
}

// Call invokes an arbitrary /groups/* endpoint.
func (s *GroupsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// CommandsService covers /commands/* (user-defined command library).

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
