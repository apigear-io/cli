# cfg

Configuration management package for the APIGear CLI application.

## Purpose

The `cfg` package handles persistent application configuration using JSON files and environment variables. It provides thread-safe access to configuration values with support for:

- Reading/writing configuration from `~/.apigear/config.json`
- Environment variable overrides via `APIGEAR_*` prefixes
- Build information storage (version, commit, date)
- Recent project entries management
- Default values for all configuration keys

## Key Exports

- `Get()`, `GetString()`, `GetInt()`, `GetBool()`, `Set()` - Configuration accessors
- `SetBuildInfo()`, `GetBuildInfo()` - Build metadata
- `AppendRecentEntry()`, `RemoveRecentEntry()`, `RecentEntries()` - Recent projects
- `ConfigDir()`, `CacheDir()`, `RegistryDir()` - Directory paths
- `EditorCommand()`, `ServerPort()`, `UpdateChannel()` - Specialized getters

## Dependencies

| Package | Purpose |
|---------|---------|
| `helper` | File operations (Join, MakeDir, IsFile, WriteFile) |
