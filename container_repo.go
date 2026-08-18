package onepanel

import "context"

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

