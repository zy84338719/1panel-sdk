package onepanel

import "context"

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
