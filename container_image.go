package onepanel

import "context"

// === Image ===

// ListImages lists images.
func (s *ContainerService) ListImages(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/image")
}

// ListAllImages lists all images including dangling.
func (s *ContainerService) ListAllImages(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/image/all")
}

// SearchImages searches local images.
func (s *ContainerService) SearchImages(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/search", body)
}

// PullImage pulls an image from a registry.
func (s *ContainerService) PullImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/pull", body)
}

// PushImage pushes an image to a registry.
func (s *ContainerService) PushImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/push", body)
}

// SaveImage saves an image to a tarball.
func (s *ContainerService) SaveImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/save", body)
}

// LoadImage loads an image from a tarball.
func (s *ContainerService) LoadImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/load", body)
}

// RemoveImage removes an image.
func (s *ContainerService) RemoveImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/remove", body)
}

// TagImage tags an image.
func (s *ContainerService) TagImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/tag", body)
}

// BuildImage builds an image from a dockerfile.
func (s *ContainerService) BuildImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/image/build", body)
}
