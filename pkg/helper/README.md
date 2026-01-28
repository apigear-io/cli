# helper

Utility package providing reusable helper functions and generic types.

## Purpose

The `helper` package is a foundational utility library used across the CLI application. It provides:

- **File Operations**: Path manipulation, file/directory checking, copying, reading/writing documents
- **Generic Data Structures**: Iterator, Emitter, Hook for event handling
- **String Utilities**: Case-insensitive matching, abbreviations, transformations
- **Document Parsing**: JSON/YAML parsing, NDJSON scanning, format conversion
- **HTTP Utilities**: HTTPSender for JSON serialization, POST helpers
- **ID Generation**: UUID generation, integer ID generators
- **Concurrency**: Signal handling, timed iteration, sender control

## Key Exports

- `Iterator[T]`, `Emitter[T]`, `Hook[T]` - Generic patterns
- `Join()`, `IsDir()`, `IsFile()`, `CopyFile()`, `CopyDir()` - File operations
- `ReadDocument()`, `WriteDocument()` - YAML/JSON I/O
- `ParseJson()`, `ParseYaml()`, `YamlToJson()` - Parsing
- `NewUUID()`, `MakeIdGenerator()` - ID generation
- `GetFreePort()`, `WaitForInterrupt()` - System utilities
- `HTTPSender`, `HttpPost()` - HTTP operations

## Dependencies

This package has no dependencies on other `pkg/` packages.
