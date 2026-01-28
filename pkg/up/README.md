# up

Self-update manager for the CLI application.

## Purpose

The `up` package provides functionality to check GitHub repositories for new releases and automatically update the current executable. It wraps the `go-selfupdate` library to provide:

- Version checking against GitHub releases
- Automatic executable update with checksum validation
- Symlink resolution for proper update paths

## Key Exports

### Types
- `Updater` - Wrapper struct managing the self-update process

### Functions
- `NewUpdater(repo, version)` - Create new updater for a GitHub repository
- `Check(ctx)` - Check GitHub for new releases, returns Release if update available
- `Update(ctx, release)` - Apply update to current executable

### Features
- Uses `checksums.txt` for update validation
- Resolves symlinks to find actual executable path
- Context-aware for cancellation support

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | File existence checking |
| `log` | Logging |
