package onepanel

import "context"

// ContainerService covers /containers/* (docker container, image, network, volume, compose, repo).
type ContainerService struct {
	ServiceBase
}

// === Container CRUD ===

// Stats returns the live stats of a container by id.
func (s *ContainerService) Stats(ctx context.Context, id string) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/stats/"+id, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a new container.
func (s *ContainerService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a container (restart policy, name, env, etc).
func (s *ContainerService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Upgrade rebuilds a container with a new image.
func (s *ContainerService) Upgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/upgrade", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Info returns detailed info for a container.
func (s *ContainerService) Info(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/info", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches containers.
func (s *ContainerService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// List lists containers (lightweight).
func (s *ContainerService) List(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/list", map[string]any{}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListByImage lists containers using a given image.
func (s *ContainerService) ListByImage(ctx context.Context, name string) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/list/byimage", map[string]string{"name": name}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Status returns aggregate container status counts.
func (s *ContainerService) Status(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/status", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListStats returns stats for a list of container ids.
func (s *ContainerService) ListStats(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/list/stats", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ItemStats returns stats for a single container.
func (s *ContainerService) ItemStats(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/item/stats", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// StreamLogs returns a path for streaming container logs.
func (s *ContainerService) StreamLogs(ctx context.Context) string { return "/containers/search/log" }

// DownloadLogs downloads the full container log file.
func (s *ContainerService) DownloadLogs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/download/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceLimit returns the system resource limits.
func (s *ContainerService) ResourceLimit(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/limit", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanLog truncates the container log file.
func (s *ContainerService) CleanLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/clean/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Inspect returns the raw docker inspect output.
func (s *ContainerService) Inspect(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/inspect", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Rename renames a container.
func (s *ContainerService) Rename(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/rename", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Commit commits a container's filesystem changes to a new image.
func (s *ContainerService) Commit(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/commit", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Operate runs a lifecycle operation (start/stop/restart/kill/pause/unpause).
func (s *ContainerService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Prune prunes stopped containers / unused data.
func (s *ContainerService) Prune(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/prune", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Users returns the OS users of a container.
func (s *ContainerService) Users(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/users", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Container files ===

// ListFiles lists container files at a path.
func (s *ContainerService) ListFiles(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/files/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UploadFile uploads a file into a container.
func (s *ContainerService) UploadFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/files/upload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// FileContent reads the content of a file inside a container.
func (s *ContainerService) FileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/files/content", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// FileSize returns the size of a file inside a container.
func (s *ContainerService) FileSize(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/files/size", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteFile removes a file from a container.
func (s *ContainerService) DeleteFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/files/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadFileURL returns the download endpoint.
func (s *ContainerService) DownloadFileURL() string { return "/containers/files/download" }

// === Image registry / repo ===

// ListRepos lists image repositories configured on this node.
func (s *ContainerService) ListRepos(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/repo", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CheckRepoStatus checks a repository's connectivity.
func (s *ContainerService) CheckRepoStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/repo/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRepos searches the configured repositories.
func (s *ContainerService) SearchRepos(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/repo/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRepo updates an existing repository.
func (s *ContainerService) UpdateRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/repo/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateRepo adds a new repository.
func (s *ContainerService) CreateRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/repo", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteRepo removes a repository.
func (s *ContainerService) DeleteRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/repo/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Compose ===

// SearchCompose searches docker-compose projects.
func (s *ContainerService) SearchCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateCompose creates a new compose project.
func (s *ContainerService) CreateCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ComposeEnv loads the env file of a compose project.
func (s *ContainerService) ComposeEnv(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/env", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TestCompose validates a compose file syntactically.
func (s *ContainerService) TestCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/test", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorCompose runs an operation (up/down/restart/...).
func (s *ContainerService) OperatorCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CleanComposeLog cleans the compose log file.
func (s *ContainerService) CleanComposeLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/clean/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateCompose updates a compose project.
func (s *ContainerService) UpdateCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PinCompose pins/unpins a compose project.
func (s *ContainerService) PinCompose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/compose/pin", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Compose templates ===

// ListComposeTemplates lists the installed compose templates.
func (s *ContainerService) ListComposeTemplates(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/template", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchComposeTemplates searches compose templates.
func (s *ContainerService) SearchComposeTemplates(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/template/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateComposeTemplate updates a compose template.
func (s *ContainerService) UpdateComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/template/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchComposeTemplate batches updates on compose templates.
func (s *ContainerService) BatchComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/template/batch", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateComposeTemplate adds a new compose template.
func (s *ContainerService) CreateComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/template", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteComposeTemplate removes a compose template.
func (s *ContainerService) DeleteComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/template/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Image ===

// ListImages lists images.
func (s *ContainerService) ListImages(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/image", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListAllImages lists all images including dangling.
func (s *ContainerService) ListAllImages(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/image/all", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchImages searches local images.
func (s *ContainerService) SearchImages(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PullImage pulls an image from a registry.
func (s *ContainerService) PullImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/pull", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PushImage pushes an image to a registry.
func (s *ContainerService) PushImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/push", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SaveImage saves an image to a tarball.
func (s *ContainerService) SaveImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/save", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LoadImage loads an image from a tarball.
func (s *ContainerService) LoadImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RemoveImage removes an image.
func (s *ContainerService) RemoveImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/remove", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TagImage tags an image.
func (s *ContainerService) TagImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/tag", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BuildImage builds an image from a dockerfile.
func (s *ContainerService) BuildImage(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/image/build", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Network ===

// ListNetworks lists docker networks.
func (s *ContainerService) ListNetworks(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/network", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteNetwork deletes a docker network.
func (s *ContainerService) DeleteNetwork(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/network/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchNetworks searches docker networks.
func (s *ContainerService) SearchNetworks(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/network/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateNetwork creates a docker network.
func (s *ContainerService) CreateNetwork(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/network", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Volume ===

// ListVolumes lists docker volumes.
func (s *ContainerService) ListVolumes(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/volume", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteVolume deletes a docker volume.
func (s *ContainerService) DeleteVolume(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/volume/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchVolumes searches docker volumes.
func (s *ContainerService) SearchVolumes(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/volume/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateVolume creates a docker volume.
func (s *ContainerService) CreateVolume(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/volume", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// === Daemon JSON / docker status ===

// DaemonJSON returns the current docker daemon configuration.
func (s *ContainerService) DaemonJSON(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/daemonjson", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonJSONFile returns the path to the daemon.json file.
func (s *ContainerService) DaemonJSONFile(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/daemonjson/file", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DockerStatus returns the running state of the docker daemon.
func (s *ContainerService) DockerStatus(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/containers/docker/status", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateDocker starts/stops/restarts the docker daemon.
func (s *ContainerService) OperateDocker(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/docker/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDaemonJSON updates the docker daemon configuration.
func (s *ContainerService) UpdateDaemonJSON(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/daemonjson/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateLogOption updates the docker log driver.
func (s *ContainerService) UpdateLogOption(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/logoption/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateIPv6Option toggles IPv6 in docker daemon.
func (s *ContainerService) UpdateIPv6Option(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/ipv6option/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDaemonJSONByFile replaces daemon.json with the supplied content.
func (s *ContainerService) UpdateDaemonJSONByFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/containers/daemonjson/update/byfile", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /containers/* endpoint.
func (s *ContainerService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}
