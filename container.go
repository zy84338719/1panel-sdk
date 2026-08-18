package onepanel

import "context"

// ContainerService covers /containers/* (docker container, image, network, volume, compose, repo).
type ContainerService struct {
	ServiceBase
}

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

// === Container files ===

// ListFiles lists container files at a path.
func (s *ContainerService) ListFiles(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/files/search", body)
}

// UploadFile uploads a file into a container.
func (s *ContainerService) UploadFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/files/upload", body)
}

// FileContent reads the content of a file inside a container.
func (s *ContainerService) FileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/files/content", body)
}

// FileSize returns the size of a file inside a container.
func (s *ContainerService) FileSize(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/files/size", body)
}

// DeleteFile removes a file from a container.
func (s *ContainerService) DeleteFile(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/files/del", body)
}

// DownloadFileURL returns the download endpoint.
func (s *ContainerService) DownloadFileURL() string { return "/containers/files/download" }

// === Image registry / repo ===

// ListRepos lists image repositories configured on this node.
func (s *ContainerService) ListRepos(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/containers/repo")
}

// CheckRepoStatus checks a repository's connectivity.
func (s *ContainerService) CheckRepoStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/repo/status", body)
}

// SearchRepos searches the configured repositories.
func (s *ContainerService) SearchRepos(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/repo/search", body)
}

// UpdateRepo updates an existing repository.
func (s *ContainerService) UpdateRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/repo/update", body)
}

// CreateRepo adds a new repository.
func (s *ContainerService) CreateRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/repo", body)
}

// DeleteRepo removes a repository.
func (s *ContainerService) DeleteRepo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/repo/del", body)
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

// === Compose templates ===

// ListComposeTemplates lists the installed compose templates.
func (s *ContainerService) ListComposeTemplates(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/containers/template")
}

// SearchComposeTemplates searches compose templates.
func (s *ContainerService) SearchComposeTemplates(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/template/search", body)
}

// UpdateComposeTemplate updates a compose template.
func (s *ContainerService) UpdateComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/template/update", body)
}

// BatchComposeTemplate batches updates on compose templates.
func (s *ContainerService) BatchComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/template/batch", body)
}

// CreateComposeTemplate adds a new compose template.
func (s *ContainerService) CreateComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/template", body)
}

// DeleteComposeTemplate removes a compose template.
func (s *ContainerService) DeleteComposeTemplate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/containers/template/del", body)
}

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
