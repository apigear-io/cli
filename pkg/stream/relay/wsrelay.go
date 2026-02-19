// Package wsrelay provides generic WebSocket relay infrastructure.
//
// This package includes low-level connection management, message streaming,
// pub/sub hubs, and client abstractions for building WebSocket-based systems.
//
// The package is organized into three main areas:
//   - Connection infrastructure (Connection, ConnectionPool, LifecycleManager)
//   - Messaging (Hub, RingBuffer, Forwarder with delay/throttle strategies)
//   - Client abstractions (Client interface, Registry, EventHub)
//
// All implementations are thread-safe and suitable for concurrent use.
package relay
