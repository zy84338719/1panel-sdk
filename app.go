package onepanel

import "context"

type AppService struct {
	ServiceBase
}

// Search searches the app store.
func (s *AppService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/search", body)
}

// GetApp returns the app metadata by its key.
func (s *AppService) GetApp(ctx context.Context, key string) (map[string]any, error) {
	return s.getMap(ctx, "/apps/"+key)
}

// Detail returns a specific app version's detail.
func (s *AppService) Detail(ctx context.Context, appID, version, kind string) (map[string]any, error) {
	return s.getMap(ctx, "/apps/detail/"+appID+"/"+version+"/"+kind)
}

// DetailForNode returns the per-node version detail.
func (s *AppService) DetailForNode(ctx context.Context, appKey, version string) (map[string]any, error) {
	return s.getMap(ctx, "/apps/detail/node/"+appKey+"/"+version)
}

// DetailByID returns an installed app detail by id.
func (s *AppService) DetailByID(ctx context.Context, id uint) (map[string]any, error) {
	return s.getMap(ctx, "/apps/details/"+itoa(id))
}

// Install installs an app.
func (s *AppService) Install(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/install", body)
}

// Tags returns the available app tags.
func (s *AppService) Tags(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/apps/tags")
}

// Icon returns the icon for an app.
func (s *AppService) Icon(ctx context.Context, key string) (map[string]any, error) {
	return s.getMap(ctx, "/apps/icon/"+key)
}

// SyncRemote syncs the remote app catalog.
func (s *AppService) SyncRemote(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/apps/sync/remote", nil)
}

// SyncLocal syncs the local app cache.
func (s *AppService) SyncLocal(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/apps/sync/local", nil)
}

// CheckUpdate lists apps that have updates available.
func (s *AppService) CheckUpdate(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/apps/checkupdate")
}

// CheckInstalled checks whether an app is already installed.
func (s *AppService) CheckInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/check", body)
}

// LoadPort returns the port allocation of an installed app.
func (s *AppService) LoadPort(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/loadport", body)
}

// ConnInfo returns the connection info of an installed app.
func (s *AppService) ConnInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/conninfo", body)
}

// DeleteCheck performs a pre-delete check for an app installation.
func (s *AppService) DeleteCheck(ctx context.Context, installID uint) (map[string]any, error) {
	return s.getMap(ctx, "/apps/installed/delete/check/"+itoa(installID))
}

// SearchInstalled searches the installed apps.
func (s *AppService) SearchInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/search", body)
}

// ListInstalled lists all installed apps.
func (s *AppService) ListInstalled(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/apps/installed/list")
}

// OperateInstalled runs a lifecycle op on an installed app.
func (s *AppService) OperateInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/op", body)
}

// SyncInstalled re-syncs the installed app state.
func (s *AppService) SyncInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/sync", body)
}

// ChangePort changes the exposed port of an installed app.
func (s *AppService) ChangePort(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/port/change", body)
}

// Services returns the services managed by an app.
func (s *AppService) Services(ctx context.Context, key string) (map[string]any, error) {
	return s.getMap(ctx, "/apps/services/"+key)
}

// DefaultConfig returns the default config for an installed app.
func (s *AppService) DefaultConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/conf", body)
}

// Params returns the runtime parameters of an installed app.
func (s *AppService) Params(ctx context.Context, installID uint) (map[string]any, error) {
	return s.getMap(ctx, "/apps/installed/params/"+itoa(installID))
}

// UpdateParams updates the runtime parameters of an installed app.
func (s *AppService) UpdateParams(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/params/update", body)
}

// UpdateVersions returns the list of versions an installed app can be updated to.
func (s *AppService) UpdateVersions(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/apps/installed/update/versions", body)
}

// Call invokes an arbitrary /apps/* endpoint.
func (s *AppService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteService covers /websites/* (PHP/Static sites, reverse proxies, runtime config).
