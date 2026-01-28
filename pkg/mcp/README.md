# mcp

Model Context Protocol (MCP) server for AI tool integration.

## Purpose

The `mcp` package implements an MCP server that exposes CLI operations as tools for Claude and other AI assistants. It enables programmatic access to:

- **Code Generation**: Generate SDKs from solution documents or with expert mode options
- **Specification Validation**: Check module, solution, and rules files
- **Template Management**: List and update templates from the registry
- **Schema Access**: Output JSON/YAML schemas for specification documents

## Key Exports

- `RunMCPServer()` - Initialize and run the MCP server via stdio

### MCP Tools Registered

- `generateSolution` - Generate SDKs from solution documents
- `generateExpert` - Advanced generation with fine-grained options
- `specificationCheck` - Validate specification files
- `specificationSchema` - Output specification schemas
- `templateList` - List available templates
- `templateUpdate` - Update template registry
- `version` - Display version information

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Build info for versioning |
| `cmd/gen` | Code generation commands |
| `cmd/tpl` | Template commands |
| `gen` | Code generation engine |
| `git` | Git operations |
| `helper` | Utilities |
| `idl` | IDL parsing |
| `log` | Logging |
| `model` | API models |
| `mon` | Monitoring |
| `net` | Network operations |
| `repos` | Template repositories |
| `sim` | Simulation |
| `sol` | Solution runner |
| `spec` | Specification validation |
| `tasks` | Task execution |
| `tpl` | Template operations |
