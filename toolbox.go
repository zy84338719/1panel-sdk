package onepanel

import "context"

type ToolboxService struct {
	ServiceBase
}

// DeviceBaseInfo returns device base info.
func (s *ToolboxService) DeviceBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/base", body)
}

// Users returns the device users.
func (s *ToolboxService) Users(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/toolbox/device/users")
}

// TimeOption returns the time-zone options.
func (s *ToolboxService) TimeOption(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/toolbox/device/zone/options")
}

// UpdateDeviceConf updates the device config.
func (s *ToolboxService) UpdateDeviceConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/update/conf", body)
}

// UpdateDeviceHost updates the hostname.
func (s *ToolboxService) UpdateDeviceHost(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/update/host", body)
}

// UpdateDevicePasswd updates the device password.
func (s *ToolboxService) UpdateDevicePasswd(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/update/passwd", body)
}

// UpdateDeviceSwap updates the device swap.
func (s *ToolboxService) UpdateDeviceSwap(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/update/swap", body)
}

// UpdateDeviceByFile updates the device by file.
func (s *ToolboxService) UpdateDeviceByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/update/byconf", body)
}

// CheckDNS checks DNS.
func (s *ToolboxService) CheckDNS(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/check/dns", body)
}

// DeviceConf returns the device conf.
func (s *ToolboxService) DeviceConf(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/device/conf", nil)
}

// ScanSystem scans the system.
func (s *ToolboxService) ScanSystem(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/scan", body)
}

// SystemClean cleans the system.
func (s *ToolboxService) SystemClean(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clean", body)
}

// Fail2BanBase returns fail2ban base info.
func (s *ToolboxService) Fail2BanBase(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/toolbox/fail2ban/base")
}

// LoadFail2BanConf loads fail2ban config.
func (s *ToolboxService) LoadFail2BanConf(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/toolbox/fail2ban/load/conf")
}

// SearchFail2Ban searches fail2ban records.
func (s *ToolboxService) SearchFail2Ban(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/fail2ban/search", body)
}

// OperateFail2Ban runs a fail2ban operation.
func (s *ToolboxService) OperateFail2Ban(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/fail2ban/operate", body)
}

// OperateSSHD runs an operation on SSHD.
func (s *ToolboxService) OperateSSHD(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/fail2ban/operate/sshd", body)
}

// UpdateFail2BanConf updates fail2ban config.
func (s *ToolboxService) UpdateFail2BanConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/fail2ban/update", body)
}

// UpdateFail2BanConfByFile updates fail2ban config by file.
func (s *ToolboxService) UpdateFail2BanConfByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/fail2ban/update/byconf", body)
}

// FtpBase returns FTP base info.
func (s *ToolboxService) FtpBase(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/toolbox/ftp/base")
}

// LoadFtpLogInfo loads FTP log info.
func (s *ToolboxService) LoadFtpLogInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/log/search", body)
}

// OperateFtp operates FTP.
func (s *ToolboxService) OperateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/operate", body)
}

// SearchFtp searches FTP.
func (s *ToolboxService) SearchFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/search", body)
}

// CreateFtp creates an FTP account.
func (s *ToolboxService) CreateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp", body)
}

// UpdateFtp updates an FTP account.
func (s *ToolboxService) UpdateFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/update", body)
}

// DeleteFtp deletes an FTP account.
func (s *ToolboxService) DeleteFtp(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/del", body)
}

// SyncFtp syncs the FTP server.
func (s *ToolboxService) SyncFtp(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/ftp/sync", nil)
}

// SearchClam searches clamav tasks.
func (s *ToolboxService) SearchClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/search", body)
}

// SearchClamRecord searches clamav records.
func (s *ToolboxService) SearchClamRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/record/search", body)
}

// CleanClamRecord cleans clamav records.
func (s *ToolboxService) CleanClamRecord(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/record/clean", body)
}

// SearchClamFile searches clamav file scan history.
func (s *ToolboxService) SearchClamFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/file/search", body)
}

// UpdateClamFile updates a clamav scanned file action.
func (s *ToolboxService) UpdateClamFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/file/update", body)
}

// CreateClam creates a clamav scan task.
func (s *ToolboxService) CreateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam", body)
}

// LoadClamBaseInfo loads clamav base info.
func (s *ToolboxService) LoadClamBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/base", body)
}

// OperateClam runs a clamav op.
func (s *ToolboxService) OperateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/operate", body)
}

// UpdateClam updates a clamav task.
func (s *ToolboxService) UpdateClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/update", body)
}

// UpdateClamStatus updates the clamav status.
func (s *ToolboxService) UpdateClamStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/status/update", body)
}

// DeleteClam deletes a clamav task.
func (s *ToolboxService) DeleteClam(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/del", body)
}

// HandleClamScan handles a clamav scan result.
func (s *ToolboxService) HandleClamScan(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/toolbox/clam/handle", body)
}

// Call invokes an arbitrary /toolbox/* endpoint.
func (s *ToolboxService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AlertsService covers /alert/* (alert rules, channels, logs).
