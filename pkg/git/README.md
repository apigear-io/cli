# git

Git repository operations abstraction layer.

## Purpose

The `git` package provides high-level functionality for Git repository operations. It wraps the `go-git` library to offer simplified APIs for:

- Cloning and pulling repositories
- Checking out specific commits or tags
- Retrieving repository metadata and version information
- Parsing and validating Git URLs
- Managing semantic versions from tags

## Key Exports

- `RepoInfo` - Repository metadata (name, path, URL, commit, version)
- `VersionInfo` - Semantic version information
- `VersionCollection` - Sortable collection of versions
- `Clone()`, `CloneOrPull()`, `Pull()` - Repository sync operations
- `CheckoutCommit()`, `CheckoutTag()` - Version switching
- `LocalRepoInfo()`, `RemoteRepoInfo()` - Metadata extraction
- `GetTagsFromRepo()`, `GetTagsFromRemote()` - Version listing
- `IsValidGitUrl()`, `ParseAsUrl()` - URL utilities

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | Directory checking (IsDir) |
| `log` | Logging |
