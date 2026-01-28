# sol

Solution execution orchestrator for code generation pipelines.

## Purpose

The `sol` package orchestrates solution builds by reading solution specifications and coordinating the code generation pipeline. It handles:

- Reading and parsing solution YAML files
- Parsing input files (YAML/JSON data or IDL specifications)
- Applying metadata overrides to system models
- Coordinating code generation through multiple targets
- File watching for development workflows

## Key Exports

### Types
- `Runner` - Main orchestrator managing solution execution tasks

### Runner Methods
- `NewRunner()` - Create new runner instance
- `HasTask()`, `TaskFiles()` - Query tasks
- `OnTask()` - Register hook for task events
- `RunSource()` - Execute solution from file path (with caching)
- `RunDoc()` - Execute pre-parsed solution document
- `WatchSource()`, `WatchDoc()` - Watch for changes and re-execute
- `StopWatch()` - Stop watching a file
- `Clear()` - Cancel all running tasks
- `ReadSolutionDoc()` - Read and parse solution YAML

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Build info for version |
| `gen` | Code generation engine |
| `git` | Git operations |
| `helper` | Path utilities, map operations |
| `idl` | IDL parsing |
| `log` | Logging |
| `model` | System and DataParser |
| `mon` | Monitoring |
| `net` | Network operations |
| `repos` | Template installation |
| `sim` | Simulation |
| `spec` | Solution document types |
| `tasks` | Task management |
