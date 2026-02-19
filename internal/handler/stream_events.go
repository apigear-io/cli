package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/apigear-io/cli/pkg/stream"
	"github.com/go-chi/chi/v5"
)

// StreamDashboardEvents godoc
// @Summary Stream dashboard events
// @Description Server-Sent Events stream for real-time dashboard statistics
// @Tags stream
// @Produce text/event-stream
// @Success 200 {object} stream.DashboardStats
// @Router /api/v1/stream/events/dashboard [get]
func StreamDashboardEvents(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial data
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"message": "Dashboard events stream connected",
		})

		// Send updates every 2 seconds
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				stats := services.GetDashboardStats()
				sendSSEEvent(w, flusher, "stats", stats)
			}
		}
	}
}

// StreamProxyEvents godoc
// @Summary Stream proxy events
// @Description Server-Sent Events stream for real-time proxy statistics and messages
// @Tags stream
// @Produce text/event-stream
// @Param name path string true "Proxy name"
// @Success 200 {object} stream.ParsedMessageEvent
// @Router /api/v1/stream/proxies/{name}/events [get]
func StreamProxyEvents(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("proxy name is required"),
				"proxy name parameter must not be empty")
			return
		}

		// Check if proxy exists
		_, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusNotFound,
				fmt.Errorf("proxy not found: %s", name),
				"Proxy not found")
			return
		}

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial connection message
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"proxy":   name,
			"message": fmt.Sprintf("Proxy events stream connected for %s", name),
		})

		// Subscribe to proxy messages (if message hub is available)
		if services.MessageHub != nil {
			msgCh := services.MessageHub.Subscribe(name)
			defer services.MessageHub.Unsubscribe(msgCh)

			// Send stats updates periodically
			statsTicker := time.NewTicker(1 * time.Second)
			defer statsTicker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-msgCh:
					// Convert to parsed message event
					event := stream.ConvertToParsedMessageEvent(name, msg.Direction, msg.Data)
					sendSSEEvent(w, flusher, "message", event)
				case <-statsTicker.C:
					// Send stats update (placeholder)
					sendSSEEvent(w, flusher, "stats", map[string]interface{}{
						"proxy":   name,
						"message": "Stats update (not yet implemented)",
					})
				}
			}
		} else {
			// No message hub, just send stats
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					// Send stats update (placeholder)
					sendSSEEvent(w, flusher, "stats", map[string]interface{}{
						"proxy":   name,
						"message": "Stats update (not yet implemented)",
					})
				}
			}
		}
	}
}

// StreamClientEvents godoc
// @Summary Stream client events
// @Description Server-Sent Events stream for real-time client status updates
// @Tags stream
// @Produce text/event-stream
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/stream/events/clients [get]
func StreamClientEvents(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial connection message
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"message": "Client events stream connected",
		})

		// Send updates every 2 seconds
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Get all client statuses
				clients := services.ClientManager.ListClients()
				sendSSEEvent(w, flusher, "clients", clients)
			}
		}
	}
}

// StreamScriptOutput godoc
// @Summary Stream script output
// @Description Server-Sent Events stream for real-time script console output
// @Tags stream
// @Produce text/event-stream
// @Param id query string true "Script ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/stream/scripts/output [get]
func StreamScriptOutput(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scriptID := r.URL.Query().Get("id")
		if scriptID == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("script id is required"),
				"script id parameter must not be empty")
			return
		}

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Get engine
		engine := services.ScriptManager.GetEngine(scriptID)
		if engine == nil {
			writeError(w, http.StatusNotFound,
				fmt.Errorf("script not found: %s", scriptID),
				"Script not found")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial connection message
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"scriptId": scriptID,
			"message":  "Script output stream connected",
		})

		// Read from engine output channel
		outputCh := engine.Output()
		for {
			select {
			case <-ctx.Done():
				return
			case entry, ok := <-outputCh:
				if !ok {
					// Channel closed, script stopped
					sendSSEEvent(w, flusher, "closed", map[string]interface{}{
						"message": "Script stopped",
					})
					return
				}
				sendSSEEvent(w, flusher, "output", entry)
			}
		}
	}
}

// StreamTracePlayback godoc
// @Summary Stream trace playback
// @Description Server-Sent Events stream for playing back a trace file
// @Tags stream
// @Produce text/event-stream
// @Param filename query string true "Trace filename"
// @Param speed query number false "Playback speed multiplier (default: 1.0)"
// @Param loop query boolean false "Loop playback"
// @Success 200 {object} stream.ParsedMessageEvent
// @Router /api/v1/stream/traces/play [get]
func StreamTracePlayback(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")
		if filename == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("filename is required"),
				"filename parameter must not be empty")
			return
		}

		// Parse speed parameter
		speed := 1.0
		if speedStr := r.URL.Query().Get("speed"); speedStr != "" {
			if _, err := fmt.Sscanf(speedStr, "%f", &speed); err != nil {
				speed = 1.0
			}
		}

		// Parse loop parameter
		loop := r.URL.Query().Get("loop") == "true"

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial connection message
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"filename": filename,
			"speed":    speed,
			"loop":     loop,
			"message":  "Trace playback stream connected",
		})

		// TODO: Implement trace playback using tracing.Player
		// For now, send a placeholder message
		sendSSEEvent(w, flusher, "info", map[string]interface{}{
			"message": "Trace playback not yet implemented - requires tracing.Player integration",
		})

		// Keep connection alive
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Send keepalive
				sendSSEEvent(w, flusher, "keepalive", map[string]interface{}{
					"timestamp": time.Now().UnixMilli(),
				})
			}
		}
	}
}

// sendSSEEvent is a helper to send a Server-Sent Event.
func sendSSEEvent(w http.ResponseWriter, flusher http.Flusher, eventType string, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Log error but don't break the stream
		return
	}

	// Send event type if provided
	if eventType != "" {
		_, _ = fmt.Fprintf(w, "event: %s\n", eventType)
	}

	// Send data
	_, _ = fmt.Fprintf(w, "data: %s\n\n", jsonData)
	flusher.Flush()
}
