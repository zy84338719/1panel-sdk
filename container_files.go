package onepanel

import "context"

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

