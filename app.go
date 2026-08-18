package onepanel

import "context"
type AppService struct {
	ServiceBase
}

// Search searches the app store.
func (s *AppService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetApp returns the app metadata by its key.
func (s *AppService) GetApp(ctx context.Context, key string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/"+key, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Detail returns a specific app version's detail.
func (s *AppService) Detail(ctx context.Context, appID, version, kind string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/detail/"+appID+"/"+version+"/"+kind, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DetailForNode returns the per-node version detail.
func (s *AppService) DetailForNode(ctx context.Context, appKey, version string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/detail/node/"+appKey+"/"+version, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DetailByID returns an installed app detail by id.
func (s *AppService) DetailByID(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/details/"+itoa(id), &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Install installs an app.
func (s *AppService) Install(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/install", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Tags returns the available app tags.
func (s *AppService) Tags(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/tags", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Icon returns the icon for an app.
func (s *AppService) Icon(ctx context.Context, key string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/icon/"+key, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SyncRemote syncs the remote app catalog.
func (s *AppService) SyncRemote(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/sync/remote", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SyncLocal syncs the local app cache.
func (s *AppService) SyncLocal(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/sync/local", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckUpdate lists apps that have updates available.
func (s *AppService) CheckUpdate(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/checkupdate", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckInstalled checks whether an app is already installed.
func (s *AppService) CheckInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadPort returns the port allocation of an installed app.
func (s *AppService) LoadPort(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/loadport", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ConnInfo returns the connection info of an installed app.
func (s *AppService) ConnInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/conninfo", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheck performs a pre-delete check for an app installation.
func (s *AppService) DeleteCheck(ctx context.Context, installID uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/installed/delete/check/"+itoa(installID), &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchInstalled searches the installed apps.
func (s *AppService) SearchInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListInstalled lists all installed apps.
func (s *AppService) ListInstalled(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/installed/list", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateInstalled runs a lifecycle op on an installed app.
func (s *AppService) OperateInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/op", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SyncInstalled re-syncs the installed app state.
func (s *AppService) SyncInstalled(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/sync", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChangePort changes the exposed port of an installed app.
func (s *AppService) ChangePort(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/port/change", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Services returns the services managed by an app.
func (s *AppService) Services(ctx context.Context, key string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/services/"+key, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultConfig returns the default config for an installed app.
func (s *AppService) DefaultConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/conf", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Params returns the runtime parameters of an installed app.
func (s *AppService) Params(ctx context.Context, installID uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/apps/installed/params/"+itoa(installID), &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateParams updates the runtime parameters of an installed app.
func (s *AppService) UpdateParams(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/params/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateVersions returns the list of versions an installed app can be updated to.
func (s *AppService) UpdateVersions(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/apps/installed/update/versions", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /apps/* endpoint.
func (s *AppService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteService covers /websites/* (PHP/Static sites, reverse proxies, runtime config).


