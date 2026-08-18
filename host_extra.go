package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
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
