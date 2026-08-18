package onepanel

import "context"

// HostService covers /hosts/* (host management, firewall, SSH, monitor, disk, tools, terminal).
type HostService struct {
	ServiceBase
}

// CreateHost registers a new managed host.
func (s *HostService) CreateHost(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetHost returns the host record by id.
func (s *HostService) GetHost(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/info", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteHost removes a managed host.
func (s *HostService) DeleteHost(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/del", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateHost updates a managed host.
func (s *HostService) UpdateHost(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateHostGroup moves hosts into a different group.
func (s *HostService) UpdateHostGroup(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/update/group", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchHosts paginates managed hosts.
func (s *HostService) SearchHosts(ctx context.Context, req PageInfo) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/search", req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// HostTree returns the host group tree.
func (s *HostService) HostTree(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tree", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TestByInfo tests a host connection described inline.
func (s *HostService) TestByInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/test/byinfo", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TestByID tests a host's connection using its stored credentials.
func (s *HostService) TestByID(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/test/byid", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Disks returns the full disk inventory.
func (s *HostService) Disks(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/disks", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PartitionDisk creates a partition table on a disk.
func (s *HostService) PartitionDisk(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/disks/partition", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// MountDisk mounts a device to a path.
func (s *HostService) MountDisk(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/disks/mount", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UnmountDisk unmounts a device.
func (s *HostService) UnmountDisk(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/disks/unmount", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckComponent verifies that a system component exists (e.g. "docker", "mysql").
func (s *HostService) CheckComponent(ctx context.Context, name string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/components/"+name, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ToolStatus returns the install/active state of a given tool.
func (s *HostService) ToolStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// InitToolConfig initializes a tool's configuration file.
func (s *HostService) InitToolConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/init", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateTool starts/stops/restarts a tool.
func (s *HostService) OperateTool(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ToolConfig returns the config of a tool.
func (s *HostService) ToolConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/config/get", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateToolConfig updates a tool's configuration.
func (s *HostService) UpdateToolConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/config/set", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateProcess controls a supervisor-managed process.
func (s *HostService) OperateProcess(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/supervisor/process", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetProcess returns the supervisor process state.
func (s *HostService) GetProcess(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/tool/supervisor/process", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetProcessFile returns the process configuration file.
func (s *HostService) GetProcessFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/supervisor/process/file/get", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateProcessFile writes/updates the process configuration file.
func (s *HostService) OperateProcessFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/tool/supervisor/process/file", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LocalTerminalURL returns the WebSocket path for a local shell.
func (s *HostService) LocalTerminalURL() string { return "/hosts/terminal/local" }

// HostSSHTerminalURL returns the WebSocket path for an SSH terminal to a managed host.
func (s *HostService) HostSSHTerminalURL() string { return "/hosts/terminal/ssh" }

// ContainerTerminalURL returns the WebSocket path for a container shell.
func (s *HostService) ContainerTerminalURL() string { return "/hosts/terminal/container" }

// Call invokes an arbitrary /hosts/* endpoint.
func (s *HostService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
