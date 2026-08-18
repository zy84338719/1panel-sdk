package onepanel

import "context"

// === Volume ===

// ListVolumes lists docker volumes.
func (s *ContainerService) ListVolumes(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/containers/volume")
}

// DeleteVolume deletes a docker volume.
func (s *ContainerService) DeleteVolume(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/volume/del", body)
}

// SearchVolumes searches docker volumes.
func (s *ContainerService) SearchVolumes(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/volume/search", body)
}

// CreateVolume creates a docker volume.
func (s *ContainerService) CreateVolume(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/volume", body)
}

