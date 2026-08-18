// Package onepanel - this file used to host the seven service types below,
// but was split into per-domain files for maintainability. The split keeps
// every exported name in the same package, so callers do not need to
// change their imports.
//
//   - AppService                 -> app.go
//   - WebsiteService             -> website.go
//   - WebsiteSSLService          -> website.go
//   - WebsiteCAService           -> website.go
//   - WebsiteDNSAccountService   -> website.go
//   - WebsiteAcmeAccountService  -> website.go
//   - WebsiteTemplateService     -> website.go
//   - DatabaseService            -> database.go
//
// Re-running scripts/gen-from-swagger.py will only re-emit the codegen file
// (services_swagger.go); these hand-written service files are the source of
// truth.
package onepanel
