// Package onepanel - this file used to hold the 21 services listed below,
// but the file was split into per-domain files to make maintenance easier.
// All the type definitions remain in this package, so callers do not need
// to change their imports.
//
//   - BackupService  -> backup_cronjob.go
//   - CronjobService -> backup_cronjob.go
//   - FileService    -> file.go
//   - SettingsService, LogsService          -> settings_logs.go
//   - GroupsService, CommandsService, ScriptService -> groups_commands_script.go
//   - ToolboxService, AlertsService         -> toolbox_alerts.go
//   - AIService, AgentsService              -> ai_agents.go
//   - SSHService, MonitorService, FirewallService -> host_extra.go
//   - NginxService, ProcessService, RuntimeService -> nginx_process_runtime.go
//   - SnapshotService, FavoriteService, TaskService -> snapshot_favorite_task.go
//
// The split preserves all exported names. Re-running the codegen
// (scripts/gen-from-swagger.py) will re-emit them into the new files.
package onepanel
