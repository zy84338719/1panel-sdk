package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type SettingsService struct {
	ServiceBase
}

// GetSettingInfo returns the panel setting info.
func (s *SettingsService) GetSettingInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/search", body)
}

// GetBaseInfo returns the panel's base info.
func (s *SettingsService) GetBaseInfo(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/settings/search/base", nil)
}

// GetTerminalSettingInfo returns the terminal setting info.
func (s *SettingsService) GetTerminalSettingInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/terminal/search", body)
}

// GetSystemAvailable returns the available system info.
func (s *SettingsService) GetSystemAvailable(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/search/available")
}

// Update updates settings.
func (s *SettingsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/update", body)
}

// UpdateTerminal updates terminal settings.
func (s *SettingsService) UpdateTerminal(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/terminal/update", body)
}

// LoadInterfaceAddr returns the panel interface address.
func (s *SettingsService) LoadInterfaceAddr(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/interface")
}

// UpdateMenu updates the menu layout.
func (s *SettingsService) UpdateMenu(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/menu/update", body)
}

// DefaultMenu resets the menu to defaults.
func (s *SettingsService) DefaultMenu(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/settings/menu/default", nil)
}

// UpdateProxy updates the proxy configuration.
func (s *SettingsService) UpdateProxy(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/proxy/update", body)
}

// UpdateBindInfo updates the bind address.
func (s *SettingsService) UpdateBindInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/bind/update", body)
}

// UpdatePort updates the panel port.
func (s *SettingsService) UpdatePort(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/port/update", body)
}

// UpdateSSL updates the panel SSL configuration.
func (s *SettingsService) UpdateSSL(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssl/update", body)
}

// LoadFromCert loads the SSL from the certificate.
func (s *SettingsService) LoadFromCert(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/ssl/info")
}

// DownloadSSL downloads the panel's SSL cert.
func (s *SettingsService) DownloadSSL(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssl/download", body)
}

// ReloadSSL reloads the SSL configuration.
func (s *SettingsService) ReloadSSL(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssl/reload", nil)
}

// Upgrade upgrades the panel.
func (s *SettingsService) Upgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/upgrade", body)
}

// GetNotesByVersion returns the release notes for a version.
func (s *SettingsService) GetNotesByVersion(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/upgrade/notes", body)
}

// LoadRelease loads the available releases.
func (s *SettingsService) LoadRelease(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/upgrade/releases")
}

// GetUpgradeInfo returns upgrade info.
func (s *SettingsService) GetUpgradeInfo(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/upgrade")
}

// UpdateAppstoreConfig updates the app store configuration.
func (s *SettingsService) UpdateAppstoreConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/apps/store/update", body)
}

// GetAppstoreConfig returns the app store configuration.
func (s *SettingsService) GetAppstoreConfig(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/apps/store/config")
}

// GetMemo returns the user memo.
func (s *SettingsService) GetMemo(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/memo")
}

// UpdateMemo updates the user memo.
func (s *SettingsService) UpdateMemo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/memo", body)
}

// DescriptionSave saves a description tag.
func (s *SettingsService) DescriptionSave(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/description/save", body)
}

// BaseDir returns the panel's base directory.
func (s *SettingsService) BaseDir(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/basedir")
}

// CheckLocalConn tests the local SSH connection.
func (s *SettingsService) CheckLocalConn(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssh/check", body)
}

// LoadLocalConn returns the local SSH connection details.
func (s *SettingsService) LoadLocalConn(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/ssh/conn")
}

// SetDefaultIsConn sets the local host as the default for SSH operations.
func (s *SettingsService) SetDefaultIsConn(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssh/default", body)
}

// SaveLocalConn saves the local SSH connection.
func (s *SettingsService) SaveLocalConn(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssh", body)
}

// CheckLocalConnByInfo tests a local SSH connection by inline info.
func (s *SettingsService) CheckLocalConnByInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/ssh/check/info", body)
}

// GetTerminalAI returns the terminal AI setting info.
func (s *SettingsService) GetTerminalAI(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/terminal/ai/search", body)
}

// UpdateTerminalAI updates the terminal AI setting.
func (s *SettingsService) UpdateTerminalAI(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/terminal/ai/update", body)
}

// GetFileManageAI returns the file manager AI setting.
func (s *SettingsService) GetFileManageAI(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/files/ai/search", body)
}

// UpdateFileManageAI updates the file manager AI setting.
func (s *SettingsService) UpdateFileManageAI(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/files/ai/update", body)
}

// GetFileHistorySetting returns the file-history setting.
func (s *SettingsService) GetFileHistorySetting(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/file-history/search", body)
}

// UpdateFileHistorySetting updates the file-history setting.
func (s *SettingsService) UpdateFileHistorySetting(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/settings/file-history/update", body)
}

// LoadWebsiteDir returns the websites root directory.
func (s *SettingsService) LoadWebsiteDir(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/settings/website/dir")
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
return s.postMap(ctx, "/logs/login", body)
}

// GetOperationLogs returns operation logs.
func (s *LogsService) GetOperationLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/logs/operation", body)
}

// CleanLogs cleans old logs.
func (s *LogsService) CleanLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/logs/clean", body)
}

// GetSystemFiles returns available system log files.
func (s *LogsService) GetSystemFiles(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/logs/system/files")
}

// GetSystemLogStatus returns the system log status.
func (s *LogsService) GetSystemLogStatus(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/logs/system/status")
}

// ReadSystemLog reads a system log.
func (s *LogsService) ReadSystemLog(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/logs/system/read", body)
}

// ListRunningServices lists running services.
func (s *LogsService) ListRunningServices(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/logs/system/services")
}

// PageTasks paginates the task list.
func (s *LogsService) PageTasks(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/logs/tasks/search", body)
}

// ReadTaskLogByLine reads a task log.
func (s *LogsService) ReadTaskLogByLine(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/logs/tasks/read", body)
}

// CountExecutingTasks returns the number of running tasks.
func (s *LogsService) CountExecutingTasks(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/logs/tasks/executing/count")
}

// Call invokes an arbitrary /logs/* endpoint.
func (s *LogsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// GroupsService covers /groups/* (host / app / website group management).
