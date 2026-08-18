package onepanel

import "context"


type CronjobService struct {
	ServiceBase
}

// Create creates a cron job.
func (s *CronjobService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs", body)
}

// LoadNextHandle previews the next fire time.
func (s *CronjobService) LoadNextHandle(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/next", body)
}

// Import imports a cron job.
func (s *CronjobService) Import(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/import", body)
}

// Export exports cron jobs.
func (s *CronjobService) Export(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/export", body)
}

// LoadCronjobInfo loads a cron job's info.
func (s *CronjobService) LoadCronjobInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/load/info", body)
}

// LoadScriptOptions lists the available scripts.
func (s *CronjobService) LoadScriptOptions(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/cronjobs/script/options")
}

// Delete deletes a cron job.
func (s *CronjobService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/del", body)
}

// Stop stops a running cron job.
func (s *CronjobService) Stop(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/stop", body)
}

// Update updates a cron job.
func (s *CronjobService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/update", body)
}

// UpdateGroup updates the group of a cron job.
func (s *CronjobService) UpdateGroup(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/group/update", body)
}

// UpdateStatus updates the enable/disable status.
func (s *CronjobService) UpdateStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/status", body)
}

// HandleOnce runs the job once.
func (s *CronjobService) HandleOnce(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/handle", body)
}

// Search searches cron jobs.
func (s *CronjobService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/search", body)
}

// SearchRecords searches job execution records.
func (s *CronjobService) SearchRecords(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/search/records", body)
}

// LoadRecordLog loads a record's log.
func (s *CronjobService) LoadRecordLog(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/records/log", body)
}

// CleanRecord cleans old records.
func (s *CronjobService) CleanRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/cronjobs/records/clean", body)
}

// Call invokes an arbitrary /cronjobs/* endpoint.
func (s *CronjobService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FileService covers /files/* (browse, upload, download, edit, compress, share).
