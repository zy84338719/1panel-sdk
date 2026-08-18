package onepanel

import "context"

// ContainerService covers /containers/* (docker container, image, network, volume, compose, repo).
type ContainerService struct {
	ServiceBase
}

// === Compose ===

// SearchCompose searches docker-compose projects.
func (s *ContainerService) SearchCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/search", body)
}

// CreateCompose creates a new compose project.
func (s *ContainerService) CreateCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose", body)
}

// ComposeEnv loads the env file of a compose project.
func (s *ContainerService) ComposeEnv(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/env", body)
}

// TestCompose validates a compose file syntactically.
func (s *ContainerService) TestCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/test", body)
}

// OperatorCompose runs an operation (up/down/restart/...).
func (s *ContainerService) OperatorCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/operate", body)
}

// CleanComposeLog cleans the compose log file.
func (s *ContainerService) CleanComposeLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/clean/log", body)
}

// UpdateCompose updates a compose project.
func (s *ContainerService) UpdateCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/update", body)
}

// PinCompose pins/unpins a compose project.
func (s *ContainerService) PinCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/compose/pin", body)
}
