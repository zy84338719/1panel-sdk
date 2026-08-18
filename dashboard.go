package onepanel

import (
	"context"
	"net/http"
)

// DashboardService covers /dashboard/* APIs.
type DashboardService struct {
	ServiceBase
}

// OSInfo returns the host OS / kernel / architecture summary.
func (s *DashboardService) OSInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/base/os", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// QuickOptions returns the dashboard quick-jump menu options.
func (s *DashboardService) QuickOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/quick/option", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateQuickJump saves the user's quick-jump selection.
func (s *DashboardService) UpdateQuickJump(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/dashboard/quick/change", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AppLauncher returns the home page app launcher list.
func (s *DashboardService) AppLauncher(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/app/launcher", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAppLauncher toggles a launcher item.
func (s *DashboardService) UpdateAppLauncher(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/dashboard/app/launcher/show", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AppLauncherOption returns the list of available launcher items.
func (s *DashboardService) AppLauncherOption(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/dashboard/app/launcher/option", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseInfo returns aggregate dashboard data with the given IO/Net options.
func (s *DashboardService) BaseInfo(ctx context.Context, ioOption, netOption string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/base/"+ioOption+"/"+netOption, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentNodeInfo returns info about the current node without sampling.
func (s *DashboardService) CurrentNodeInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/current/node", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentInfo returns the live dashboard snapshot.
func (s *DashboardService) CurrentInfo(ctx context.Context, ioOption, netOption string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/current/"+ioOption+"/"+netOption, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TopCPU returns the top CPU-consuming processes.
func (s *DashboardService) TopCPU(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/current/top/cpu", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TopMem returns the top memory-consuming processes.
func (s *DashboardService) TopMem(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/dashboard/current/top/mem", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SystemRestart issues a system restart / shutdown. operation = "restart" | "shutdown".
func (s *DashboardService) SystemRestart(ctx context.Context, operation string) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/dashboard/system/restart/"+operation, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /dashboard/* endpoint.
func (s *DashboardService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// silence unused import
var _ = http.MethodGet
