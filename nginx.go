package onepanel

import "context"

type NginxService struct {
	ServiceBase
}

// LoadConf loads the Nginx configuration.
func (s *NginxService) LoadConf(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/nginx/conf")
}

// UpdateConf updates the Nginx configuration.
func (s *NginxService) UpdateConf(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/nginx/conf", body)
}

// LoadServerInfo loads a server block.
func (s *NginxService) LoadServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/nginx/server", body)
}

// UpdateServerInfo updates a server block.
func (s *NginxService) UpdateServerInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/nginx/server/update", body)
}

// Operate runs an Nginx op (start/stop/reload/...).
func (s *NginxService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/nginx/operate", body)
}

// Call invokes an arbitrary /nginx/* endpoint.
func (s *NginxService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// ProcessService covers /processes/* (process listing, killing).
