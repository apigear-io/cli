package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/stream/config"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testWSServer creates a test WebSocket server that echoes messages
func testWSServer(t *testing.T) *httptest.Server {
	upgrader := websocket.Upgrader{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("Upgrade error: %v", err)
			return
		}
		defer conn.Close()

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Echo back the message
			if err := conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	return httptest.NewServer(handler)
}

// testWSClient creates a WebSocket client and connects to the given URL
func testWSClient(t *testing.T, url string) *websocket.Conn {
	// Convert http:// to ws://
	wsURL := strings.Replace(url, "http://", "ws://", 1)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "Failed to connect WebSocket client")

	return conn
}

// TestEchoMode tests the echo proxy mode
func TestEchoMode(t *testing.T) {
	// Create echo proxy
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("test-echo", "ws://localhost:0/ws", "", cfg)
	require.NotNil(t, proxy)
	proxy.SetTrace(false) // Disable tracing for tests

	// Start proxy in background
	errCh := make(chan error, 1)
	go func() {
		errCh <- proxy.Start()
	}()

	// Wait for proxy to be ready
	time.Sleep(100 * time.Millisecond)

	// Get the actual listen address
	listenAddr := proxy.GetListenAddr()
	require.NotEmpty(t, listenAddr, "Proxy listen address should not be empty")

	// Connect client to proxy
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "Failed to connect to echo proxy")
	defer client.Close()

	// Send a test message
	testMsg := []byte(`{"type":"test","data":"hello"}`)
	err = client.WriteMessage(websocket.TextMessage, testMsg)
	require.NoError(t, err, "Failed to send message")

	// Read echoed message
	msgType, msg, err := client.ReadMessage()
	require.NoError(t, err, "Failed to read echoed message")
	assert.Equal(t, websocket.TextMessage, msgType)
	assert.Equal(t, testMsg, msg)

	// Stop proxy
	err = proxy.Stop()
	assert.NoError(t, err, "Proxy should stop cleanly")

	// Wait for Start() to return
	select {
	case err := <-errCh:
		// Start() should return nil or http.ErrServerClosed
		if err != nil && err != http.ErrServerClosed {
			t.Errorf("Unexpected error from Start(): %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Proxy did not stop in time")
	}
}

// TestProxyMode tests the proxy forwarding mode
func TestProxyMode(t *testing.T) {
	// Create backend server
	backend := testWSServer(t)
	defer backend.Close()

	// Create proxy in forwarding mode
	backendURL := strings.Replace(backend.URL, "http://", "ws://", 1)
	cfg := config.ProxyConfig{
		Mode: "proxy",
	}

	proxy := NewProxy("test-proxy", "ws://localhost:0/ws", backendURL, cfg)
	require.NotNil(t, proxy)
	proxy.SetTrace(false)

	// Start proxy
	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	// Connect client to proxy
	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "Failed to connect to proxy")
	defer client.Close()

	// Send message through proxy
	testMsg := []byte(`{"type":"test","data":"proxied"}`)
	err = client.WriteMessage(websocket.TextMessage, testMsg)
	require.NoError(t, err, "Failed to send message")

	// Read response (backend echoes it back)
	msgType, msg, err := client.ReadMessage()
	require.NoError(t, err, "Failed to read response")
	assert.Equal(t, websocket.TextMessage, msgType)
	assert.Equal(t, testMsg, msg)

	// Stop proxy
	err = proxy.Stop()
	assert.NoError(t, err)
	time.Sleep(100 * time.Millisecond)
}

// TestInboundOnlyMode tests inbound-only mode (accepts connections but doesn't forward)
func TestInboundOnlyMode(t *testing.T) {
	cfg := config.ProxyConfig{
		Mode: "inbound-only",
	}

	proxy := NewProxy("test-inbound", "ws://localhost:0/ws", "", cfg)
	require.NotNil(t, proxy)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	// Connect client
	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "Failed to connect to inbound-only proxy")
	defer client.Close()

	// Send message (should be accepted but not echoed/forwarded)
	testMsg := []byte(`{"type":"test"}`)
	err = client.WriteMessage(websocket.TextMessage, testMsg)
	require.NoError(t, err, "Failed to send message")

	// Set read deadline to avoid blocking forever
	client.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

	// Try to read (should timeout since inbound-only doesn't respond)
	_, _, err = client.ReadMessage()
	assert.Error(t, err, "Should timeout waiting for response in inbound-only mode")

	err = proxy.Stop()
	assert.NoError(t, err)
}

// TestMultipleClients tests multiple simultaneous WebSocket connections
func TestMultipleClients(t *testing.T) {
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("test-multi", "ws://localhost:0/ws", "", cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)

	// Connect multiple clients
	numClients := 5
	clients := make([]*websocket.Conn, numClients)
	for i := 0; i < numClients; i++ {
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		require.NoError(t, err, "Failed to connect client %d", i)
		clients[i] = conn
		defer conn.Close()
	}

	// Each client sends and receives a message
	for i, client := range clients {
		msg := []byte(fmt.Sprintf(`{"client":%d}`, i))
		err := client.WriteMessage(websocket.TextMessage, msg)
		require.NoError(t, err, "Client %d failed to send", i)

		_, received, err := client.ReadMessage()
		require.NoError(t, err, "Client %d failed to receive", i)
		assert.Equal(t, msg, received, "Client %d received wrong message", i)
	}

	err := proxy.Stop()
	assert.NoError(t, err)
}

// TestObjectLinkMessages tests ObjectLink protocol message handling
func TestObjectLinkMessages(t *testing.T) {
	backend := testWSServer(t)
	defer backend.Close()

	backendURL := strings.Replace(backend.URL, "http://", "ws://", 1)
	cfg := config.ProxyConfig{
		Mode: "proxy",
	}

	proxy := NewProxy("test-objectlink", "ws://localhost:0/ws", backendURL, cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer client.Close()

	// Test ObjectLink LINK message
	linkMsg := []interface{}{10, "demo.Counter"}
	linkData, _ := json.Marshal(linkMsg)

	err = client.WriteMessage(websocket.TextMessage, linkData)
	require.NoError(t, err)

	// Read response
	_, response, err := client.ReadMessage()
	require.NoError(t, err)

	// Verify it's the same message (backend echoes)
	var responseMsg []interface{}
	err = json.Unmarshal(response, &responseMsg)
	require.NoError(t, err)
	assert.Equal(t, float64(10), responseMsg[0]) // JSON numbers are float64
	assert.Equal(t, "demo.Counter", responseMsg[1])

	// Test INVOKE message
	invokeMsg := []interface{}{30, 1, "demo.Counter/increment", []interface{}{}}
	invokeData, _ := json.Marshal(invokeMsg)

	err = client.WriteMessage(websocket.TextMessage, invokeData)
	require.NoError(t, err)

	_, response, err = client.ReadMessage()
	require.NoError(t, err)
	assert.NotEmpty(t, response)

	err = proxy.Stop()
	assert.NoError(t, err)
}

// TestProxyReconnect tests backend reconnection logic
func TestProxyReconnect(t *testing.T) {
	t.Skip("Reconnection logic not yet implemented")

	// TODO: Test that proxy reconnects to backend when connection drops
	// 1. Start backend
	// 2. Start proxy pointing to backend
	// 3. Stop backend
	// 4. Verify proxy detects disconnection
	// 5. Restart backend
	// 6. Verify proxy reconnects
}

// TestProxyStats tests statistics collection
func TestProxyStats(t *testing.T) {
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("test-stats", "ws://localhost:0/ws", "", cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	// Connect and send messages
	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer client.Close()

	// Send several messages
	for i := 0; i < 3; i++ {
		msg := []byte(fmt.Sprintf(`{"count":%d}`, i))
		err = client.WriteMessage(websocket.TextMessage, msg)
		require.NoError(t, err)

		_, _, err = client.ReadMessage()
		require.NoError(t, err)
	}

	// Verify stats were collected
	info := proxy.Info()
	assert.Greater(t, info.MessagesReceived, int64(0), "Should have received messages")
	assert.Greater(t, info.MessagesSent, int64(0), "Should have sent messages")

	err = proxy.Stop()
	assert.NoError(t, err)
}

// TestBinaryMessages tests binary WebSocket message handling
func TestBinaryMessages(t *testing.T) {
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("test-binary", "ws://localhost:0/ws", "", cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer client.Close()

	// Send binary message
	binaryData := []byte{0x01, 0x02, 0x03, 0x04, 0xFF, 0xFE}
	err = client.WriteMessage(websocket.BinaryMessage, binaryData)
	require.NoError(t, err)

	// Read echoed binary message
	msgType, received, err := client.ReadMessage()
	require.NoError(t, err)
	assert.Equal(t, websocket.BinaryMessage, msgType)
	assert.Equal(t, binaryData, received)

	err = proxy.Stop()
	assert.NoError(t, err)
}

// TestProxyClose tests graceful shutdown
func TestProxyClose(t *testing.T) {
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("test-close", "ws://localhost:0/ws", "", cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	// Connect client
	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)

	// Close proxy
	err = proxy.Stop()
	assert.NoError(t, err)
	time.Sleep(100 * time.Millisecond)

	// Client should detect connection closed - first write might succeed due to buffering,
	// but read should definitely fail
	client.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	_, _, err = client.ReadMessage()
	assert.Error(t, err, "Read should fail after proxy closes")

	client.Close()
}

// BenchmarkEchoMode benchmarks echo mode throughput
func BenchmarkEchoMode(b *testing.B) {
	cfg := config.ProxyConfig{
		Mode: "echo",
	}

	proxy := NewProxy("bench-echo", "ws://localhost:0/ws", "", cfg)
	proxy.SetTrace(false)

	go func() {
		_ = proxy.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	listenAddr := proxy.GetListenAddr()
	wsURL := fmt.Sprintf("ws://%s/ws", listenAddr)
	client, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(b, err)
	defer client.Close()

	testMsg := []byte(`{"type":"benchmark","data":"test message"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = client.WriteMessage(websocket.TextMessage, testMsg)
		if err != nil {
			b.Fatal(err)
		}

		_, _, err = client.ReadMessage()
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()

	_ = proxy.Stop()
}
