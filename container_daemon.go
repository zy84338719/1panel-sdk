package onepanel

import "context"

// === Daemon JSON / docker status ===

// DaemonJSON returns the current docker daemon configuration.
func (s *ContainerService) DaemonJSON(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/daemonjson")
}

// DaemonJSONFile returns the path to the daemon.json file.
func (s *ContainerService) DaemonJSONFile(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/daemonjson/file")
}

// DockerStatus returns the running state of the docker daemon.
func (s *ContainerService) DockerStatus(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/containers/docker/status")
}

// OperateDocker starts/stops/restarts the docker daemon.
func (s *ContainerService) OperateDocker(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/docker/operate", body)
}

// UpdateDaemonJSON updates the docker daemon configuration.
func (s *ContainerService) UpdateDaemonJSON(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/daemonjson/update", body)
}

// UpdateLogOption updates the docker log driver.
func (s *ContainerService) UpdateLogOption(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/logoption/update", body)
}

// UpdateIPv6Option toggles IPv6 in docker daemon.
func (s *ContainerService) UpdateIPv6Option(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/ipv6option/update", body)
}

// UpdateDaemonJSONByFile replaces daemon.json with the supplied content.
func (s *ContainerService) UpdateDaemonJSONByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/containers/daemonjson/update/byfile", body)
}

// Call invokes an arbitrary /containers/* endpoint.
func (s *ContainerService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
