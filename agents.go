package onepanel

import "context"

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
