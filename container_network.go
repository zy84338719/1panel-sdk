package onepanel

import "context"

// === Network ===

// ListNetworks lists docker networks.
func (s *ContainerService) ListNetworks(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/network")
}

// DeleteNetwork deletes a docker network.
func (s *ContainerService) DeleteNetwork(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/network/del", body)
}

// SearchNetworks searches docker networks.
func (s *ContainerService) SearchNetworks(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/network/search", body)
}

// CreateNetwork creates a docker network.
func (s *ContainerService) CreateNetwork(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/network", body)
}
