# mon

Monitoring and event tracking system for API activity.

## Purpose

The `mon` package enables recording and processing of API events including calls, signals, and state changes. It provides:

- Event creation and sanitization
- Multiple input formats (CSV, NDJSON)
- JavaScript-based event generation scripts
- Event emission via hooks

## Key Exports

### Types
- `Event` - Monitored API event with Id, Source, Type, Timestamp, Symbol, Data
- `EventFactory` - Factory for creating and sanitizing events
- `EventScript` - JavaScript runtime for event generation

### Constants
- `TypeCall`, `TypeSignal`, `TypeState` - Event type constants

### Functions
- `MakeEvent()`, `MakeCall()`, `MakeSignal()`, `MakeState()` - Event constructors
- `ReadCsvEvents()` - Parse events from CSV files
- `ReadJsonEvents()` - Parse NDJSON event streams
- `Emitter` - Global Hook for event emission

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | Hook pattern for event emission |
| `log` | Logging |
