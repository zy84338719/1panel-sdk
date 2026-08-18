package onepanel

import "context"


type AlertsService struct {
	ServiceBase
}

// Create creates an alert.
func (s *AlertsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert", body)
}

// Update updates an alert.
func (s *AlertsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/update", body)
}

// Search searches alerts.
func (s *AlertsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/search", body)
}

// UpdateStatus updates an alert's status.
func (s *AlertsService) UpdateStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/status", body)
}

// Delete deletes an alert.
func (s *AlertsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/del", body)
}

// DisksList lists disks for alerts.
func (s *AlertsService) DisksList(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/alert/disks/list")
}

// SearchLogs searches alert logs.
func (s *AlertsService) SearchLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/logs/search", body)
}

// CleanLogs cleans alert logs.
func (s *AlertsService) CleanLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/logs/clean", body)
}

// ClamsList lists installed clamav instances.
func (s *AlertsService) ClamsList(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/alert/clams/list")
}

// CronjobList lists cronjobs for alerts.
func (s *AlertsService) CronjobList(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/cronjob/list", body)
}

// UpdateConfig updates the alert config.
func (s *AlertsService) UpdateConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/config/update", body)
}

// GetConfig returns the alert config.
func (s *AlertsService) GetConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/config/info", body)
}

// SearchConfigs searches alert configs.
func (s *AlertsService) SearchConfigs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/config/search", body)
}

// DeleteConfig deletes an alert config.
func (s *AlertsService) DeleteConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/config/del", body)
}

// TestConfig tests the alert config.
func (s *AlertsService) TestConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/alert/config/test", body)
}

// Call invokes an arbitrary /alert/* endpoint.
func (s *AlertsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AIService covers /ai/* (Ollama models, GPU, MCP servers, TensorRT-LLM, agents, accounts).
