package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/itchyny/gojq"

	"github.com/apigear-io/cli/pkg/stream/tracing"
)

// EditorMessage represents a parsed ObjectLink message with metadata
type EditorMessage struct {
	Index     int                    `json:"index"`
	Timestamp int64                  `json:"timestamp"` // Unix ms
	Direction string                 `json:"direction"` // SEND or RECV
	Proxy     string                 `json:"proxy"`
	Raw       map[string]interface{} `json:"raw"`    // Raw message object
	Parsed    ParsedObjectLink       `json:"parsed"` // Parsed ObjectLink message
}

// ParsedObjectLink represents a parsed ObjectLink message
type ParsedObjectLink struct {
	MsgType     int         `json:"msgType"`
	MsgTypeName string      `json:"msgTypeName"`
	Symbol      string      `json:"symbol,omitempty"`
	ObjectID    string      `json:"objectId,omitempty"`
	RequestID   int         `json:"requestId,omitempty"`
	Args        interface{} `json:"args,omitempty"`
}

// TimeRange represents a time range with start and end timestamps
type TimeRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// EditorSession holds a loaded trace file session
type EditorSession struct {
	ID         string
	Filename   string
	Messages   []EditorMessage
	Proxies    []string
	Interfaces []string
	TimeRange  TimeRange
	CreatedAt  time.Time
	mu         sync.RWMutex
}

// EditorManager manages editor sessions
type EditorManager struct {
	sessions map[string]*EditorSession
	mu       sync.RWMutex
	stopChan chan struct{}
}

// NewEditorManager creates a new editor manager
func NewEditorManager() *EditorManager {
	em := &EditorManager{
		sessions: make(map[string]*EditorSession),
		stopChan: make(chan struct{}),
	}
	go em.cleanupLoop()
	return em
}

// cleanupLoop removes stale sessions (older than 30 minutes)
func (em *EditorManager) cleanupLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			em.cleanup()
		case <-em.stopChan:
			return
		}
	}
}

// cleanup removes sessions older than 30 minutes
func (em *EditorManager) cleanup() {
	em.mu.Lock()
	defer em.mu.Unlock()

	cutoff := time.Now().Add(-30 * time.Minute)
	for id, session := range em.sessions {
		if session.CreatedAt.Before(cutoff) {
			delete(em.sessions, id)
		}
	}
}

// Stop stops the editor manager
func (em *EditorManager) Stop() {
	close(em.stopChan)
}

// CreateSession creates a new editor session
func (em *EditorManager) CreateSession(filename string, messages []EditorMessage) *EditorSession {
	em.mu.Lock()
	defer em.mu.Unlock()

	session := &EditorSession{
		ID:        uuid.New().String(),
		Filename:  filename,
		Messages:  messages,
		CreatedAt: time.Now(),
	}

	// Calculate metadata
	proxies := make(map[string]bool)
	interfaces := make(map[string]bool)

	for _, msg := range messages {
		proxies[msg.Proxy] = true
		if msg.Parsed.Symbol != "" {
			interfaces[msg.Parsed.Symbol] = true
		}
	}

	session.Proxies = make([]string, 0, len(proxies))
	for p := range proxies {
		session.Proxies = append(session.Proxies, p)
	}

	session.Interfaces = make([]string, 0, len(interfaces))
	for i := range interfaces {
		session.Interfaces = append(session.Interfaces, i)
	}

	if len(messages) > 0 {
		session.TimeRange.Start = messages[0].Timestamp
		session.TimeRange.End = messages[len(messages)-1].Timestamp
	}

	em.sessions[session.ID] = session
	return session
}

// GetSession retrieves a session by ID
func (em *EditorManager) GetSession(id string) *EditorSession {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return em.sessions[id]
}

// parseTraceEntry converts a trace entry to EditorMessage
func parseTraceEntry(entry tracing.TraceEntry, index int) EditorMessage {
	msg := EditorMessage{
		Index:     index,
		Timestamp: entry.Timestamp,
		Direction: entry.Direction,
		Proxy:     entry.Proxy,
		Raw:       make(map[string]interface{}),
	}

	// Unmarshal message to interface{}
	var msgData interface{}
	if err := json.Unmarshal(entry.Message, &msgData); err == nil {
		msg.Raw["msg"] = msgData
		msg.Parsed = parseObjectLinkMessage(msgData)
	}

	return msg
}

// parseObjectLinkMessage extracts ObjectLink-specific fields
func parseObjectLinkMessage(data interface{}) ParsedObjectLink {
	parsed := ParsedObjectLink{
		MsgTypeName: "UNKNOWN",
	}

	msgMap, ok := data.(map[string]interface{})
	if !ok {
		return parsed
	}

	// Extract msgType
	if msgType, ok := msgMap["msgType"].(float64); ok {
		parsed.MsgType = int(msgType)
		parsed.MsgTypeName = getMessageTypeName(int(msgType))
	}

	// Extract symbol (interface name)
	if symbol, ok := msgMap["symbol"].(string); ok {
		parsed.Symbol = symbol
	}

	// Extract objectId
	if objectID, ok := msgMap["objectId"].(string); ok {
		parsed.ObjectID = objectID
	}

	// Extract requestId
	if requestID, ok := msgMap["requestId"].(float64); ok {
		parsed.RequestID = int(requestID)
	}

	// Extract args
	if args, ok := msgMap["args"]; ok {
		parsed.Args = args
	}

	return parsed
}

// getMessageTypeName returns the message type name
func getMessageTypeName(msgType int) string {
	names := map[int]string{
		10: "LINK",
		11: "INIT",
		12: "UNLINK",
		20: "SET_PROPERTY",
		21: "PROPERTY_CHANGE",
		30: "INVOKE",
		31: "INVOKE_REPLY",
		40: "SIGNAL",
		90: "ERROR",
	}
	if name, ok := names[msgType]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_%d", msgType)
}

// loadTraceFile loads a trace file and converts it to EditorMessages
func loadTraceFile(filename string) ([]EditorMessage, error) {
	entries, err := tracing.ReadTraceFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read trace file: %w", err)
	}

	messages := make([]EditorMessage, len(entries))
	for i, entry := range entries {
		messages[i] = parseTraceEntry(entry, i)
	}

	return messages, nil
}

// parseUploadedFile parses an uploaded JSONL file
func parseUploadedFile(file io.Reader) ([]EditorMessage, error) {
	scanner := bufio.NewScanner(file)
	messages := make([]EditorMessage, 0, 1000)
	index := 0

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(line, &data); err != nil {
			continue
		}

		msg := EditorMessage{
			Index: index,
			Raw:   make(map[string]interface{}),
		}

		if ts, ok := data["ts"].(float64); ok {
			msg.Timestamp = int64(ts)
		}

		if dir, ok := data["dir"].(string); ok {
			msg.Direction = strings.ToUpper(dir)
		}

		if proxy, ok := data["proxy"].(string); ok {
			msg.Proxy = proxy
		}

		if rawMsg, ok := data["msg"]; ok {
			msg.Raw["msg"] = rawMsg
			msg.Parsed = parseObjectLinkMessage(rawMsg)
		}

		messages = append(messages, msg)
		index++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

// EditorStatsResponse is the response for editor session creation
type EditorStatsResponse struct {
	SessionID  string    `json:"sessionId"`
	Filename   string    `json:"filename"`
	TotalCount int       `json:"totalCount"`
	TimeRange  TimeRange `json:"timeRange"`
	Proxies    []string  `json:"proxies"`
	Interfaces []string  `json:"interfaces"`
}

// EditorMessagesResponse is the response for paginated messages
type EditorMessagesResponse struct {
	Messages []EditorMessage `json:"messages"`
	Total    int             `json:"total"`
	Offset   int             `json:"offset"`
	Limit    int             `json:"limit"`
}

// EditorTimelineResponse is the response for timeline data
type EditorTimelineResponse struct {
	Buckets   []EditorBucket `json:"buckets"`
	TimeRange TimeRange      `json:"timeRange"`
}

// EditorBucket represents a time bucket in the timeline
type EditorBucket struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	SendCount int   `json:"sendCount"`
	RecvCount int   `json:"recvCount"`
}

// EditorSeekResponse is the response for seek operation
type EditorSeekResponse struct {
	Offset       int `json:"offset"`
	MessageIndex int `json:"messageIndex"`
}

// EditorJQResponse is the response for JQ query
type EditorJQResponse struct {
	Matches      []EditorJQMatch `json:"matches"`
	TotalMatches int             `json:"totalMatches"`
}

// EditorJQMatch represents a JQ query match
type EditorJQMatch struct {
	Index  int         `json:"index"`
	Result interface{} `json:"result"`
}

// Global editor manager
var editorManager *EditorManager

func init() {
	editorManager = NewEditorManager()
}

// LoadStreamEditor loads a trace file for editing
// @Summary Load trace file for editing
// @Description Load a trace file (upload or from server) and create an editor session
// @Tags stream
// @Accept json,multipart/form-data
// @Produce json
// @Success 200 {object} EditorStatsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/editor/load [post]
func LoadStreamEditor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		var filename string
		var messages []EditorMessage
		var err error

		if strings.HasPrefix(contentType, "multipart/form-data") {
			// Handle file upload
			if err := r.ParseMultipartForm(100 << 20); err != nil { // 100 MB max
				writeError(w, http.StatusBadRequest, err, "failed to parse form")
				return
			}

			file, header, err := r.FormFile("file")
			if err != nil {
				writeError(w, http.StatusBadRequest, err, "missing file")
				return
			}
			defer file.Close()

			filename = header.Filename
			messages, err = parseUploadedFile(file)
			if err != nil {
				writeError(w, http.StatusBadRequest, err, "failed to parse file")
				return
			}
		} else {
			// Handle JSON request (load from server)
			var req struct {
				Filename string `json:"filename"`
			}

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, err, "invalid request body")
				return
			}

			if req.Filename == "" {
				writeError(w, http.StatusBadRequest, nil, "missing filename")
				return
			}

			filename = req.Filename
			messages, err = loadTraceFile(filename)
			if err != nil {
				writeError(w, http.StatusInternalServerError, err, "failed to load file")
				return
			}
		}

		// Create session
		session := editorManager.CreateSession(filename, messages)

		// Return stats
		resp := EditorStatsResponse{
			SessionID:  session.ID,
			Filename:   session.Filename,
			TotalCount: len(session.Messages),
			TimeRange:  session.TimeRange,
			Proxies:    session.Proxies,
			Interfaces: session.Interfaces,
		}

		writeJSON(w, http.StatusOK, resp)
	}
}

// GetStreamEditorMessages gets paginated messages from an editor session
// @Summary Get paginated messages
// @Description Get messages from an editor session with optional filters
// @Tags stream
// @Produce json
// @Param sessionId query string true "Session ID"
// @Param offset query int false "Starting offset (default 0)"
// @Param limit query int false "Number of messages (default 100)"
// @Param proxy query string false "Filter by proxy name"
// @Param interface query string false "Filter by interface name"
// @Param direction query string false "Filter by direction (SEND/RECV)"
// @Param type query string false "Filter by message type"
// @Success 200 {object} EditorMessagesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/editor/messages [get]
func GetStreamEditorMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.URL.Query().Get("sessionId")
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		if limit == 0 {
			limit = 100
		}

		// Get filters
		proxyFilter := r.URL.Query().Get("proxy")
		interfaceFilter := r.URL.Query().Get("interface")
		directionFilter := r.URL.Query().Get("direction")
		typeFilter := r.URL.Query().Get("type")

		// Get session
		session := editorManager.GetSession(sessionID)
		if session == nil {
			writeError(w, http.StatusNotFound, nil, "session not found")
			return
		}

		session.mu.RLock()
		defer session.mu.RUnlock()

		// Filter messages
		filtered := filterEditorMessages(session.Messages, proxyFilter, interfaceFilter, directionFilter, typeFilter)

		// Paginate
		start := offset
		end := offset + limit
		if start > len(filtered) {
			start = len(filtered)
		}
		if end > len(filtered) {
			end = len(filtered)
		}

		resp := EditorMessagesResponse{
			Messages: filtered[start:end],
			Total:    len(filtered),
			Offset:   offset,
			Limit:    limit,
		}

		writeJSON(w, http.StatusOK, resp)
	}
}

// filterEditorMessages applies filters to messages
func filterEditorMessages(messages []EditorMessage, proxy, iface, direction, msgType string) []EditorMessage {
	if proxy == "" && iface == "" && direction == "" && msgType == "" {
		return messages
	}

	filtered := make([]EditorMessage, 0, len(messages))
	for _, msg := range messages {
		if proxy != "" && msg.Proxy != proxy {
			continue
		}
		if iface != "" && msg.Parsed.Symbol != iface {
			continue
		}
		if direction != "" && msg.Direction != direction {
			continue
		}
		if msgType != "" && msg.Parsed.MsgTypeName != msgType {
			continue
		}
		filtered = append(filtered, msg)
	}
	return filtered
}

// GetStreamEditorTimeline gets timeline buckets for visualization
// @Summary Get timeline buckets
// @Description Get 200 time buckets for timeline visualization
// @Tags stream
// @Produce json
// @Param sessionId query string true "Session ID"
// @Success 200 {object} EditorTimelineResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/editor/timeline [get]
func GetStreamEditorTimeline() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.URL.Query().Get("sessionId")

		session := editorManager.GetSession(sessionID)
		if session == nil {
			writeError(w, http.StatusNotFound, nil, "session not found")
			return
		}

		session.mu.RLock()
		defer session.mu.RUnlock()

		// Create 200 buckets
		const numBuckets = 200
		buckets := make([]EditorBucket, numBuckets)

		if len(session.Messages) == 0 {
			resp := EditorTimelineResponse{
				Buckets:   buckets,
				TimeRange: TimeRange{Start: 0, End: 0},
			}
			writeJSON(w, http.StatusOK, resp)
			return
		}

		timeSpan := session.TimeRange.End - session.TimeRange.Start
		if timeSpan == 0 {
			timeSpan = 1
		}
		bucketSize := float64(timeSpan) / float64(numBuckets)

		// Initialize buckets
		for i := 0; i < numBuckets; i++ {
			buckets[i].StartTime = session.TimeRange.Start + int64(float64(i)*bucketSize)
			buckets[i].EndTime = session.TimeRange.Start + int64(float64(i+1)*bucketSize)
		}

		// Count messages per bucket
		for _, msg := range session.Messages {
			bucketIdx := int(float64(msg.Timestamp-session.TimeRange.Start) / bucketSize)
			if bucketIdx < 0 {
				bucketIdx = 0
			}
			if bucketIdx >= numBuckets {
				bucketIdx = numBuckets - 1
			}

			if msg.Direction == "SEND" {
				buckets[bucketIdx].SendCount++
			} else if msg.Direction == "RECV" {
				buckets[bucketIdx].RecvCount++
			}
		}

		resp := EditorTimelineResponse{
			Buckets:   buckets,
			TimeRange: session.TimeRange,
		}

		writeJSON(w, http.StatusOK, resp)
	}
}

// SeekStreamEditor finds the message offset at a specific timestamp
// @Summary Seek to timestamp
// @Description Find the message offset at a specific timestamp
// @Tags stream
// @Produce json
// @Param sessionId query string true "Session ID"
// @Param timestamp query int64 true "Unix timestamp in milliseconds"
// @Param proxy query string false "Filter by proxy name"
// @Param interface query string false "Filter by interface name"
// @Param direction query string false "Filter by direction (SEND/RECV)"
// @Param type query string false "Filter by message type"
// @Success 200 {object} EditorSeekResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/editor/seek [get]
func SeekStreamEditor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.URL.Query().Get("sessionId")
		timestamp, _ := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)

		// Get filters
		proxyFilter := r.URL.Query().Get("proxy")
		interfaceFilter := r.URL.Query().Get("interface")
		directionFilter := r.URL.Query().Get("direction")
		typeFilter := r.URL.Query().Get("type")

		session := editorManager.GetSession(sessionID)
		if session == nil {
			writeError(w, http.StatusNotFound, nil, "session not found")
			return
		}

		session.mu.RLock()
		defer session.mu.RUnlock()

		// Filter messages
		filtered := filterEditorMessages(session.Messages, proxyFilter, interfaceFilter, directionFilter, typeFilter)

		// Find first message at or after timestamp
		idx := 0
		for i, msg := range filtered {
			if msg.Timestamp >= timestamp {
				idx = i
				break
			}
		}

		resp := EditorSeekResponse{
			Offset:       idx,
			MessageIndex: idx,
		}

		writeJSON(w, http.StatusOK, resp)
	}
}

// ExportStreamEditor exports messages as JSONL file
// @Summary Export messages
// @Description Export messages as JSONL file (all or specific indices)
// @Tags stream
// @Accept json
// @Produce application/x-ndjson
// @Success 200 {file} file
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/editor/export [post]
func ExportStreamEditor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			SessionID string `json:"sessionId"`
			Indices   []int  `json:"indices,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		session := editorManager.GetSession(req.SessionID)
		if session == nil {
			writeError(w, http.StatusNotFound, nil, "session not found")
			return
		}

		session.mu.RLock()
		defer session.mu.RUnlock()

		// Determine which messages to export
		var toExport []EditorMessage
		if len(req.Indices) > 0 {
			// Create index lookup
			indexSet := make(map[int]bool)
			for _, idx := range req.Indices {
				indexSet[idx] = true
			}

			// Export only selected indices
			for _, msg := range session.Messages {
				if indexSet[msg.Index] {
					toExport = append(toExport, msg)
				}
			}
		} else {
			// Export all
			toExport = session.Messages
		}

		// Write JSONL
		w.Header().Set("Content-Type", "application/x-ndjson")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", session.Filename))

		encoder := json.NewEncoder(w)
		for _, msg := range toExport {
			// Write original format
			line := map[string]interface{}{
				"ts":    msg.Timestamp,
				"dir":   msg.Direction,
				"proxy": msg.Proxy,
				"msg":   msg.Raw["msg"],
			}
			encoder.Encode(line)
		}
	}
}

// RunStreamEditorJQ runs a JQ query on messages
// @Summary Run JQ query
// @Description Run a JQ query on messages and return matching results
// @Tags stream
// @Accept json
// @Produce json
// @Success 200 {object} EditorJQResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/editor/jq [post]
func RunStreamEditorJQ() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			SessionID string `json:"sessionId"`
			Query     string `json:"query"`
			Limit     int    `json:"limit"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if req.Limit == 0 {
			req.Limit = 100
		}

		session := editorManager.GetSession(req.SessionID)
		if session == nil {
			writeError(w, http.StatusNotFound, nil, "session not found")
			return
		}

		session.mu.RLock()
		defer session.mu.RUnlock()

		// Parse JQ query
		query, err := gojq.Parse(req.Query)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid JQ query")
			return
		}

		// Compile query
		code, err := gojq.Compile(query)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to compile JQ query")
			return
		}

		// Run query on each message
		matches := make([]EditorJQMatch, 0, req.Limit)
		ctx := context.Background()

		for _, msg := range session.Messages {
			if len(matches) >= req.Limit {
				break
			}

			// Convert message to map for JQ
			msgData := map[string]interface{}{
				"index":     msg.Index,
				"timestamp": msg.Timestamp,
				"direction": msg.Direction,
				"proxy":     msg.Proxy,
				"msg":       msg.Raw["msg"],
			}

			// Run query
			iter := code.RunWithContext(ctx, msgData)
			for {
				v, ok := iter.Next()
				if !ok {
					break
				}
				if err, ok := v.(error); ok {
					// Query returned error for this message, skip
					_ = err
					break
				}
				// Query matched (returned non-false/null value)
				if v != nil && v != false {
					matches = append(matches, EditorJQMatch{
						Index:  msg.Index,
						Result: v,
					})
					break
				}
			}
		}

		resp := EditorJQResponse{
			Matches:      matches,
			TotalMatches: len(matches),
		}

		writeJSON(w, http.StatusOK, resp)
	}
}
