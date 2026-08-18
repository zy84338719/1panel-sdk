// Package onepanel is a complete Go SDK for the 1Panel server management panel
// (https://github.com/1Panel-dev/1Panel). It exposes every public endpoint
// surfaced by the panel frontend and ships typed helpers for the most common
// workflows: authentication, host/firewall/SSH management, container control,
// application store operations, website + SSL automation, database admin,
// backups, scheduled tasks, file management, and the dashboard APIs.
//
// # Usage
//
//	c, err := onepanel.New(onepanel.Options{
//	    BaseURL:  "https://1panel.example.com",
//	    Entrance: "1panel_entrance",
//	    Username: "admin",
//	    Password: "secret",
//	})
//	if err != nil { log.Fatal(err) }
//
//	// ... call any method on c.Auth, c.Host, c.Container, ...
//	hosts, _ := c.Host.SearchHosts(ctx, onepanel.PageInfo{Page: 1, PageSize: 20})
//
// # Endpoints covered
//
// The SDK mirrors the public URL layout used by the 1Panel frontend: /core/*
// for the master panel APIs and /<resource>/* for node-facing APIs. The
// "CurrentNode" header is set automatically per call.
package onepanel

import (
	"context"

	"github.com/zy84338719/1panel-sdk/client"
)

// doer is the minimal HTTP surface used by sub-services. Implemented by both
// *client.Client (default node) and *client.NodeClient (explicit node).
type doer interface {
	Get(ctx context.Context, path string, out any) error
	Post(ctx context.Context, path string, body, out any) error
	Put(ctx context.Context, path string, body, out any) error
	Delete(ctx context.Context, path string, body, out any) error
	Do(ctx context.Context, method, path string, body, out any) error
}

// Options configures a SDK instance.
type Options struct {
	// BaseURL is the public panel URL, e.g. "https://1panel.example.com".
	// Either BaseURL or Endpoint must be set.
	BaseURL string

	// Endpoint is the direct host[:port] of the core API. Use when the public
	// entrance is unavailable (e.g. accessing the core service directly on 9999).
	Endpoint string

	// Entrance is the entrance sub-path used to obscure the panel. Empty means
	// the panel is served at the root path.
	Entrance string

	// Language is sent as the "Accept-Language" header. Defaults to "zh-CN".
	Language string

	// NodeID is the default node id used for node-facing APIs. Empty means local.
	NodeID string

	// Login credentials. If both are set, New logs in synchronously.
	Username string
	Password string
	// OnLogin is invoked right after a successful login so callers can persist
	// the cookie jar or refresh tokens.
	OnLogin func(*client.LoginResult)
}

// SDK bundles a configured client and all sub-services.
type SDK struct {
	C *client.Client

	Auth        *AuthService
	Dashboard   *DashboardService
	Host        *HostService
	Container   *ContainerService
	App         *AppService
	Website     *WebsiteService
	Database    *DatabaseService
	Backup      *BackupService
	Cronjob     *CronjobService
	File        *FileService
	Settings    *SettingsService
	Logs        *LogsService
	Groups      *GroupsService
	Commands    *CommandsService
	Script      *ScriptService
	Toolbox     *ToolboxService
	Alerts      *AlertsService
	AI          *AIService
	Agent       *AgentsService // alias kept for back-compat
	Agents      *AgentsService
	AIAccount   *AIAccountService
	AIAgent     *AIAgentService
	AIDomain    *AIDomainService
	AIMcp       *AIMcpService
	AITensor    *AITensorService
	CoreAuth    *CoreAuthService
	CoreBackup  *CoreBackupsService
	CoreCommand *CoreCommandsService
	CoreGroup   *CoreGroupsService
	CoreLog     *CoreLogsService
	CoreScript  *CoreScriptService
	CoreSetting *CoreSettingsService
	Health      *HealthService
	OpenResty   *OpenRestyService
	SSH         *SSHService
	Monitor     *MonitorService
	Firewall    *FirewallService
	SSL         *SSLService
	WebsiteCA   *WebsiteCAService
	WebsiteDNS  *WebsiteDNSAccountService
	WebsiteAcme *WebsiteAcmeAccountService
	WebsiteTpl  *WebsiteTemplateService
	Nginx       *NginxService
	Process     *ProcessService
	Runtime     *RuntimeService
	Snapshot    *SnapshotService
	Favorite    *FavoriteService
	Task        *TaskService
}

// New constructs a SDK. If Username+Password are set, it performs the initial
// login. The login honors MFA via the Options.MFA field when present.
func New(opt Options) (*SDK, error) {
	c, err := client.New(client.Config{
		BaseURL:  opt.BaseURL,
		Endpoint: opt.Endpoint,
		Entrance: opt.Entrance,
		Language: opt.Language,
		NodeID:   opt.NodeID,
	})
	if err != nil {
		return nil, err
	}
	s := &SDK{C: c}
	s.bind(c)
	if opt.Username != "" && opt.Password != "" {
		login, err := s.Auth.Login(LoginForm{
			Name:     opt.Username,
			Password: opt.Password,
		})
		if err != nil {
			return nil, err
		}
		if opt.OnLogin != nil {
			opt.OnLogin(login)
		}
	}
	return s, nil
}

// NewFromClient wraps an existing client.Client. Useful when the caller wants
// to fully control the underlying HTTP transport.
func NewFromClient(c *client.Client) *SDK {
	s := &SDK{C: c}
	s.bind(c)
	return s
}

// OnNode returns a shallow copy of the SDK whose default node id is overridden.
// All sub-services of the returned SDK target that node.
func (s *SDK) OnNode(nodeID string) *SDK {
	clone := *s
	node := s.C.OnNode(nodeID)
	clone.bind(node)
	return &clone
}

// SetLanguage changes the Accept-Language header.
func (s *SDK) SetLanguage(lang string) { s.C.SetLanguage(lang) }

// SetNode changes the default node id on the underlying client.
func (s *SDK) SetNode(id string) { s.C.SetNode(id) }

func (s *SDK) bind(d doer) {
	s.Auth = &AuthService{ServiceBase: ServiceBase{d: d}}
	s.Dashboard = &DashboardService{ServiceBase: ServiceBase{d: d}}
	s.Host = &HostService{ServiceBase: ServiceBase{d: d}}
	s.Container = &ContainerService{ServiceBase: ServiceBase{d: d}}
	s.App = &AppService{ServiceBase: ServiceBase{d: d}}
	s.Website = &WebsiteService{ServiceBase: ServiceBase{d: d}}
	s.Database = &DatabaseService{ServiceBase: ServiceBase{d: d}}
	s.Backup = &BackupService{ServiceBase: ServiceBase{d: d}}
	s.Cronjob = &CronjobService{ServiceBase: ServiceBase{d: d}}
	s.File = &FileService{ServiceBase: ServiceBase{d: d}}
	s.Settings = &SettingsService{ServiceBase: ServiceBase{d: d}}
	s.Logs = &LogsService{ServiceBase: ServiceBase{d: d}}
	s.Groups = &GroupsService{ServiceBase: ServiceBase{d: d}}
	s.Commands = &CommandsService{ServiceBase: ServiceBase{d: d}}
	s.Script = &ScriptService{ServiceBase: ServiceBase{d: d}}
	s.Toolbox = &ToolboxService{ServiceBase: ServiceBase{d: d}}
	s.Alerts = &AlertsService{ServiceBase: ServiceBase{d: d}}
	s.AI = &AIService{ServiceBase: ServiceBase{d: d}}
	s.Agent = &AgentsService{ServiceBase: ServiceBase{d: d}}
	s.Agents = &AgentsService{ServiceBase: ServiceBase{d: d}}
	s.AIAccount = &AIAccountService{ServiceBase: ServiceBase{d: d}}
	s.AIAgent = &AIAgentService{ServiceBase: ServiceBase{d: d}}
	s.AIDomain = &AIDomainService{ServiceBase: ServiceBase{d: d}}
	s.AIMcp = &AIMcpService{ServiceBase: ServiceBase{d: d}}
	s.AITensor = &AITensorService{ServiceBase: ServiceBase{d: d}}
	s.CoreAuth = &CoreAuthService{ServiceBase: ServiceBase{d: d}}
	s.CoreBackup = &CoreBackupsService{ServiceBase: ServiceBase{d: d}}
	s.CoreCommand = &CoreCommandsService{ServiceBase: ServiceBase{d: d}}
	s.CoreGroup = &CoreGroupsService{ServiceBase: ServiceBase{d: d}}
	s.CoreLog = &CoreLogsService{ServiceBase: ServiceBase{d: d}}
	s.CoreScript = &CoreScriptService{ServiceBase: ServiceBase{d: d}}
	s.CoreSetting = &CoreSettingsService{ServiceBase: ServiceBase{d: d}}
	s.Health = &HealthService{ServiceBase: ServiceBase{d: d}}
	s.OpenResty = &OpenRestyService{ServiceBase: ServiceBase{d: d}}
	s.SSH = &SSHService{ServiceBase: ServiceBase{d: d}}
	s.Monitor = &MonitorService{ServiceBase: ServiceBase{d: d}}
	s.Firewall = &FirewallService{ServiceBase: ServiceBase{d: d}}
	s.SSL = &SSLService{ServiceBase: ServiceBase{d: d}}
	s.WebsiteCA = &WebsiteCAService{ServiceBase: ServiceBase{d: d}}
	s.WebsiteDNS = &WebsiteDNSAccountService{ServiceBase: ServiceBase{d: d}}
	s.WebsiteAcme = &WebsiteAcmeAccountService{ServiceBase: ServiceBase{d: d}}
	s.WebsiteTpl = &WebsiteTemplateService{ServiceBase: ServiceBase{d: d}}
	s.Nginx = &NginxService{ServiceBase: ServiceBase{d: d}}
	s.Process = &ProcessService{ServiceBase: ServiceBase{d: d}}
	s.Runtime = &RuntimeService{ServiceBase: ServiceBase{d: d}}
	s.Snapshot = &SnapshotService{ServiceBase: ServiceBase{d: d}}
	s.Favorite = &FavoriteService{ServiceBase: ServiceBase{d: d}}
	s.Task = &TaskService{ServiceBase: ServiceBase{d: d}}
}
