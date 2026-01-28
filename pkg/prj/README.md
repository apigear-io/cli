# prj

Project lifecycle management for APIGear projects.

## Purpose

The `prj` package handles creation, discovery, and management of APIGear projects. A project is a directory containing an `apigear/` subdirectory with configuration documents. The package provides:

- Project initialization with demo files
- Project discovery and reading
- Document management (modules, solutions, simulations)
- Project archiving/export
- Git-based project import
- Editor/IDE integration

## Key Exports

### Types
- `ProjectInfo` - Project with Name, Path, and Documents
- `DocumentInfo` - Document with Name, Path, Type
- `DemoType` - Enum for demo types (module, solution, scenario)

### Functions
- `OpenProject()` - Open existing project
- `InitProject()` - Initialize new project with demos
- `GetProjectInfo()` - Retrieve project information
- `CurrentProject()` - Get currently loaded project
- `RecentProjectInfos()` - List recently accessed projects
- `ReadProject()` - Parse project structure
- `ImportProject()` - Import from Git repository
- `PackProject()` - Export as tar.gz archive
- `AddDocument()` - Add new documents
- `OpenEditor()`, `OpenStudio()` - Launch external tools

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Editor preferences, recent entries |
| `git` | Git URL validation, cloning |
| `helper` | Path utilities, document detection |
| `log` | Logging |
| `vfs` | Demo template content |
