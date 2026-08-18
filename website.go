package onepanel

import "context"

type WebsiteService struct {
	ServiceBase
}

// CreateWebsite creates a website record.
func (s *WebsiteService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites", body)
}

// DeleteWebsite deletes a website by id.
func (s *WebsiteService) Delete(ctx context.Context, id uint) (map[string]any, error) {
	return s.postMap(ctx, "/websites/del", IDReq{ID: id})
}

// UpdateWebsite updates a website's config.
func (s *WebsiteService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/update", body)
}

// GetWebsite fetches a website by id.
func (s *WebsiteService) GetWebsite(ctx context.Context, id uint) (map[string]any, error) {
	return s.postMap(ctx, "/websites/info", IDReq{ID: id})
}

// SearchWebsites paginates websites.
func (s *WebsiteService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/search", body)
}

// ListWebsites lists websites (lightweight).
func (s *WebsiteService) List(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/websites/list", map[string]any{})
}

// OperateWebsite runs a lifecycle op (start/stop/restart/reload).
func (s *WebsiteService) Operate(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/operate", body)
}

// UpdateTags updates a website's tags.
func (s *WebsiteService) UpdateTags(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/tag", body)
}

// Logs returns the website's runtime log.
func (s *WebsiteService) Logs(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/log", body)
}

// ClearLog clears the website's log file.
func (s *WebsiteService) ClearLog(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/clean/log", body)
}

// HTTPSetting toggles the HTTPS configuration.
func (s *WebsiteService) HTTPSetting(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/https", body)
}

// HTTPSStatus returns whether HTTPS is enabled.
func (s *WebsiteService) HTTPSStatus(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/https/status", body)
}

// AddAlias adds an alias domain to a website.
func (s *WebsiteService) AddAlias(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/alias", body)
}

// DeleteAlias deletes an alias domain.
func (s *WebsiteService) DeleteAlias(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/alias/del", body)
}

// AddDomain adds a primary domain to a website.
func (s *WebsiteService) AddDomain(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/domain", body)
}

// DeleteDomain removes a primary domain.
func (s *WebsiteService) DeleteDomain(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/domain/del", body)
}

// Domains returns the domains associated with a website.
func (s *WebsiteService) Domains(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/domains", body)
}

// CertList lists certificates available to a website.
func (s *WebsiteService) CertList(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/certs", body)
}

// CertInstall requests/links an SSL certificate.
func (s *WebsiteService) CertInstall(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/install", body)
}

// CertUpdate updates the certificate assignment.
func (s *WebsiteService) CertUpdate(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/update", body)
}

// CertDetail returns details of a single certificate.
func (s *WebsiteService) CertDetail(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/detail", body)
}

// CertObtain obtains a new certificate via ACME.
func (s *WebsiteService) CertObtain(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/obtain", body)
}

// CertRenew renews an existing certificate.
func (s *WebsiteService) CertRenew(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/renew", body)
}

// CertReplace replaces the certificate of a website.
func (s *WebsiteService) CertReplace(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/replace", body)
}

// CertApply applies a CA-signed certificate.
func (s *WebsiteService) CertApply(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/apply", body)
}

// CertCancel cancels a pending certificate request.
func (s *WebsiteService) CertCancel(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/cancel", body)
}

// CertDelete removes a certificate.
func (s *WebsiteService) CertDelete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/del", body)
}

// CertDownloadURL returns the download URL for a certificate bundle.
func (s *WebsiteService) CertDownloadURL() string { return "/websites/cert/download" }

// CertUpload uploads a certificate.
func (s *WebsiteService) CertUpload(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/upload", body)
}

// CertSearch searches certificates.
func (s *WebsiteService) CertSearch(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/cert/search", body)
}

// CertGroupList lists the certificate groups.
func (s *WebsiteService) CertGroupList(ctx context.Context) (map[string]any, error) {
	return s.getMap(ctx, "/websites/cert/group")
}

// DNSAccountList lists DNS provider accounts.
func (s *WebsiteService) DNSAccountList(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/websites/dns/account", nil)
}

// AcmeAccountList lists ACME accounts.
func (s *WebsiteService) AcmeAccountList(ctx context.Context) (map[string]any, error) {
	return s.postMap(ctx, "/websites/acme/account", nil)
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
	return s.postMap(ctx, "/websites/ssl/search", body)
}

// Create creates an SSL resource.
func (s *WebsiteSSLService) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/ssl", body)
}

// Update updates an SSL resource.
func (s *WebsiteSSLService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/ssl/update", body)
}

// Delete deletes an SSL resource.
func (s *WebsiteSSLService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/ssl/del", body)
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
	return s.postMap(ctx, "/websites/ca", body)
}

// Search searches private CAs.
func (s *WebsiteCAService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/ca/search", body)
}

// Delete removes a private CA.
func (s *WebsiteCAService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/ca/del", body)
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
	return s.postMap(ctx, "/websites/dns", body)
}

// Update updates a DNS provider account.
func (s *WebsiteDNSAccountService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/dns/update", body)
}

// Search searches DNS provider accounts.
func (s *WebsiteDNSAccountService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/dns/search", body)
}

// Delete deletes a DNS provider account.
func (s *WebsiteDNSAccountService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/dns/del", body)
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
	return s.postMap(ctx, "/websites/acme", body)
}

// Update updates an ACME account.
func (s *WebsiteAcmeAccountService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/acme/update", body)
}

// Search searches ACME accounts.
func (s *WebsiteAcmeAccountService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/acme/search", body)
}

// Delete deletes an ACME account.
func (s *WebsiteAcmeAccountService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/acme/del", body)
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
	return s.postMap(ctx, "/websites/template", body)
}

// Update updates a website template.
func (s *WebsiteTemplateService) Update(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/template/update", body)
}

// Delete deletes a website template.
func (s *WebsiteTemplateService) Delete(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/template/del", body)
}

// Search searches website templates.
func (s *WebsiteTemplateService) Search(ctx context.Context, body map[string]any) (map[string]any, error) {
	return s.postMap(ctx, "/websites/template/search", body)
}

// Call invokes an arbitrary /websites/template/* endpoint.
func (s *WebsiteTemplateService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// DatabaseService covers /databases/* (mysql, postgres, mongodb, redis).
