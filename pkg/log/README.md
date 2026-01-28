# log

Structured logging system for the CLI application.

## Purpose

The `log` package provides multi-destination structured logging using zerolog. It supports:

- Console output with configurable log levels
- Rolling file logging to `~/.apigear/apigear.log`
- Event emission for external system integration
- Log level control via `DEBUG` environment variable (1=debug, 2=trace)
- Automatic UUID tagging for log entries
- Topic-based logging for component isolation

## Key Exports

- `Debug()`, `Info()`, `Warn()`, `Error()`, `Fatal()`, `Panic()` - Log level shortcuts
- `Topic(topic string)` - Create logger with topic label
- `OnReportEvent()` - Register callback for parsed log events
- `OnReportBytes()` - Register callback for raw log bytes
- `UUIDHook` - Zerolog hook adding unique IDs
- `EventLogWriter` - Custom writer for event emission

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Config directory for log file path |
| `helper` | UUID generation and path joining |
