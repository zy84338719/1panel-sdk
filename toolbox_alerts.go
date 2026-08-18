package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
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
