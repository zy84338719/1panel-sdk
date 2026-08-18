package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type BackupService struct {
	ServiceBase
}

// CheckBackupUsed checks whether a backup is used.
func (s *BackupService) CheckBackupUsed(ctx context.Context, name string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/backups/check/"+name, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBackupOptions returns the available backup options.
func (s *BackupService) LoadBackupOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/backups/options", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchBackups searches backup records.
func (s *BackupService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetLocalDir returns the local backup directory.
func (s *BackupService) GetLocalDir(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/backups/local", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RefreshToken refreshes the OAuth token of a backup destination.
func (s *RefreshToken) RefreshToken(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/refresh/token", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListBuckets lists buckets on a storage destination.
func (s *BackupService) ListBuckets(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/buckets", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a new backup destination.
func (s *BackupService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckBackup validates the backup destination.
func (s *BackupService) CheckBackup(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/conn/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a backup destination.
func (s *BackupService) Delete(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/del", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a backup destination.
func (s *BackupService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RunBackup triggers a backup job.
func (s *BackupService) RunBackup(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/backup", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Recover restores from a backup.
func (s *BackupService) Recover(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/recover", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UploadForRecover uploads a local backup file to recover.
func (s *BackupService) UploadForRecover(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/upload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RecoverByUpload recovers from an uploaded file.
func (s *BackupService) RecoverByUpload(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/recover/byupload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadFilesFromBackup browses files inside a backup.
func (s *BackupService) LoadFilesFromBackup(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/search/files", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRecords searches backup records.
func (s *BackupService) SearchRecords(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRecordSize returns the total size of records.
func (s *BackupService) LoadRecordSize(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/size", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRecordsByCronjob searches records generated by a cron job.
func (s *BackupService) SearchRecordsByCronjob(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/search/bycronjob", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadRecord downloads a backup record.
func (s *BackupService) DownloadRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/download", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteRecord deletes a backup record.
func (s *BackupService) DeleteRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRecordDescription updates a record's description.
func (s *BackupService) UpdateRecordDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/backups/record/description/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /backups/* endpoint.
func (s *BackupService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// RefreshToken is exposed for the rare case where the caller wants to refresh a
// destination's OAuth token without touching the BackupService helpers.
type RefreshToken struct{ ServiceBase }

// CronjobService covers /cronjobs/* (scheduled tasks).
type CronjobService struct {
	ServiceBase
}

// Create creates a cron job.
func (s *CronjobService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadNextHandle previews the next fire time.
func (s *CronjobService) LoadNextHandle(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/next", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Import imports a cron job.
func (s *CronjobService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/import", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Export exports cron jobs.
func (s *CronjobService) Export(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/export", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadCronjobInfo loads a cron job's info.
func (s *CronjobService) LoadCronjobInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/load/info", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadScriptOptions lists the available scripts.
func (s *CronjobService) LoadScriptOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/cronjobs/script/options", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a cron job.
func (s *CronjobService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Stop stops a running cron job.
func (s *CronjobService) Stop(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/stop", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a cron job.
func (s *CronjobService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateGroup updates the group of a cron job.
func (s *CronjobService) UpdateGroup(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/group/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateStatus updates the enable/disable status.
func (s *CronjobService) UpdateStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// HandleOnce runs the job once.
func (s *CronjobService) HandleOnce(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/handle", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches cron jobs.
func (s *CronjobService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRecords searches job execution records.
func (s *CronjobService) SearchRecords(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/search/records", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadRecordLog loads a record's log.
func (s *CronjobService) LoadRecordLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/records/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanRecord cleans old records.
func (s *CronjobService) CleanRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/cronjobs/records/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /cronjobs/* endpoint.
func (s *CronjobService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FileService covers /files/* (browse, upload, download, edit, compress, share).
type FileService struct {
	ServiceBase
}

// List lists files at a path.
func (s *FileService) List(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadDirs loads the available directories (quick navigation).
func (s *FileService) LoadDirs(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/files/dir", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadFileContent reads the content of a file.
func (s *FileService) LoadFileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/content", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SaveFileContent writes content to a file.
func (s *FileService) SaveFileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/save", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Upload uploads a file (multipart).
func (s *FileService) Upload(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/upload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a new file or directory.
func (s *FileService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a file or directory.
func (s *FileService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchOperate batch-deletes files.
func (s *FileService) BatchOperate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/batch/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Compress compresses files.
func (s *FileService) Compress(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/compress", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Decompress decompresses an archive.
func (s *FileService) Decompress(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/decompress", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangePermissions chmods a file.
func (s *FileService) ChangePermissions(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/permission", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeOwner chowns a file.
func (s *FileService) ChangeOwner(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/owner", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Move moves a file.
func (s *FileService) Move(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/move", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Rename renames a file.
func (s *FileService) Rename(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/rename", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadURL returns the download endpoint for files.
func (s *FileService) DownloadURL() string { return "/files/download" }

// Search performs a deep file search.
func (s *FileService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/grep", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Wget downloads a file via HTTP.
func (s *FileService) Wget(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/wget", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Token issues a download token for sharing.
func (s *FileService) Token(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/token", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkUploadInit starts a chunked upload.
func (s *FileService) ChunkUploadInit(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/chunkupload/init", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkUpload appends a chunk.
func (s *FileService) ChunkUpload(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/chunkupload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkMerge merges uploaded chunks.
func (s *FileService) ChunkMerge(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/chunkupload/merge", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// FileHistory searches file-edit history.
func (s *FileService) FileHistory(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/history", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RecoverFile recovers a deleted file from the recycle bin.
func (s *FileService) RecoverFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/files/recover", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /files/* endpoint.
func (s *FileService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SettingsService covers /settings/* (panel configuration).
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
type ToolboxService struct {
	ServiceBase
}

// DeviceBaseInfo returns device base info.
func (s *ToolboxService) DeviceBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/base", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Users returns the device users.
func (s *ToolboxService) Users(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/toolbox/device/users", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeOption returns the time-zone options.
func (s *ToolboxService) TimeOption(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/toolbox/device/zone/options", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDeviceConf updates the device config.
func (s *ToolboxService) UpdateDeviceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/update/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDeviceHost updates the hostname.
func (s *ToolboxService) UpdateDeviceHost(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/update/host", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDevicePasswd updates the device password.
func (s *ToolboxService) UpdateDevicePasswd(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/update/passwd", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDeviceSwap updates the device swap.
func (s *ToolboxService) UpdateDeviceSwap(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/update/swap", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDeviceByFile updates the device by file.
func (s *ToolboxService) UpdateDeviceByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/update/byconf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckDNS checks DNS.
func (s *ToolboxService) CheckDNS(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/check/dns", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceConf returns the device conf.
func (s *ToolboxService) DeviceConf(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/device/conf", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ScanSystem scans the system.
func (s *ToolboxService) ScanSystem(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/scan", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SystemClean cleans the system.
func (s *ToolboxService) SystemClean(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Fail2BanBase returns fail2ban base info.
func (s *ToolboxService) Fail2BanBase(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/toolbox/fail2ban/base", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadFail2BanConf loads fail2ban config.
func (s *ToolboxService) LoadFail2BanConf(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/toolbox/fail2ban/load/conf", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchFail2Ban searches fail2ban records.
func (s *ToolboxService) SearchFail2Ban(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/fail2ban/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateFail2Ban runs a fail2ban operation.
func (s *ToolboxService) OperateFail2Ban(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/fail2ban/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateSSHD runs an operation on SSHD.
func (s *ToolboxService) OperateSSHD(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/fail2ban/operate/sshd", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateFail2BanConf updates fail2ban config.
func (s *ToolboxService) UpdateFail2BanConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/fail2ban/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateFail2BanConfByFile updates fail2ban config by file.
func (s *ToolboxService) UpdateFail2BanConfByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/fail2ban/update/byconf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// FtpBase returns FTP base info.
func (s *ToolboxService) FtpBase(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/toolbox/ftp/base", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadFtpLogInfo loads FTP log info.
func (s *ToolboxService) LoadFtpLogInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/log/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateFtp operates FTP.
func (s *ToolboxService) OperateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchFtp searches FTP.
func (s *ToolboxService) SearchFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateFtp creates an FTP account.
func (s *ToolboxService) CreateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateFtp updates an FTP account.
func (s *ToolboxService) UpdateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteFtp deletes an FTP account.
func (s *ToolboxService) DeleteFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SyncFtp syncs the FTP server.
func (s *ToolboxService) SyncFtp(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/ftp/sync", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchClam searches clamav tasks.
func (s *ToolboxService) SearchClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchClamRecord searches clamav records.
func (s *ToolboxService) SearchClamRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/record/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanClamRecord cleans clamav records.
func (s *ToolboxService) CleanClamRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/record/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchClamFile searches clamav file scan history.
func (s *ToolboxService) SearchClamFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/file/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateClamFile updates a clamav scanned file action.
func (s *ToolboxService) UpdateClamFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/file/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateClam creates a clamav scan task.
func (s *ToolboxService) CreateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadClamBaseInfo loads clamav base info.
func (s *ToolboxService) LoadClamBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/base", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateClam runs a clamav op.
func (s *ToolboxService) OperateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateClam updates a clamav task.
func (s *ToolboxService) UpdateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateClamStatus updates the clamav status.
func (s *ToolboxService) UpdateClamStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/status/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteClam deletes a clamav task.
func (s *ToolboxService) DeleteClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// HandleClamScan handles a clamav scan result.
func (s *ToolboxService) HandleClamScan(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/toolbox/clam/handle", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /toolbox/* endpoint.
func (s *ToolboxService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AlertsService covers /alert/* (alert rules, channels, logs).
type AlertsService struct {
	ServiceBase
}

// Create creates an alert.
func (s *AlertsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates an alert.
func (s *AlertsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches alerts.
func (s *AlertsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateStatus updates an alert's status.
func (s *AlertsService) UpdateStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an alert.
func (s *AlertsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DisksList lists disks for alerts.
func (s *AlertsService) DisksList(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/alert/disks/list", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchLogs searches alert logs.
func (s *AlertsService) SearchLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/logs/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanLogs cleans alert logs.
func (s *AlertsService) CleanLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/logs/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ClamsList lists installed clamav instances.
func (s *AlertsService) ClamsList(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/alert/clams/list", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CronjobList lists cronjobs for alerts.
func (s *AlertsService) CronjobList(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/cronjob/list", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateConfig updates the alert config.
func (s *AlertsService) UpdateConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/config/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetConfig returns the alert config.
func (s *AlertsService) GetConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/config/info", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchConfigs searches alert configs.
func (s *AlertsService) SearchConfigs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/config/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteConfig deletes an alert config.
func (s *AlertsService) DeleteConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/config/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TestConfig tests the alert config.
func (s *AlertsService) TestConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/alert/config/test", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /alert/* endpoint.
func (s *AlertsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AIService covers /ai/* (Ollama models, GPU, MCP servers, TensorRT-LLM, agents, accounts).
type AIService struct {
	ServiceBase
}

// OllamaClose closes an Ollama model.
func (s *AIService) OllamaClose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/close", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaCreate creates an Ollama model.
func (s *AIService) OllamaCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaRecreate recreates an Ollama model.
func (s *AIService) OllamaRecreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/recreate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaSearch searches Ollama models.
func (s *AIService) OllamaSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaSync syncs Ollama models.
func (s *AIService) OllamaSync(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/sync", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaLoadDetail loads Ollama model details.
func (s *AIService) OllamaLoadDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaDelete deletes an Ollama model.
func (s *AIService) OllamaDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUInfo returns the GPU info.
func (s *AIService) GPUInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/gpu/load", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUMonitor loads the GPU monitor data.
func (s *AIService) GPUMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/gpu/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUOptions returns the GPU monitoring options.
func (s *AIService) GPUOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/gpu/options", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpSearch searches MCP servers.
func (s *AIService) McpSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerDetail returns an MCP server's detail.
func (s *AIService) McpServerDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/detail", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerCreate creates an MCP server.
func (s *AIService) McpServerCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerUpdate updates an MCP server.
func (s *AIService) McpServerUpdate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerDelete deletes an MCP server.
func (s *AIService) McpServerDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerOperate operates an MCP server.
func (s *AIService) McpServerOperate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/op", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /ai/* endpoint.
func (s *AIService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AgentsService covers /agents/* (managed agent instances).
type AgentsService struct {
	ServiceBase
}

// Create creates an agent.
func (s *AgentsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchInstall installs agents in batch.
func (s *AgentsService) BatchInstall(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/batch/install", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchUpgrade upgrades agents in batch.
func (s *AgentsService) BatchUpgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/batch/upgrade", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches agents.
func (s *AgentsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheck pre-flight check for agent deletion.
func (s *AgentsService) DeleteCheck(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/delete/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an agent.
func (s *AgentsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/delete", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResetToken resets an agent's access token.
func (s *AgentsService) ResetToken(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/token/reset", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Overview returns the agent overview.
func (s *AgentsService) Overview(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/overview", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Remark updates an agent's remark.
func (s *AgentsService) Remark(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/agents/remark", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Providers returns the available agent providers.
func (s *AgentsService) Providers(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/accounts/providers", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAccount creates an agent account.
func (s *AgentsService) CreateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAccount updates an agent account.
func (s *AgentsService) UpdateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchAccounts searches agent accounts.
func (s *AgentsService) SearchAccounts(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteAccount deletes an agent account.
func (s *AgentsService) DeleteAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/delete", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyAccount verifies an agent account.
func (s *AgentsService) VerifyAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/verify", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /agents/* or /ai/accounts/* endpoint.
func (s *AgentsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SSHService covers /hosts/ssh/* (SSH config, logs, root cert, host SSH).
type SSHService struct {
	ServiceBase
}

// GetSSHInfo returns the SSH config.
func (s *SSHService) GetSSHInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateSSH updates the SSH config.
func (s *SSHService) UpdateSSH(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadSSHLogs loads SSH logs.
func (s *SSHService) LoadSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ExportSSHLogs exports SSH logs.
func (s *SSHService) ExportSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/log/export", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanSSHLogs cleans SSH logs.
func (s *SSHService) CleanSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/log/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Operate runs a SSH op (start/stop/restart).
func (s *SSHService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadSSHFile loads the SSH config file.
func (s *SSHService) LoadSSHFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/file", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateSSHByFile replaces the SSH config file.
func (s *SSHService) UpdateSSHByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/file/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateRootCert creates a root CA certificate.
func (s *SSHService) CreateRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/cert", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// EditRootCert edits a root CA certificate.
func (s *SSHService) EditRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/cert/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SyncRootCert syncs the root CA to nodes.
func (s *SSHService) SyncRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/cert/sync", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRootCert searches root certificates.
func (s *SSHService) SearchRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/cert/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteRootCert deletes a root certificate.
func (s *SSHService) DeleteRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/ssh/cert/delete", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /hosts/ssh/* endpoint.
func (s *SSHService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// MonitorService covers /hosts/monitor/* (CPU / memory / disk / network monitor).
type MonitorService struct {
	ServiceBase
}

// LoadMonitor loads the monitor data.
func (s *MonitorService) LoadMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/monitor/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanMonitor cleans old monitor data.
func (s *MonitorService) CleanMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/monitor/clean", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkOptions returns the network monitor options.
func (s *MonitorService) NetworkOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/monitor/netoptions", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// IOOptions returns the disk-IO monitor options.
func (s *MonitorService) IOOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/monitor/iooptions", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadMonitorSetting returns the monitor configuration.
func (s *MonitorService) LoadMonitorSetting(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/monitor/setting", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateMonitorSetting updates the monitor configuration.
func (s *MonitorService) UpdateMonitorSetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/monitor/setting/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /hosts/monitor/* endpoint.
func (s *MonitorService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FirewallService covers /hosts/firewall/* (firewall rules and chains).
type FirewallService struct {
	ServiceBase
}

// LoadBaseInfo returns the firewall base info.
func (s *FirewallService) LoadBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/base", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRule searches firewall rules.
func (s *FirewallService) SearchRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Operate runs a firewall op (start/stop/restart).
func (s *FirewallService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperatePortRule creates a port rule.
func (s *FirewallService) OperatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/port", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateForwardRule creates a forward rule.
func (s *FirewallService) OperateForwardRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/forward", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateIPRule creates an IP rule.
func (s *FirewallService) OperateIPRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/ip", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchOperateRule batch-operates rules.
func (s *FirewallService) BatchOperateRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/batch", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatePortRule updates a port rule.
func (s *FirewallService) UpdatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/update/port", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAddrRule updates an address rule.
func (s *FirewallService) UpdateAddrRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/update/addr", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDescription updates a firewall rule's description.
func (s *FirewallService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/update/description", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchFilterRules searches iptables filter rules.
func (s *FirewallService) SearchFilterRules(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/filter/rule/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateFilterRule operates an iptables filter rule.
func (s *FirewallService) OperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/filter/rule/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchOperateFilterRule batch operates filter rules.
func (s *FirewallService) BatchOperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/filter/rule/batch", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateFilterChain operates an iptables filter chain.
func (s *FirewallService) OperateFilterChain(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/filter/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadChainStatus returns the chain status.
func (s *FirewallService) LoadChainStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/firewall/filter/chain/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /hosts/firewall/* endpoint.
func (s *FirewallService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// NginxService covers /nginx/* (Nginx config management).
type NginxService struct {
	ServiceBase
}

// LoadConf loads the Nginx configuration.
func (s *NginxService) LoadConf(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/nginx/conf", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateConf updates the Nginx configuration.
func (s *NginxService) UpdateConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/nginx/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServerInfo loads a server block.
func (s *NginxService) LoadServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/nginx/server", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServerInfo updates a server block.
func (s *NginxService) UpdateServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/nginx/server/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Operate runs an Nginx op (start/stop/reload/...).
func (s *NginxService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/nginx/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/processes", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Kill terminates a process.
func (s *ProcessService) Kill(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/processes/kill", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Get(ctx, "/hosts/diagnostics/summary", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Goroutines returns the goroutine stack dump.
func (s *RuntimeService) Goroutines(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/hosts/diagnostics/goroutines", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateProfile starts a profile capture.
func (s *RuntimeService) CreateProfile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/hosts/diagnostics/profiles", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /hosts/diagnostics/* endpoint.
func (s *RuntimeService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SnapshotService covers /settings/snapshot/* (panel snapshots).
type SnapshotService struct {
	ServiceBase
}

// LoadSnapshotData loads a snapshot file.
func (s *SnapshotService) LoadSnapshotData(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/settings/snapshot/load", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a snapshot.
func (s *SnapshotService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Recreate recreates a snapshot from a previous state.
func (s *SnapshotService) Recreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/recreate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches snapshots.
func (s *SnapshotService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Import imports a snapshot.
func (s *SnapshotService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/import", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a snapshot.
func (s *SnapshotService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Recover recovers a snapshot.
func (s *SnapshotService) Recover(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/recover", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Rollback rolls back to a snapshot.
func (s *SnapshotService) Rollback(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/rollback", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDescription updates a snapshot's description.
func (s *SnapshotService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/settings/snapshot/description/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /settings/snapshot/* endpoint.
func (s *SnapshotService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FavoriteService covers /favorites/* (user favorites/bookmarks).
type FavoriteService struct {
	ServiceBase
}

// Create creates a favorite.
func (s *FavoriteService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/favorites", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a favorite.
func (s *FavoriteService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/favorites/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches favorites.
func (s *FavoriteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/favorites/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /favorites/* endpoint.
func (s *FavoriteService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// TaskService covers /tasks/* (long-running task progress queries).
type TaskService struct {
	ServiceBase
}

// SSLService is a thin alias for WebsiteSSLService kept for naming symmetry
// with the rest of the SDK. It exposes the same /websites/ssl/* endpoints.
type SSLService = WebsiteSSLService

// GetTask returns a task by id.

// GetTask returns a task by id.
func (s *TaskService) GetTask(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListTasks lists tasks.
func (s *TaskService) ListTasks(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CancelTask cancels a running task.
func (s *TaskService) CancelTask(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/cancel", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TaskLog returns the log of a task.
func (s *TaskService) TaskLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/tasks/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /tasks/* endpoint.
func (s *TaskService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
