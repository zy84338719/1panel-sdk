package onepanel

import "context"
type WebsiteService struct {
	ServiceBase
}

// CreateWebsite creates a website record.
func (s *WebsiteService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteWebsite deletes a website by id.
func (s *WebsiteService) Delete(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/del", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateWebsite updates a website's config.
func (s *WebsiteService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetWebsite fetches a website by id.
func (s *WebsiteService) GetWebsite(ctx context.Context, id uint) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/info", IDReq{ID: id}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SearchWebsites paginates websites.
func (s *WebsiteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListWebsites lists websites (lightweight).
func (s *WebsiteService) List(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/list", map[string]any{}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OperateWebsite runs a lifecycle op (start/stop/restart/reload).
func (s *WebsiteService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/operate", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateTags updates a website's tags.
func (s *WebsiteService) UpdateTags(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/tag", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Logs returns the website's runtime log.
func (s *WebsiteService) Logs(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ClearLog clears the website's log file.
func (s *WebsiteService) ClearLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/clean/log", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPSetting toggles the HTTPS configuration.
func (s *WebsiteService) HTTPSetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/https", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPSStatus returns whether HTTPS is enabled.
func (s *WebsiteService) HTTPSStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/https/status", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AddAlias adds an alias domain to a website.
func (s *WebsiteService) AddAlias(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/alias", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteAlias deletes an alias domain.
func (s *WebsiteService) DeleteAlias(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/alias/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AddDomain adds a primary domain to a website.
func (s *WebsiteService) AddDomain(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/domain", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteDomain removes a primary domain.
func (s *WebsiteService) DeleteDomain(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/domain/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Domains returns the domains associated with a website.
func (s *WebsiteService) Domains(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/domains", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertList lists certificates available to a website.
func (s *WebsiteService) CertList(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/certs", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertInstall requests/links an SSL certificate.
func (s *WebsiteService) CertInstall(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/install", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertUpdate updates the certificate assignment.
func (s *WebsiteService) CertUpdate(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertDetail returns details of a single certificate.
func (s *WebsiteService) CertDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/detail", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertObtain obtains a new certificate via ACME.
func (s *WebsiteService) CertObtain(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/obtain", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertRenew renews an existing certificate.
func (s *WebsiteService) CertRenew(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/renew", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertReplace replaces the certificate of a website.
func (s *WebsiteService) CertReplace(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/replace", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertApply applies a CA-signed certificate.
func (s *WebsiteService) CertApply(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/apply", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertCancel cancels a pending certificate request.
func (s *WebsiteService) CertCancel(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/cancel", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertDelete removes a certificate.
func (s *WebsiteService) CertDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertDownloadURL returns the download URL for a certificate bundle.
func (s *WebsiteService) CertDownloadURL() string { return "/websites/cert/download" }

// CertUpload uploads a certificate.
func (s *WebsiteService) CertUpload(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/upload", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertSearch searches certificates.
func (s *WebsiteService) CertSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/cert/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CertGroupList lists the certificate groups.
func (s *WebsiteService) CertGroupList(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Get(ctx, "/websites/cert/group", &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DNSAccountList lists DNS provider accounts.
func (s *WebsiteService) DNSAccountList(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/dns/account", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AcmeAccountList lists ACME accounts.
func (s *WebsiteService) AcmeAccountList(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/acme/account", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/* endpoint.
func (s *WebsiteService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteSSLService covers /websites/ssl/* (acme accounts, dns accounts, CAs).

type WebsiteSSLService struct {
	ServiceBase
}

// Search searches the SSL resources.
func (s *WebsiteSSLService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ssl/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates an SSL resource.
func (s *WebsiteSSLService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ssl", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates an SSL resource.
func (s *WebsiteSSLService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ssl/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an SSL resource.
func (s *WebsiteSSLService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ssl/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/ssl/* endpoint.
func (s *WebsiteSSLService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteCAService covers /websites/ca/* (private CA management).

type WebsiteCAService struct {
	ServiceBase
}

// Create creates a private CA.
func (s *WebsiteCAService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ca", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches private CAs.
func (s *WebsiteCAService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ca/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete removes a private CA.
func (s *WebsiteCAService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/ca/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/ca/* endpoint.
func (s *WebsiteCAService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteDNSAccountService covers /websites/dns/* (DNS provider accounts for ACME DNS-01).

type WebsiteDNSAccountService struct {
	ServiceBase
}

// Create creates a DNS provider account.
func (s *WebsiteDNSAccountService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/dns", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a DNS provider account.
func (s *WebsiteDNSAccountService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/dns/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches DNS provider accounts.
func (s *WebsiteDNSAccountService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/dns/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a DNS provider account.
func (s *WebsiteDNSAccountService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/dns/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/dns/* endpoint.
func (s *WebsiteDNSAccountService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteAcmeAccountService covers /websites/acme/* (ACME account management).

type WebsiteAcmeAccountService struct {
	ServiceBase
}

// Create creates an ACME account.
func (s *WebsiteAcmeAccountService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/acme", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates an ACME account.
func (s *WebsiteAcmeAccountService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/acme/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches ACME accounts.
func (s *WebsiteAcmeAccountService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/acme/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an ACME account.
func (s *WebsiteAcmeAccountService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/acme/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/acme/* endpoint.
func (s *WebsiteAcmeAccountService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// WebsiteTemplateService covers /websites/template/* (website templates).

type WebsiteTemplateService struct {
	ServiceBase
}

// Create creates a website template.
func (s *WebsiteTemplateService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/template", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a website template.
func (s *WebsiteTemplateService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/template/update", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a website template.
func (s *WebsiteTemplateService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/template/del", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Search searches website templates.
func (s *WebsiteTemplateService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	if err := s.Post(ctx, "/websites/template/search", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Call invokes an arbitrary /websites/template/* endpoint.
func (s *WebsiteTemplateService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// DatabaseService covers /databases/* (mysql, postgres, mongodb, redis).


