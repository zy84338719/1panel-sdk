package onepanel

import "context"


type FirewallService struct {
	ServiceBase
}

// LoadBaseInfo returns the firewall base info.
func (s *FirewallService) LoadBaseInfo(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/base", body)
}

// SearchRule searches firewall rules.
func (s *FirewallService) SearchRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/search", body)
}

// Operate runs a firewall op (start/stop/restart).
func (s *FirewallService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/operate", body)
}

// OperatePortRule creates a port rule.
func (s *FirewallService) OperatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/port", body)
}

// OperateForwardRule creates a forward rule.
func (s *FirewallService) OperateForwardRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/forward", body)
}

// OperateIPRule creates an IP rule.
func (s *FirewallService) OperateIPRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/ip", body)
}

// BatchOperateRule batch-operates rules.
func (s *FirewallService) BatchOperateRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/batch", body)
}

// UpdatePortRule updates a port rule.
func (s *FirewallService) UpdatePortRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/port", body)
}

// UpdateAddrRule updates an address rule.
func (s *FirewallService) UpdateAddrRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/addr", body)
}

// UpdateDescription updates a firewall rule's description.
func (s *FirewallService) UpdateDescription(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/update/description", body)
}

// SearchFilterRules searches iptables filter rules.
func (s *FirewallService) SearchFilterRules(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/search", body)
}

// OperateFilterRule operates an iptables filter rule.
func (s *FirewallService) OperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/operate", body)
}

// BatchOperateFilterRule batch operates filter rules.
func (s *FirewallService) BatchOperateFilterRule(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/rule/batch", body)
}

// OperateFilterChain operates an iptables filter chain.
func (s *FirewallService) OperateFilterChain(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/operate", body)
}

// LoadChainStatus returns the chain status.
func (s *FirewallService) LoadChainStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/hosts/firewall/filter/chain/status", body)
}

// Call invokes an arbitrary /hosts/firewall/* endpoint.
func (s *FirewallService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// NginxService covers /nginx/* (Nginx config management).
