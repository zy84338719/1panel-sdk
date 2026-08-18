package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type GroupsService struct {
	ServiceBase
}

// Create creates a group.
func (s *GroupsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/groups", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a group.
func (s *GroupsService) Delete(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/groups/del", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a group.
func (s *GroupsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/groups/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// List lists groups.
func (s *GroupsService) List(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/groups/search", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/commands/list", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a command.
func (s *CommandsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a command.
func (s *CommandsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches commands.
func (s *CommandsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTree returns the command tree.
func (s *CommandsService) SearchTree(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/tree", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a command.
func (s *CommandsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Export exports commands to CSV.
func (s *CommandsService) Export(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/export", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UploadCsv uploads a CSV of commands.
func (s *CommandsService) UploadCsv(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/upload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Import imports commands.
func (s *CommandsService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/commands/import", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/script", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches scripts.
func (s *ScriptService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/script/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a script.
func (s *ScriptService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/script/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a script.
func (s *ScriptService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/script/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Sync syncs scripts to the connected nodes.
func (s *ScriptService) Sync(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/script/sync", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Run runs a script.
func (s *ScriptService) Run(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/script/run", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /script/* endpoint.
func (s *ScriptService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// ToolboxService covers /toolbox/* (device, fail2ban, ftp, clamav).
