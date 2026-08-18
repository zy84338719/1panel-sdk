package onepanel

import "context"

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
