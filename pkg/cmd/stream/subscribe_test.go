package stream

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// dialSubscribe connects a WS client to the subscribe server at the given address.
func dialSubscribe(t *testing.T, addr string) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
	conn, _, err := dialer.Dial(addr, nil)
	if err != nil {
		t.Fatalf("failed to connect to subscribe server: %v", err)
	}
	return conn
}

func TestSubscribe_ReceivesMessages(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := make(chan string, 1)
	opts := subscribeOptions{
		Listen: "ws://localhost:0/ws",
		Format: "text",
		onReady: func(addr string) {
			ready <- addr
		},
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- runSubscribe(ctx, &stdout, &stderr, opts)
	}()

	// Wait for server to be ready
	var addr string
	select {
	case addr = <-ready:
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not become ready")
	}

	// Connect and send messages
	conn := dialSubscribe(t, addr)
	defer conn.Close()

	messages := []string{"hello", "world", "done"}
	for _, msg := range messages {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			t.Fatalf("failed to send message: %v", err)
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Give server time to process
	time.Sleep(100 * time.Millisecond)
	cancel()

	<-errCh

	out := stdout.String()
	for _, msg := range messages {
		if !strings.Contains(out, msg) {
			t.Errorf("expected %q in output, got: %s", msg, out)
		}
	}
	if !strings.Contains(stderr.String(), "Listening on") {
		t.Errorf("expected 'Listening on' on stderr, got: %s", stderr.String())
	}
}

func TestSubscribe_EchoMode(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := make(chan string, 1)
	opts := subscribeOptions{
		Listen: "ws://localhost:0/ws",
		Echo:   true,
		onReady: func(addr string) {
			ready <- addr
		},
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- runSubscribe(ctx, &stdout, &stderr, opts)
	}()

	var addr string
	select {
	case addr = <-ready:
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not become ready")
	}

	conn := dialSubscribe(t, addr)
	defer conn.Close()

	// Send a message
	if err := conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	// Read the echoed response
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, msg, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("failed to read echo response: %v", err)
	}
	if string(msg) != "ping" {
		t.Errorf("expected echoed 'ping', got: %s", string(msg))
	}

	// Give server time to write output
	time.Sleep(100 * time.Millisecond)
	cancel()

	<-errCh

	out := stdout.String()
	if !strings.Contains(out, "-> ping") {
		t.Errorf("expected '-> ping' in output, got: %s", out)
	}
	if !strings.Contains(out, "<- ping") {
		t.Errorf("expected '<- ping' in output, got: %s", out)
	}
}

func TestSubscribe_NoEchoByDefault(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := make(chan string, 1)
	opts := subscribeOptions{
		Listen: "ws://localhost:0/ws",
		Format: "text",
		onReady: func(addr string) {
			ready <- addr
		},
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- runSubscribe(ctx, &stdout, &stderr, opts)
	}()

	var addr string
	select {
	case addr = <-ready:
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not become ready")
	}

	conn := dialSubscribe(t, addr)
	defer conn.Close()

	if err := conn.WriteMessage(websocket.TextMessage, []byte("test")); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	// Try to read — should timeout because echo is off
	conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	_, _, err := conn.ReadMessage()
	if err == nil {
		t.Error("expected no echo response, but got one")
	}

	time.Sleep(100 * time.Millisecond)
	cancel()

	<-errCh

	// Message should still appear in stdout
	if !strings.Contains(stdout.String(), "test") {
		t.Errorf("expected 'test' in output, got: %s", stdout.String())
	}
}

func TestSubscribe_JSONFormat(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := make(chan string, 1)
	opts := subscribeOptions{
		Listen: "ws://localhost:0/ws",
		Format: "json",
		onReady: func(addr string) {
			ready <- addr
		},
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- runSubscribe(ctx, &stdout, &stderr, opts)
	}()

	var addr string
	select {
	case addr = <-ready:
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not become ready")
	}

	conn := dialSubscribe(t, addr)
	defer conn.Close()

	if err := conn.WriteMessage(websocket.TextMessage, []byte("json-test")); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	time.Sleep(100 * time.Millisecond)
	cancel()

	<-errCh

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) < 1 {
		t.Fatal("expected at least one JSON line")
	}

	var om outputMessage
	if err := json.Unmarshal([]byte(lines[0]), &om); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if om.Data != "json-test" {
		t.Errorf("expected data 'json-test', got %q", om.Data)
	}
	if om.Type != "text" {
		t.Errorf("expected type 'text', got %q", om.Type)
	}
	if om.Timestamp == "" {
		t.Error("expected non-empty timestamp")
	}
}

func TestSubscribe_Cancellation(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())

	ready := make(chan string, 1)
	opts := subscribeOptions{
		Listen: "ws://localhost:0/ws",
		Format: "text",
		onReady: func(addr string) {
			ready <- addr
		},
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- runSubscribe(ctx, &stdout, &stderr, opts)
	}()

	select {
	case <-ready:
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not become ready")
	}

	// Cancel immediately
	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("subscribe server did not shut down")
	}

	if !strings.Contains(stderr.String(), "Stopped") {
		t.Errorf("expected 'Stopped' on stderr, got: %s", stderr.String())
	}
}
