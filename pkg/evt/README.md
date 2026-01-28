# evt

Event-driven messaging system built on NATS.

## Purpose

The `evt` package provides an event bus abstraction for publish/subscribe and request/response patterns. It enables asynchronous communication between components using NATS as the messaging backend.

Features:
- Event publishing without waiting for response
- Request/response pattern with 10-second timeout
- Handler registration for specific event types
- Middleware support for event processing

## Key Exports

- `Event` - Message struct with Kind, Value, Error, and Meta fields
- `IEventBus` - Interface for event operations (Publish, Request, Register, Use)
- `NewEvent()`, `NewErrorEvent()` - Event constructors
- `NewNatsEventBus()` - Creates NATS-backed event bus
- `HandlerFunc` - Function type for event handlers

## Dependencies

This package has no dependencies on other `pkg/` packages.
