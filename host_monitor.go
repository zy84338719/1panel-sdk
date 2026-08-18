package onepanel

import "context"

type MonitorService struct {
	ServiceBase
}

// LoadMonitor loads the monitor data.
func (s *MonitorService) LoadMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/hosts/monitor/search", body)
}

// CleanMonitor cleans old monitor data.
func (s *MonitorService) CleanMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/hosts/monitor/clean", body)
}

// NetworkOptions returns the network monitor options as a JSON array.
func (s *MonitorService) NetworkOptions(ctx context.Context) ([]any, error) {
	return s.GetList(ctx, "/hosts/monitor/netoptions")
}

// IOOptions returns the disk-IO monitor options as a JSON array.
func (s *MonitorService) IOOptions(ctx context.Context) ([]any, error) {
	return s.GetList(ctx, "/hosts/monitor/iooptions")
}

// LoadMonitorSetting returns the monitor configuration.
func (s *MonitorService) LoadMonitorSetting(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/hosts/monitor/setting")
}

// UpdateMonitorSetting updates the monitor configuration.
func (s *MonitorService) UpdateMonitorSetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/hosts/monitor/setting/update", body)
}

// Call invokes an arbitrary /hosts/monitor/* endpoint.
func (s *MonitorService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// FirewallService covers /hosts/firewall/* (firewall rules and chains).
