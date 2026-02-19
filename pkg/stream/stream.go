// Package stream provides WebSocket streaming and proxy functionality for ApiGear CLI.
//
// The stream package integrates WebSocket proxy capabilities with ObjectLink protocol support,
// message tracing, JavaScript scripting, and real-time monitoring.
//
// Key features:
//   - WebSocket proxy with multiple modes (proxy, echo, backend, inbound-only)
//   - ObjectLink protocol client management
//   - Message tracing and replay (JSONL format)
//   - JavaScript-based custom backends and message transformation
//   - Real-time monitoring and statistics
//
// Package structure:
//   - relay: WebSocket infrastructure (connections, clients, servers, hub)
//   - protocol: ObjectLink message parsing for logging/UI (best-effort)
//   - config: Configuration management (YAML/JSON loading, watching)
//   - proxy: Core WebSocket proxy implementation
//   - client: ObjectLink client management (using objectlink-core-go)
//   - scripting: JavaScript engine (Goja runtime)
//   - tracing: Trace file management (reader, writer, player, filter)
package stream

import (
	"fmt"
)

// Version information
const (
	Version = "0.1.0"
	Name    = "ApiGear Stream"
)

// String returns the version string
func String() string {
	return fmt.Sprintf("%s v%s", Name, Version)
}
