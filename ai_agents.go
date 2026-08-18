package onepanel

import "context"

// BackupService covers /backups/* (cloud backup destinations and scheduled jobs).
type AIService struct {
	ServiceBase
}

// OllamaClose closes an Ollama model.
func (s *AIService) OllamaClose(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/close", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaCreate creates an Ollama model.
func (s *AIService) OllamaCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaRecreate recreates an Ollama model.
func (s *AIService) OllamaRecreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/recreate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaSearch searches Ollama models.
func (s *AIService) OllamaSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaSync syncs Ollama models.
func (s *AIService) OllamaSync(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/sync", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaLoadDetail loads Ollama model details.
func (s *AIService) OllamaLoadDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/load", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OllamaDelete deletes an Ollama model.
func (s *AIService) OllamaDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/ollama/model/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUInfo returns the GPU info.
func (s *AIService) GPUInfo(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/gpu/load", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUMonitor loads the GPU monitor data.
func (s *AIService) GPUMonitor(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/gpu/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GPUOptions returns the GPU monitoring options.
func (s *AIService) GPUOptions(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/gpu/options", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpSearch searches MCP servers.
func (s *AIService) McpSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerDetail returns an MCP server's detail.
func (s *AIService) McpServerDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/detail", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerCreate creates an MCP server.
func (s *AIService) McpServerCreate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerUpdate updates an MCP server.
func (s *AIService) McpServerUpdate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerDelete deletes an MCP server.
func (s *AIService) McpServerDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// McpServerOperate operates an MCP server.
func (s *AIService) McpServerOperate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/mcp/server/op", body, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchInstall installs agents in batch.
func (s *AgentsService) BatchInstall(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/batch/install", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// BatchUpgrade upgrades agents in batch.
func (s *AgentsService) BatchUpgrade(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/batch/upgrade", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches agents.
func (s *AgentsService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCheck pre-flight check for agent deletion.
func (s *AgentsService) DeleteCheck(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/delete/check", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an agent.
func (s *AgentsService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/delete", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResetToken resets an agent's access token.
func (s *AgentsService) ResetToken(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/token/reset", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Overview returns the agent overview.
func (s *AgentsService) Overview(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/overview", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Remark updates an agent's remark.
func (s *AgentsService) Remark(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/agents/remark", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Providers returns the available agent providers.
func (s *AgentsService) Providers(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/ai/accounts/providers", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAccount creates an agent account.
func (s *AgentsService) CreateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAccount updates an agent account.
func (s *AgentsService) UpdateAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchAccounts searches agent accounts.
func (s *AgentsService) SearchAccounts(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteAccount deletes an agent account.
func (s *AgentsService) DeleteAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/delete", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyAccount verifies an agent account.
func (s *AgentsService) VerifyAccount(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/ai/accounts/verify", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /agents/* or /ai/accounts/* endpoint.
func (s *AgentsService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// SSHService covers /hosts/ssh/* (SSH config, logs, root cert, host SSH).
