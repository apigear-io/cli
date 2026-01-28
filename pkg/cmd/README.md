# cmd

CLI command layer for the APIGear application using the Cobra framework.

## Purpose

The `cmd` package serves as the entry point for all user-facing CLI commands. It orchestrates the various CLI subcommands and delegates to specialized domain packages for actual functionality. The package exposes commands for:

- Code generation (`gen`)
- Project management (`prj`)
- Template management (`tpl`)
- Monitoring (`mon`)
- Simulation (`sim`)
- Specification handling (`spec`)
- Configuration (`cfg`)
- MCP server (`mcp`)

## Key Exports

- `Run()` - Main entry point for CLI execution
- `NewRootCommand()` - Creates the root Cobra command
- `NewServeCommand()` - Starts the APIGear server
- `NewVersionCommand()` - Displays version info
- `NewUpdateCommand()` - CLI self-update
- `NewMCPCommand()` - Starts MCP server

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Build info and configuration |
| `gen` | Code generation engine |
| `git` | Git operations |
| `helper` | Utility functions |
| `idl` | IDL parsing |
| `log` | Logging |
| `mcp` | MCP server |
| `model` | API models |
| `mon` | Monitoring |
| `net` | Network management |
| `prj` | Project management |
| `repos` | Template repositories |
| `sim` | Simulation |
| `sol` | Solution runner |
| `spec` | Specifications |
| `tasks` | Task execution |
| `tpl` | Template operations |
| `up` | Self-update |
| `vfs` | Virtual filesystem |
