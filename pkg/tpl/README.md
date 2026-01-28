# tpl

Template creation and management operations.

## Purpose

The `tpl` package manages template operations for code generation. It provides functionality to create, inspect, and manage templates for multiple programming languages:

- C++
- Go
- Python
- TypeScript
- Rust
- Unreal Engine

## Key Exports

### Types
- `TemplateInfo` - Template metadata with Rules and Files list

### Functions
- `CreateCustomTemplate(dir, lang)` - Create template structure for a language
- `Info(dir)` - Read and return template information
- `PublishTemplate(dir)` - Publish template (placeholder)

### Supported Languages
Templates include `rules.yaml` configuration and language-specific template files from the `apigear-by-example` repository.

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | Path joining utilities |
| `log` | Logging |
