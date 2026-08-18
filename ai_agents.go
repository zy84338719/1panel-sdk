package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type AIService struct {
	ServiceBase
}

// OllamaClose closes an Ollama model.
func (s *AIService) OllamaClose(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/close", body)
}

// OllamaCreate creates an Ollama model.
func (s *AIService) OllamaCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model", body)
}

// OllamaRecreate recreates an Ollama model.
func (s *AIService) OllamaRecreate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model/recreate", body)
}

// OllamaSearch searches Ollama models.
func (s *AIService) OllamaSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model/search", body)
}

// OllamaSync syncs Ollama models.
func (s *AIService) OllamaSync(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model/sync", body)
}

// OllamaLoadDetail loads Ollama model details.
func (s *AIService) OllamaLoadDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model/load", body)
}

// OllamaDelete deletes an Ollama model.
func (s *AIService) OllamaDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/ollama/model/del", body)
}

// GPUInfo returns the GPU info.
func (s *AIService) GPUInfo(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/ai/gpu/load")
}

// GPUMonitor loads the GPU monitor data.
func (s *AIService) GPUMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/gpu/search", body)
}

// GPUOptions returns the GPU monitoring options.
func (s *AIService) GPUOptions(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/ai/gpu/options")
}

// McpSearch searches MCP servers.
func (s *AIService) McpSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/search", body)
}

// McpServerDetail returns an MCP server's detail.
func (s *AIService) McpServerDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/server/detail", body)
}

// McpServerCreate creates an MCP server.
func (s *AIService) McpServerCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/server", body)
}

// McpServerUpdate updates an MCP server.
func (s *AIService) McpServerUpdate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/server/update", body)
}

// McpServerDelete deletes an MCP server.
func (s *AIService) McpServerDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/server/del", body)
}

// McpServerOperate operates an MCP server.
func (s *AIService) McpServerOperate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/mcp/server/op", body)
}

// Call invokes an arbitrary /ai/* endpoint.
func (s *AIService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// AgentsService covers /agents/* (managed agent instances).

type AgentsService struct {
	ServiceBase
}

// Create creates an agent.
func (s *AgentsService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents", body)
}

// BatchInstall installs agents in batch.
func (s *AgentsService) BatchInstall(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/batch/install", body)
}

// BatchUpgrade upgrades agents in batch.
func (s *AgentsService) BatchUpgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/batch/upgrade", body)
}

// Search searches agents.
func (s *AgentsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/search", body)
}

// DeleteCheck pre-flight check for agent deletion.
func (s *AgentsService) DeleteCheck(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/delete/check", body)
}

// Delete deletes an agent.
func (s *AgentsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/delete", body)
}

// ResetToken resets an agent's access token.
func (s *AgentsService) ResetToken(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/token/reset", body)
}

// Overview returns the agent overview.
func (s *AgentsService) Overview(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/overview", body)
}

// Remark updates an agent's remark.
func (s *AgentsService) Remark(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/agents/remark", body)
}

// Providers returns the available agent providers.
func (s *AgentsService) Providers(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/ai/accounts/providers")
}

// CreateAccount creates an agent account.
func (s *AgentsService) CreateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/accounts", body)
}

// UpdateAccount updates an agent account.
func (s *AgentsService) UpdateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/accounts/update", body)
}

// SearchAccounts searches agent accounts.
func (s *AgentsService) SearchAccounts(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/accounts/search", body)
}

// DeleteAccount deletes an agent account.
func (s *AgentsService) DeleteAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/accounts/delete", body)
}

// VerifyAccount verifies an agent account.
func (s *AgentsService) VerifyAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/ai/accounts/verify", body)
}

// Call invokes an arbitrary /agents/* or /ai/accounts/* endpoint.
func (s *AgentsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SSHService covers /hosts/ssh/* (SSH config, logs, root cert, host SSH).
