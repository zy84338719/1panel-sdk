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
