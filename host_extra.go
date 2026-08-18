package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type SSHService struct {
	ServiceBase
}

// GetSSHInfo returns the SSH config.
func (s *SSHService) GetSSHInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/search", body)
}

// UpdateSSH updates the SSH config.
func (s *SSHService) UpdateSSH(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/update", body)
}

// LoadSSHLogs loads SSH logs.
func (s *SSHService) LoadSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log", body)
}

// ExportSSHLogs exports SSH logs.
func (s *SSHService) ExportSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log/export", body)
}

// CleanSSHLogs cleans SSH logs.
func (s *SSHService) CleanSSHLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/log/clean", body)
}

// Operate runs a SSH op (start/stop/restart).
func (s *SSHService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/operate", body)
}

// LoadSSHFile loads the SSH config file.
func (s *SSHService) LoadSSHFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/file", body)
}

// UpdateSSHByFile replaces the SSH config file.
func (s *SSHService) UpdateSSHByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/file/update", body)
}

// CreateRootCert creates a root CA certificate.
func (s *SSHService) CreateRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert", body)
}

// EditRootCert edits a root CA certificate.
func (s *SSHService) EditRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/update", body)
}

// SyncRootCert syncs the root CA to nodes.
func (s *SSHService) SyncRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/sync", body)
}

// SearchRootCert searches root certificates.
func (s *SSHService) SearchRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/search", body)
}

// DeleteRootCert deletes a root certificate.
func (s *SSHService) DeleteRootCert(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/ssh/cert/delete", body)
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
return s.postMap(ctx, "/hosts/monitor/search", body)
}

// CleanMonitor cleans old monitor data.
func (s *MonitorService) CleanMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/monitor/clean", body)
}

// NetworkOptions returns the network monitor options.
func (s *MonitorService) NetworkOptions(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/hosts/monitor/netoptions")
}

// IOOptions returns the disk-IO monitor options.
func (s *MonitorService) IOOptions(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/hosts/monitor/iooptions")
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

type FirewallService struct {
	ServiceBase
}

// LoadBaseInfo returns the firewall base info.
func (s *FirewallService) LoadBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/base", body)
}

// SearchRule searches firewall rules.
func (s *FirewallService) SearchRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/search", body)
}

// Operate runs a firewall op (start/stop/restart).
func (s *FirewallService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/operate", body)
}

// OperatePortRule creates a port rule.
func (s *FirewallService) OperatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/port", body)
}

// OperateForwardRule creates a forward rule.
func (s *FirewallService) OperateForwardRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/forward", body)
}

// OperateIPRule creates an IP rule.
func (s *FirewallService) OperateIPRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/ip", body)
}

// BatchOperateRule batch-operates rules.
func (s *FirewallService) BatchOperateRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/batch", body)
}

// UpdatePortRule updates a port rule.
func (s *FirewallService) UpdatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/port", body)
}

// UpdateAddrRule updates an address rule.
func (s *FirewallService) UpdateAddrRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/addr", body)
}

// UpdateDescription updates a firewall rule's description.
func (s *FirewallService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/description", body)
}

// SearchFilterRules searches iptables filter rules.
func (s *FirewallService) SearchFilterRules(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/search", body)
}

// OperateFilterRule operates an iptables filter rule.
func (s *FirewallService) OperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/operate", body)
}

// BatchOperateFilterRule batch operates filter rules.
func (s *FirewallService) BatchOperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/batch", body)
}

// OperateFilterChain operates an iptables filter chain.
func (s *FirewallService) OperateFilterChain(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/operate", body)
}

// LoadChainStatus returns the chain status.
func (s *FirewallService) LoadChainStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/chain/status", body)
}

// Call invokes an arbitrary /hosts/firewall/* endpoint.
func (s *FirewallService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// NginxService covers /nginx/* (Nginx config management).
