package onepanel

import "context"

type FavoriteService struct {
	ServiceBase
}

// Create creates a favorite.
func (s *FavoriteService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/favorite", body)
}

// Delete deletes a favorite.
func (s *FavoriteService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/favorite/del", body)
}

// Search searches favorites.
func (s *FavoriteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/favorite/search", body)
}

// Call invokes an arbitrary /files/favorite/* endpoint.
func (s *FavoriteService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// TaskService covers /tasks/* (long-running task progress queries).
