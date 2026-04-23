package net

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// echoServer upgrades to WebSocket and echoes what it receives.
func echoServer(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err := conn.WriteMessage(messageType, payload); err != nil {
			return
		}
	}
}

func TestWSProxy_BasicTextRoundTrip(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(echoServer))
	t.Cleanup(upstream.Close)

	upstreamWSURL := strings.Replace(upstream.URL, "http", "ws", 1)

	proxy, err := NewWSProxy(ProxyOptions{
		BasePath: "/ws",
		Routes: []RouteConfig{{
			Path:    "/:id",
			Param:   "id",
			Targets: map[string]string{"abc": upstreamWSURL},
			Mode:    MessageModeText,
		}},
	})
	if err != nil {
		t.Fatalf("create proxy: %v", err)
	}

	server := httptest.NewServer(proxy)
	t.Cleanup(server.Close)

	dialer := websocket.Dialer{}
	proxyURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/abc"
	clientConn, _, err := dialer.Dial(proxyURL, nil)
	if err != nil {
		t.Fatalf("dial proxy: %v", err)
	}
	t.Cleanup(func() {
		if err := clientConn.Close(); err != nil {
			t.Errorf("close client connection: %v", err)
		}
	})

	payload := map[string]string{"hello": "world"}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}
	if err := clientConn.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatalf("write message: %v", err)
	}

	messageType, resp, err := clientConn.ReadMessage()
	if err != nil {
		t.Fatalf("read message: %v", err)
	}
	if messageType != websocket.TextMessage {
		t.Fatalf("unexpected message type: got %d", messageType)
	}
	if string(resp) != string(data) {
		t.Fatalf("unexpected payload: got %s want %s", resp, data)
	}
}

func TestWSProxy_MiddlewareDrop(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(echoServer))
	t.Cleanup(upstream.Close)

	var received atomic.Bool

	upstreamWSURL := strings.Replace(upstream.URL, "http", "ws", 1)

	proxy, err := NewWSProxy(ProxyOptions{
		BasePath: "/ws",
		Routes: []RouteConfig{{
			Path:    "",
			Targets: map[string]string{"": upstreamWSURL},
			Mode:    MessageModeBinary,
		}},
	})
	if err != nil {
		t.Fatalf("create proxy: %v", err)
	}

	proxy.Use(func(ctx context.Context, msg *ProxyMessage) error {
		if msg.Direction == DirectionClientToUpstream {
			msg.Drop = true
		}
		return nil
	})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received.Store(true)
		proxy.ServeHTTP(w, r)
	}))
	t.Cleanup(server.Close)

	dialer := websocket.Dialer{}
	proxyURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	clientConn, _, err := dialer.Dial(proxyURL, nil)
	if err != nil {
		t.Fatalf("dial proxy: %v", err)
	}
	t.Cleanup(func() {
		if err := clientConn.Close(); err != nil {
			t.Errorf("close client connection: %v", err)
		}
	})

	if err := clientConn.WriteMessage(websocket.BinaryMessage, []byte("ignored")); err != nil {
		t.Fatalf("write message: %v", err)
	}

	if err := clientConn.SetReadDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
		t.Fatalf("set read deadline: %v", err)
	}
	_, _, err = clientConn.ReadMessage()
	if err == nil {
		t.Fatalf("expected read error due to dropped message")
	}
	if !received.Load() {
		t.Fatalf("expected HTTP handler to be invoked")
	}
}
