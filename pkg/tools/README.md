# tools

Low-level utility tools and helper components.

## Purpose

The `tools` package provides foundational utility components. Currently contains:

- **Hook[T]** - Generic thread-safe event hook system with handler registration
- **ColorWriter** - Colored stderr output for error messages

> **Note**: The `Hook[T]` implementation in this package may be superseded by the version in `pkg/helper`. Most of the codebase uses `helper.Hook` instead.

## Key Exports

### Hook[T]
- `NewHook[T]()` - Create new hook instance
- `Add()` - Register handler, returns unsubscribe function
- `PreAdd()` - Add handler to beginning (higher priority)
- `Fire()` - Fire all handlers with event
- `Connect()` - Chain hooks together
- `Clear()` - Remove all handlers
- `Len()` - Handler count

### ColorWriter
- `NewErrWriter()` - Create writer for red-colored stderr output

## Dependencies

This package has no dependencies on other `pkg/` packages.
