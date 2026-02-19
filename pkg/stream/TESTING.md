# WebSocket Testing Guide

This guide explains how to test WebSocket functionality and the different proxy modes in the Stream module.

## Testing Strategy

### 1. Unit Tests
Test individual components in isolation using mocks.

### 2. Integration Tests
Test WebSocket connections with real servers and clients.

### 3. E2E Tests
Test complete scenarios from frontend to backend.

## WebSocket Test Patterns

### Pattern 1: Test Server Helper

Create a test WebSocket server that echoes messages:

```go
func testWSServer(t *testing.T) *httptest.Server {
    upgrader := websocket.Upgrader{}

    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            return
        }
        defer conn.Close()

        for {
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
            conn.WriteMessage(msgType, msg) // Echo back
        }
    })

    return httptest.NewServer(handler)
}
```

### Pattern 2: Test Client Helper

Create a WebSocket client for testing:

```go
func testWSClient(t *testing.T, url string) *websocket.Conn {
    wsURL := strings.Replace(url, "http://", "ws://", 1)
    conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    require.NoError(t, err)
    return conn
}
```

### Pattern 3: Async Proxy Testing

Start proxy in background, test, then clean up:

```go
func TestProxy(t *testing.T) {
    proxy := New("test", config, nil)

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        _ = proxy.Start(ctx)
    }()

    time.Sleep(100 * time.Millisecond) // Wait for startup

    // ... test operations ...

    cancel() // Graceful shutdown
}
```

## Testing Each Proxy Mode

### 1. Echo Mode

Echo mode reflects all messages back to the sender.

**Test Checklist:**
- ✅ Connection establishment
- ✅ Text message echo
- ✅ Binary message echo
- ✅ Multiple clients simultaneously
- ✅ Message order preservation
- ✅ Graceful disconnection

**Example:**
```go
func TestEchoMode(t *testing.T) {
    config := &Config{
        Name:    "test-echo",
        Listen:  "localhost:0",
        Mode:    ModeEcho,
        Enabled: true,
    }

    proxy := New("test-echo", config, nil)

    // Start proxy
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go proxy.Start(ctx)
    time.Sleep(100 * time.Millisecond)

    // Connect client
    client := testWSClient(t, proxy.GetListenAddr())
    defer client.Close()

    // Send and verify echo
    testMsg := []byte(`{"test":"echo"}`)
    client.WriteMessage(websocket.TextMessage, testMsg)

    _, received, _ := client.ReadMessage()
    assert.Equal(t, testMsg, received)
}
```

### 2. Proxy Mode (Forwarding)

Proxy mode forwards messages between frontend and backend.

**Test Checklist:**
- ✅ Message forwarding frontend → backend
- ✅ Message forwarding backend → frontend
- ✅ Bidirectional communication
- ✅ Connection tracking
- ✅ Backend reconnection
- ✅ Error handling

**Example:**
```go
func TestProxyMode(t *testing.T) {
    // Setup backend
    backend := testWSServer(t)
    defer backend.Close()

    // Setup proxy
    config := &Config{
        Name:    "test-proxy",
        Listen:  "localhost:0",
        Backend: strings.Replace(backend.URL, "http://", "ws://", 1),
        Mode:    ModeProxy,
        Enabled: true,
    }

    proxy := New("test-proxy", config, nil)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go proxy.Start(ctx)
    time.Sleep(100 * time.Millisecond)

    // Connect client to proxy
    client := testWSClient(t, proxy.GetListenAddr())
    defer client.Close()

    // Send through proxy, verify backend echoes
    testMsg := []byte(`{"proxied":"message"}`)
    client.WriteMessage(websocket.TextMessage, testMsg)

    _, received, _ := client.ReadMessage()
    assert.Equal(t, testMsg, received) // Backend echoed it
}
```

### 3. Backend Mode

Backend mode acts as an ObjectLink server.

**Test Checklist:**
- ✅ ObjectLink LINK messages
- ✅ ObjectLink INIT responses
- ✅ ObjectLink INVOKE handling
- ✅ Property change notifications
- ✅ Signal emission
- ✅ Multiple object support

**Example:**
```go
func TestBackendMode(t *testing.T) {
    // Requires scripting integration
    t.Skip("Backend mode requires scripting engine")

    config := &Config{
        Name:    "test-backend",
        Listen:  "localhost:0",
        Mode:    ModeBackend,
        Enabled: true,
    }

    proxy := New("test-backend", config, nil)
    // TODO: Set up script with object definitions

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go proxy.Start(ctx)
    time.Sleep(100 * time.Millisecond)

    client := testWSClient(t, proxy.GetListenAddr())
    defer client.Close()

    // Send LINK message
    linkMsg := []interface{}{10, "demo.Counter"}
    data, _ := json.Marshal(linkMsg)
    client.WriteMessage(websocket.TextMessage, data)

    // Expect INIT response
    _, response, _ := client.ReadMessage()
    var initMsg []interface{}
    json.Unmarshal(response, &initMsg)

    assert.Equal(t, float64(11), initMsg[0]) // INIT = 11
    assert.Equal(t, "demo.Counter", initMsg[1])
}
```

### 4. Inbound-Only Mode

Inbound-only accepts connections but doesn't forward or respond.

**Test Checklist:**
- ✅ Connection acceptance
- ✅ Message reception
- ✅ No responses sent
- ✅ Message logging/tracing
- ✅ Connection tracking

**Example:**
```go
func TestInboundOnlyMode(t *testing.T) {
    config := &Config{
        Name:    "test-inbound",
        Listen:  "localhost:0",
        Mode:    ModeInboundOnly,
        Enabled: true,
    }

    proxy := New("test-inbound", config, nil)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go proxy.Start(ctx)
    time.Sleep(100 * time.Millisecond)

    client := testWSClient(t, proxy.GetListenAddr())
    defer client.Close()

    // Send message (should be accepted)
    client.WriteMessage(websocket.TextMessage, []byte(`{"test":"data"}`))

    // Set timeout for read
    client.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

    // Should timeout (no response in inbound-only mode)
    _, _, err := client.ReadMessage()
    assert.Error(t, err) // Timeout expected
}
```

## Testing ObjectLink Protocol

### Testing Message Types

```go
func TestObjectLinkMessages(t *testing.T) {
    // Test each message type
    messages := []struct{
        name string
        data []interface{}
    }{
        {"LINK", []interface{}{10, "demo.Counter"}},
        {"INIT", []interface{}{11, "demo.Counter", map[string]interface{}{"count": 0}}},
        {"INVOKE", []interface{}{30, 1, "demo.Counter/increment", []interface{}{}}},
        {"SIGNAL", []interface{}{40, "demo.Counter/changed", []interface{}{5}}},
        {"PROPERTY_CHANGE", []interface{}{50, "demo.Counter/count", 42}},
    }

    for _, msg := range messages {
        t.Run(msg.name, func(t *testing.T) {
            // Test sending this message type
            data, _ := json.Marshal(msg.data)
            // ... send and verify
        })
    }
}
```

## Testing Scenarios

### Scenario 1: Multiple Clients

Test that multiple clients can connect simultaneously:

```go
func TestMultipleClients(t *testing.T) {
    proxy := setupEchoProxy(t)
    defer proxy.Close()

    // Connect 10 clients
    clients := make([]*websocket.Conn, 10)
    for i := 0; i < 10; i++ {
        clients[i] = testWSClient(t, proxy.GetListenAddr())
        defer clients[i].Close()
    }

    // Each client sends a unique message
    for i, client := range clients {
        msg := []byte(fmt.Sprintf(`{"client":%d}`, i))
        client.WriteMessage(websocket.TextMessage, msg)

        _, received, _ := client.ReadMessage()
        assert.Equal(t, msg, received)
    }
}
```

### Scenario 2: High Throughput

Benchmark message throughput:

```go
func BenchmarkMessageThroughput(b *testing.B) {
    proxy := setupEchoProxy(b)
    defer proxy.Close()

    client := testWSClient(b, proxy.GetListenAddr())
    defer client.Close()

    msg := []byte(`{"benchmark":"test"}`)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.WriteMessage(websocket.TextMessage, msg)
        client.ReadMessage()
    }
}
```

### Scenario 3: Connection Loss and Reconnect

Test reconnection behavior:

```go
func TestReconnection(t *testing.T) {
    t.Skip("Reconnection not yet implemented")

    // 1. Start backend
    backend := testWSServer(t)

    // 2. Start proxy
    proxy := setupProxyMode(t, backend.URL)

    // 3. Connect client
    client := testWSClient(t, proxy.GetListenAddr())

    // 4. Stop backend
    backend.Close()

    // 5. Verify proxy detects disconnect
    time.Sleep(500 * time.Millisecond)
    // assert proxy is in reconnecting state

    // 6. Restart backend
    backend = testWSServer(t)
    defer backend.Close()

    // 7. Verify proxy reconnects
    time.Sleep(1 * time.Second)
    // assert proxy is connected

    client.Close()
}
```

### Scenario 4: Message Ordering

Verify messages are processed in order:

```go
func TestMessageOrdering(t *testing.T) {
    proxy := setupEchoProxy(t)
    defer proxy.Close()

    client := testWSClient(t, proxy.GetListenAddr())
    defer client.Close()

    // Send 100 numbered messages
    for i := 0; i < 100; i++ {
        msg := []byte(fmt.Sprintf(`{"seq":%d}`, i))
        client.WriteMessage(websocket.TextMessage, msg)
    }

    // Verify order
    for i := 0; i < 100; i++ {
        _, received, _ := client.ReadMessage()
        var data map[string]int
        json.Unmarshal(received, &data)
        assert.Equal(t, i, data["seq"])
    }
}
```

## Running Tests

### Run all WebSocket tests:
```bash
go test ./pkg/stream/proxy -v
```

### Run specific test:
```bash
go test ./pkg/stream/proxy -v -run TestEchoMode
```

### Run benchmarks:
```bash
go test ./pkg/stream/proxy -bench=. -benchmem
```

### Run with race detector:
```bash
go test ./pkg/stream/proxy -race
```

### Run with coverage:
```bash
go test ./pkg/stream/proxy -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Common Pitfalls

### 1. Not Waiting for Proxy Startup
Always add a small sleep after starting the proxy:
```go
go proxy.Start(ctx)
time.Sleep(100 * time.Millisecond) // Let it start
```

### 2. Forgetting to Close Connections
Use defer to ensure cleanup:
```go
client := testWSClient(t, url)
defer client.Close() // Always clean up
```

### 3. Read Timeouts
Set deadlines when expecting no response:
```go
client.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
_, _, err := client.ReadMessage()
// err will be timeout error if no message
```

### 4. Context Cancellation
Always cancel contexts to avoid goroutine leaks:
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Ensure cleanup
```

### 5. Random Ports
Use port 0 for tests to avoid conflicts:
```go
Listen: "localhost:0" // OS assigns random port
```

Get actual address with:
```go
actualAddr := proxy.GetListenAddr()
```

## Debugging Tips

### 1. Enable Verbose Logging
```go
// Add to test:
t.Setenv("LOG_LEVEL", "debug")
```

### 2. Use Test Helpers
```go
func setupEchoProxy(t *testing.T) *Proxy {
    config := &Config{
        Name:    t.Name(),
        Listen:  "localhost:0",
        Mode:    ModeEcho,
        Enabled: true,
    }

    proxy := New(t.Name(), config, nil)
    ctx, cancel := context.WithCancel(context.Background())

    t.Cleanup(func() {
        cancel()
        time.Sleep(50 * time.Millisecond)
    })

    go proxy.Start(ctx)
    time.Sleep(100 * time.Millisecond)

    return proxy
}
```

### 3. Capture Messages
```go
var receivedMessages [][]byte

go func() {
    for {
        _, msg, err := client.ReadMessage()
        if err != nil {
            return
        }
        receivedMessages = append(receivedMessages, msg)
    }
}()
```

### 4. Test with Real ObjectLink Client

Use objectlink-core-go for integration tests:
```go
import "github.com/apigear-io/objectlink-core-go"

func TestWithObjectLinkClient(t *testing.T) {
    proxy := setupBackendProxy(t)

    client := objectlink.NewClient("test", proxy.GetListenAddr(), []string{"demo.Counter"}, true, true)
    defer client.Stop()

    err := client.Start()
    require.NoError(t, err)

    // Use client to interact with backend
}
```

## Next Steps

1. **Implement missing features**: GetStats(), reconnection logic
2. **Add more test scenarios**: Large messages, malformed data, stress tests
3. **E2E tests**: Test with real frontend and backend applications
4. **Performance tests**: Measure latency, throughput under load
5. **Integration tests**: Test with objectlink-core-go clients/servers
