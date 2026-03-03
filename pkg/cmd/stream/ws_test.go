package stream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// --- Test helpers ---

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// testEchoServer echoes all messages back to the client.
func testEchoServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if err := conn.WriteMessage(mt, msg); err != nil {
				return
			}
		}
	}))
}

// testPushServer sends predefined messages to every connected client, then closes.
func testPushServer(t *testing.T, messages []string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for _, msg := range messages {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}))
}

// testSilentServer accepts connections but never sends messages.
func testSilentServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		// Read messages but discard them, wait for close
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
}

// httpToWS converts an http:// test server URL to a ws:// URL.
func httpToWS(s *httptest.Server) string {
	return "ws" + strings.TrimPrefix(s.URL, "http")
}

// --- Publish tests ---

func TestPublish_BasicEcho(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "hello", publishOptions{
		Count: 1,
		Wait:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	if !strings.Contains(out, "> hello") {
		t.Errorf("expected sent message in output, got: %s", out)
	}
	if !strings.Contains(out, "< hello") {
		t.Errorf("expected echoed message in output, got: %s", out)
	}
	if !strings.Contains(stderr.String(), "Connected to") {
		t.Errorf("expected Connected message on stderr, got: %s", stderr.String())
	}
	if !strings.Contains(stderr.String(), "Disconnected") {
		t.Errorf("expected Disconnected message on stderr, got: %s", stderr.String())
	}
}

func TestPublish_CountMultiple(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "tick", publishOptions{
		Count:    3,
		Wait:     2 * time.Second,
		Interval: 10 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	sentCount := strings.Count(out, "> tick")
	if sentCount != 3 {
		t.Errorf("expected 3 sent messages, got %d: %s", sentCount, out)
	}
}

func TestPublish_NoResponse(t *testing.T) {
	srv := testSilentServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "fire-and-forget", publishOptions{
		Count:      1,
		NoResponse: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	if !strings.Contains(out, "> fire-and-forget") {
		t.Errorf("expected sent message, got: %s", out)
	}
	if strings.Contains(out, "<") {
		t.Errorf("expected no response lines, got: %s", out)
	}
}

func TestPublish_WaitTimeout(t *testing.T) {
	srv := testSilentServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "hello", publishOptions{
		Count: 1,
		Wait:  200 * time.Millisecond,
	})
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should have waited approximately the --wait duration
	if elapsed < 150*time.Millisecond {
		t.Errorf("expected to wait ~200ms, but finished in %v", elapsed)
	}
}

func TestPublish_InvalidURL(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx := context.Background()

	err := runPublish(ctx, &stdout, &stderr, "http://localhost:1234", "hello", publishOptions{Count: 1})
	if err == nil {
		t.Fatal("expected error for http:// URL")
	}
	if !strings.Contains(err.Error(), "ws:// or wss://") {
		t.Errorf("expected scheme error, got: %v", err)
	}
}

func TestPublish_ConnectionRefused(t *testing.T) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, "ws://127.0.0.1:1/ws", "hello", publishOptions{Count: 1})
	if err == nil {
		t.Fatal("expected connection error")
	}
	if !strings.Contains(err.Error(), "failed to connect") {
		t.Errorf("expected connection failure, got: %v", err)
	}
}

func TestPublish_Cancellation(t *testing.T) {
	srv := testSilentServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "hello", publishOptions{
		Count: 1,
		Wait:  10 * time.Second, // Long wait, should be cancelled
	})
	// Either nil or context.Canceled is acceptable
	if err != nil && err != context.Canceled {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Template tests ---

func TestPublish_TemplateWithIndex(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), `msg-{{.Index}}`, publishOptions{
		Count:    3,
		Wait:     2 * time.Second,
		Interval: 10 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	for i := 0; i < 3; i++ {
		expected := fmt.Sprintf("> msg-%d", i)
		if !strings.Contains(out, expected) {
			t.Errorf("expected %q in output, got: %s", expected, out)
		}
	}
}

func TestPublish_TemplateWithUUID(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), `{"id":"{{uuid}}"}`, publishOptions{
		Count: 2,
		Wait:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	// Each line should have a different UUID (36 chars with dashes)
	lines := strings.Split(strings.TrimSpace(out), "\n")
	var uuids []string
	for _, line := range lines {
		if strings.HasPrefix(line, "> ") {
			// Extract UUID from {"id":"<uuid>"}
			start := strings.Index(line, `"id":"`) + 6
			end := strings.LastIndex(line, `"`)
			if start > 6 && end > start {
				uuids = append(uuids, line[start:end])
			}
		}
	}
	if len(uuids) < 2 {
		t.Fatalf("expected at least 2 UUIDs, got %d from: %s", len(uuids), out)
	}
	if uuids[0] == uuids[1] {
		t.Errorf("expected different UUIDs, got same: %s", uuids[0])
	}
}

func TestPublish_TemplateWithFaker(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), `{{name}} <{{email}}>`, publishOptions{
		Count: 1,
		Wait:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	// Should contain an @ from the email
	if !strings.Contains(out, "@") {
		t.Errorf("expected email with @ in output, got: %s", out)
	}
	// Should contain a < and > from formatting
	if !strings.Contains(out, "<") || !strings.Contains(out, ">") {
		t.Errorf("expected angle brackets in output, got: %s", out)
	}
}

func TestPublish_TemplateJSON(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tmpl := `{"index":{{.Index}},"total":{{.Count}},"rand":{{intRange 1 1000}}}`
	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), tmpl, publishOptions{
		Count: 1,
		Wait:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Extract the sent message (after "> ")
	out := stdout.String()
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "> ") {
			payload := strings.TrimPrefix(line, "> ")
			var result map[string]interface{}
			if err := json.Unmarshal([]byte(payload), &result); err != nil {
				t.Fatalf("sent message is not valid JSON: %v\npayload: %s", err, payload)
			}
			if result["index"] != float64(0) {
				t.Errorf("expected index=0, got %v", result["index"])
			}
			if result["total"] != float64(1) {
				t.Errorf("expected total=1, got %v", result["total"])
			}
			if result["rand"] == nil {
				t.Error("expected rand field")
			}
			break
		}
	}
}

func TestPublish_TemplateFromFile(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	// Write template to temp file
	dir := t.TempDir()
	tmplFile := filepath.Join(dir, "msg.tmpl")
	if err := os.WriteFile(tmplFile, []byte(`file-msg-{{.Index}}`), 0644); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "", publishOptions{
		Count: 2,
		Wait:  2 * time.Second,
		File:  tmplFile,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := stdout.String()
	if !strings.Contains(out, "> file-msg-0") {
		t.Errorf("expected file-msg-0 in output, got: %s", out)
	}
	if !strings.Contains(out, "> file-msg-1") {
		t.Errorf("expected file-msg-1 in output, got: %s", out)
	}
}

func TestPublish_InvalidTemplate(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), `{{invalid`, publishOptions{Count: 1})
	if err == nil {
		t.Fatal("expected error for invalid template")
	}
	if !strings.Contains(err.Error(), "invalid template") {
		t.Errorf("expected template parse error, got: %v", err)
	}
}

func TestPublish_PlainTextStillWorks(t *testing.T) {
	srv := testEchoServer(t)
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Plain text without any template syntax should pass through unchanged
	err := runPublish(ctx, &stdout, &stderr, httpToWS(srv), "plain hello world", publishOptions{
		Count: 1,
		Wait:  2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(stdout.String(), "> plain hello world") {
		t.Errorf("expected plain text in output, got: %s", stdout.String())
	}
}

// --- validateWSURL tests ---

func TestValidateWSURL(t *testing.T) {
	tests := []struct {
		url     string
		wantErr bool
	}{
		{"ws://localhost:8080/ws", false},
		{"wss://example.com/ws", false},
		{"http://localhost:8080", true},
		{"https://example.com", true},
		{"ftp://example.com", true},
		{"not-a-url", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("url=%s", tt.url), func(t *testing.T) {
			err := validateWSURL(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateWSURL(%q) error = %v, wantErr %v", tt.url, err, tt.wantErr)
			}
		})
	}
}
