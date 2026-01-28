# repos

Template repository management with two-layer caching.

## Purpose

The `repos` package manages a template repository system consisting of:

1. **Registry** - A git repository catalog of available templates with metadata
2. **Cache** - Local directory storing cloned template repositories in versioned subdirectories

It provides APIs for discovering, installing, and upgrading template repositories.

## Key Exports

### Singletons
- `Registry` - Global default registry instance
- `Cache` - Global default cache instance

### RepoID Functions
- `EnsureRepoID()` - Normalize to "name@version" format
- `SplitRepoID()` - Split into name and version
- `MakeRepoID()` - Construct repo ID
- `NameFromRepoID()`, `VersionFromRepoID()` - Extractors
- `IsRepoID()` - Check if string is valid repo ID

### Registry Methods
- `Load()`, `Save()` - Persist registry
- `List()`, `Search()`, `Get()` - Query templates
- `Update()`, `Reset()` - Sync with remote

### Cache Methods
- `List()`, `Search()` - Query cached templates
- `Install()` - Clone specific template version
- `Upgrade()`, `UpgradeAll()` - Update templates
- `Remove()`, `Clean()` - Cleanup
- `GetTemplateDir()` - Get local filesystem path

### High-level API
- `GetOrInstallTemplateFromRepoID()` - Install if not cached

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Cache/registry directories and URLs |
| `git` | Clone, pull, checkout, repo info |
| `helper` | File/directory operations |
| `log` | Logging |
