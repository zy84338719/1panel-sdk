package onepanel

import "context"


type GroupsService struct {
	ServiceBase
}

// Create creates a group.
func (s *GroupsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/groups", body)
}

// Delete deletes a group.
func (s *GroupsService) Delete(ctx context.Context, id uint) (map[string]any, error) {
return s.postMap(ctx, "/groups/del", IDReq{ID: id})
}

// Update updates a group.
func (s *GroupsService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/groups/update", body)
}

// List lists groups.
func (s *GroupsService) List(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/groups/search", nil)
}

// Call invokes an arbitrary /groups/* endpoint.
func (s *GroupsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// CommandsService covers /commands/* (user-defined command library).

