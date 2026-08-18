package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type SettingsService struct {
	ServiceBase
}

// GetSettingInfo returns the panel setting info.
func (s *SettingsService) GetSettingInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetBaseInfo returns the panel's base info.
func (s *SettingsService) GetBaseInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/search/base", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetTerminalSettingInfo returns the terminal setting info.
func (s *SettingsService) GetTerminalSettingInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/terminal/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetSystemAvailable returns the available system info.
func (s *SettingsService) GetSystemAvailable(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/search/available", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates settings.
func (s *SettingsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateTerminal updates terminal settings.
func (s *SettingsService) UpdateTerminal(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/terminal/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadInterfaceAddr returns the panel interface address.
func (s *SettingsService) LoadInterfaceAddr(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/interface", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMenu updates the menu layout.
func (s *SettingsService) UpdateMenu(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/menu/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultMenu resets the menu to defaults.
func (s *SettingsService) DefaultMenu(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/menu/default", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateProxy updates the proxy configuration.
func (s *SettingsService) UpdateProxy(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/proxy/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateBindInfo updates the bind address.
func (s *SettingsService) UpdateBindInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/bind/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatePort updates the panel port.
func (s *SettingsService) UpdatePort(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/port/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateSSL updates the panel SSL configuration.
func (s *SettingsService) UpdateSSL(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssl/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadFromCert loads the SSL from the certificate.
func (s *SettingsService) LoadFromCert(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/ssl/info", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadSSL downloads the panel's SSL cert.
func (s *SettingsService) DownloadSSL(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssl/download", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ReloadSSL reloads the SSL configuration.
func (s *SettingsService) ReloadSSL(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssl/reload", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Upgrade upgrades the panel.
func (s *SettingsService) Upgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/upgrade", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetNotesByVersion returns the release notes for a version.
func (s *SettingsService) GetNotesByVersion(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/upgrade/notes", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRelease loads the available releases.
func (s *SettingsService) LoadRelease(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/upgrade/releases", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetUpgradeInfo returns upgrade info.
func (s *SettingsService) GetUpgradeInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/upgrade", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAppstoreConfig updates the app store configuration.
func (s *SettingsService) UpdateAppstoreConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/apps/store/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetAppstoreConfig returns the app store configuration.
func (s *SettingsService) GetAppstoreConfig(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/apps/store/config", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetMemo returns the user memo.
func (s *SettingsService) GetMemo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/memo", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMemo updates the user memo.
func (s *SettingsService) UpdateMemo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/memo", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DescriptionSave saves a description tag.
func (s *SettingsService) DescriptionSave(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/description/save", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseDir returns the panel's base directory.
func (s *SettingsService) BaseDir(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/basedir", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckLocalConn tests the local SSH connection.
func (s *SettingsService) CheckLocalConn(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssh/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadLocalConn returns the local SSH connection details.
func (s *SettingsService) LoadLocalConn(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/ssh/conn", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SetDefaultIsConn sets the local host as the default for SSH operations.
func (s *SettingsService) SetDefaultIsConn(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssh/default", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SaveLocalConn saves the local SSH connection.
func (s *SettingsService) SaveLocalConn(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssh", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckLocalConnByInfo tests a local SSH connection by inline info.
func (s *SettingsService) CheckLocalConnByInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/ssh/check/info", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetTerminalAI returns the terminal AI setting info.
func (s *SettingsService) GetTerminalAI(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/terminal/ai/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateTerminalAI updates the terminal AI setting.
func (s *SettingsService) UpdateTerminalAI(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/terminal/ai/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetFileManageAI returns the file manager AI setting.
func (s *SettingsService) GetFileManageAI(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/files/ai/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateFileManageAI updates the file manager AI setting.
func (s *SettingsService) UpdateFileManageAI(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/files/ai/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetFileHistorySetting returns the file-history setting.
func (s *SettingsService) GetFileHistorySetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/file-history/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateFileHistorySetting updates the file-history setting.
func (s *SettingsService) UpdateFileHistorySetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/file-history/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadWebsiteDir returns the websites root directory.
func (s *SettingsService) LoadWebsiteDir(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/website/dir", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /settings/* endpoint.
func (s *SettingsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// LogsService covers /logs/* (operation logs, system logs, tasks).

type LogsService struct {
	ServiceBase
}

// GetLoginLogs returns login logs.
func (s *LogsService) GetLoginLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/login", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetOperationLogs returns operation logs.
func (s *LogsService) GetOperationLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/operation", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanLogs cleans old logs.
func (s *LogsService) CleanLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetSystemFiles returns available system log files.
func (s *LogsService) GetSystemFiles(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/logs/system/files", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetSystemLogStatus returns the system log status.
func (s *LogsService) GetSystemLogStatus(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/logs/system/status", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ReadSystemLog reads a system log.
func (s *LogsService) ReadSystemLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/system/read", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListRunningServices lists running services.
func (s *LogsService) ListRunningServices(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/logs/system/services", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PageTasks paginates the task list.
func (s *LogsService) PageTasks(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/tasks/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ReadTaskLogByLine reads a task log.
func (s *LogsService) ReadTaskLogByLine(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/logs/tasks/read", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CountExecutingTasks returns the number of running tasks.
func (s *LogsService) CountExecutingTasks(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/logs/tasks/executing/count", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /logs/* endpoint.
func (s *LogsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// GroupsService covers /groups/* (host / app / website group management).
