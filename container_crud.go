package onepanel

import "context"

// === Container CRUD ===

// Stats returns the live stats of a container by id.
func (s *ContainerService) Stats(ctx context.Context, id string) (map[string]any, error) {
	return s.getMap(ctx, "/containers/stats/"+id)
}

// Create creates a new container.
func (s *ContainerService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers", body)
}

// Update updates a container (restart policy, name, env, etc).
func (s *ContainerService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/update", body)
}

// Upgrade rebuilds a container with a new image.
func (s *ContainerService) Upgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/upgrade", body)
}

// Info returns detailed info for a container.
func (s *ContainerService) Info(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/info", body)
}

// Search searches containers.
func (s *ContainerService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/search", body)
}

// List lists containers (lightweight).
func (s *ContainerService) List(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/containers/list", map[string]any{})
}

// ListByImage lists containers using a given image.
func (s *ContainerService) ListByImage(ctx context.Context, name string) (map[string]any, error) {
	return s.postMap(ctx, "/containers/list/byimage", map[string]string{"name": name})
}

// Status returns aggregate container status counts.
func (s *ContainerService) Status(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/status")
}

// ListStats returns stats for a list of container ids.
func (s *ContainerService) ListStats(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/list/stats")
}

// ItemStats returns stats for a single container.
func (s *ContainerService) ItemStats(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/item/stats", body)
}

// StreamLogs returns a path for streaming container logs.
func (s *ContainerService) StreamLogs(ctx context.Context) string { return "/containers/search/log" }

// DownloadLogs downloads the full container log file.
func (s *ContainerService) DownloadLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/download/log", body)
}

// ResourceLimit returns the system resource limits.
func (s *ContainerService) ResourceLimit(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/limit")
}

// CleanLog truncates the container log file.
func (s *ContainerService) CleanLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/clean/log", body)
}

// Inspect returns the raw docker inspect output.
func (s *ContainerService) Inspect(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/inspect", body)
}

// Rename renames a container.
func (s *ContainerService) Rename(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/rename", body)
}

// Commit commits a container's filesystem changes to a new image.
func (s *ContainerService) Commit(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/commit", body)
}

// Operate runs a lifecycle operation (start/stop/restart/kill/pause/unpause).
func (s *ContainerService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/operate", body)
}

// Prune prunes stopped containers / unused data.
func (s *ContainerService) Prune(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/prune", body)
}

// Users returns the OS users of a container.
func (s *ContainerService) Users(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/users", body)
}
