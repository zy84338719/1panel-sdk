package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type FileService struct {
	ServiceBase
}

// List lists files at a path.
func (s *FileService) List(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/search", body)
}

// LoadDirs loads the available directories (quick navigation).
func (s *FileService) LoadDirs(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/files/dir")
}

// LoadFileContent reads the content of a file.
func (s *FileService) LoadFileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/content", body)
}

// SaveFileContent writes content to a file.
func (s *FileService) SaveFileContent(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/save", body)
}

// Upload uploads a file (multipart).
func (s *FileService) Upload(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/upload", body)
}

// Create creates a new file or directory.
func (s *FileService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files", body)
}

// Delete deletes a file or directory.
func (s *FileService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/del", body)
}

// BatchOperate batch-deletes files.
func (s *FileService) BatchOperate(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/batch/operate", body)
}

// Compress compresses files.
func (s *FileService) Compress(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/compress", body)
}

// Decompress decompresses an archive.
func (s *FileService) Decompress(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/decompress", body)
}

// ChangePermissions chmods a file.
func (s *FileService) ChangePermissions(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/permission", body)
}

// ChangeOwner chowns a file.
func (s *FileService) ChangeOwner(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/owner", body)
}

// Move moves a file.
func (s *FileService) Move(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/move", body)
}

// Rename renames a file.
func (s *FileService) Rename(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/rename", body)
}

// DownloadURL returns the download endpoint for files.
func (s *FileService) DownloadURL() string { return "/files/download" }

// Search performs a deep file search.
func (s *FileService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/grep", body)
}

// Wget downloads a file via HTTP.
func (s *FileService) Wget(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/wget", body)
}

// Token issues a download token for sharing.
func (s *FileService) Token(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/token", body)
}

// ChunkUploadInit starts a chunked upload.
func (s *FileService) ChunkUploadInit(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/chunkupload/init", body)
}

// ChunkUpload appends a chunk.
func (s *FileService) ChunkUpload(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/chunkupload", body)
}

// ChunkMerge merges uploaded chunks.
func (s *FileService) ChunkMerge(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/chunkupload/merge", body)
}

// FileHistory searches file-edit history.
func (s *FileService) FileHistory(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/history", body)
}

// RecoverFile recovers a deleted file from the recycle bin.
func (s *FileService) RecoverFile(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/files/recover", body)
}

// Call invokes an arbitrary /files/* endpoint.
func (s *FileService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SettingsService covers /settings/* (panel configuration).
