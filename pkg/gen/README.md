# gen

Code generation engine for transforming API specifications into source code.

## Purpose

The `gen` package is the core code generation engine that transforms API specifications into source code across multiple programming languages. It works by:

1. Parsing template rules documents (YAML/JSON specs)
2. Reading Go text templates from a template directory
3. Applying templates to API models (systems, modules, interfaces, structs, enums)
4. Writing generated code to an output directory

Features:
- Multi-language support via template filters (C++, Go, Java, Python, TypeScript, Rust, Qt, Unreal Engine)
- Feature-based generation with configurable options
- Dry-run mode for previewing changes
- Generation statistics and reporting

## Key Exports

- `Generator` - Main generator struct via `New()` constructor
- `Options` - Configuration for output, templates, features
- `GeneratorStats` - Tracks generation metrics
- `ProcessRules()` - Main entry point for code generation
- `RenderString()` - Template string rendering utility

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `git` | Git operations for templates |
| `helper` | File operations and utilities |
| `idl` | IDL parsing |
| `log` | Logging |
| `model` | API data models |
| `mon` | Monitoring |
| `net` | Network operations |
| `repos` | Template repository management |
| `sim` | Simulation engine |
| `spec` | Rules document types |
