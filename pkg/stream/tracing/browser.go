package tracing

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Browser provides high-level operations for browsing trace files.
type Browser struct {
	traceDir string
}

// NewBrowser creates a new trace browser.
func NewBrowser() *Browser {
	return &Browser{
		traceDir: GetTraceDir(),
	}
}

// SearchOptions specifies criteria for searching traces.
type SearchOptions struct {
	ProxyName  string    // Filter by proxy name
	Direction  string    // Filter by direction (SEND/RECV)
	StartTime  time.Time // Filter by start time
	EndTime    time.Time // Filter by end time
	MaxFiles   int       // Maximum number of files to search (0 = all)
	MaxEntries int       // Maximum number of entries to return (0 = all)
}

// SearchResult contains a trace entry with its source file information.
type SearchResult struct {
	Entry    TraceEntry    `json:"entry"`
	File     TraceFileInfo `json:"file"`
	Position int           `json:"position"` // Position in file
}

// Search searches across all trace files for entries matching the criteria.
func (b *Browser) Search(options SearchOptions) ([]SearchResult, error) {
	// List all trace files
	files, err := ListTraceFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to list trace files: %w", err)
	}

	// Filter files by proxy name if specified
	if options.ProxyName != "" {
		filtered := make([]TraceFileInfo, 0)
		for _, f := range files {
			if f.ProxyName == options.ProxyName {
				filtered = append(filtered, f)
			}
		}
		files = filtered
	}

	// Limit number of files
	if options.MaxFiles > 0 && len(files) > options.MaxFiles {
		files = files[:options.MaxFiles]
	}

	// Search each file
	var results []SearchResult
	for _, file := range files {
		// Create filter for reader
		filter := FilterOptions{
			Direction: options.Direction,
			StartTime: options.StartTime.UnixMilli(),
			EndTime:   options.EndTime.UnixMilli(),
		}

		// Read entries
		entries, err := ReadTraceFileFiltered(file.Name, filter)
		if err != nil {
			// Skip files that can't be read
			continue
		}

		// Add to results
		for i, entry := range entries {
			results = append(results, SearchResult{
				Entry:    entry,
				File:     file,
				Position: i,
			})

			// Check max entries limit
			if options.MaxEntries > 0 && len(results) >= options.MaxEntries {
				return results, nil
			}
		}
	}

	// Sort by timestamp (newest first)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Entry.Timestamp > results[j].Entry.Timestamp
	})

	return results, nil
}

// GetFileStats returns statistics for a specific trace file.
func (b *Browser) GetFileStats(filename string) (*FileStats, error) {
	entries, err := ReadTraceFile(filename)
	if err != nil {
		return nil, err
	}

	stats := &FileStats{
		Filename:     filename,
		EntryCount:   len(entries),
		DirectionCounts: make(map[string]int),
		ProxyCounts:    make(map[string]int),
	}

	if len(entries) == 0 {
		return stats, nil
	}

	stats.FirstTimestamp = entries[0].Timestamp
	stats.LastTimestamp = entries[len(entries)-1].Timestamp
	stats.Duration = time.Duration(stats.LastTimestamp-stats.FirstTimestamp) * time.Millisecond

	// Count by direction and proxy
	for _, entry := range entries {
		stats.DirectionCounts[entry.Direction]++
		stats.ProxyCounts[entry.Proxy]++
	}

	return stats, nil
}

// FileStats contains statistics about a trace file.
type FileStats struct {
	Filename        string            `json:"filename"`
	EntryCount      int               `json:"entryCount"`
	FirstTimestamp  int64             `json:"firstTimestamp"`
	LastTimestamp   int64             `json:"lastTimestamp"`
	Duration        time.Duration     `json:"duration"`
	DirectionCounts map[string]int    `json:"directionCounts"` // "SEND" -> count, "RECV" -> count
	ProxyCounts     map[string]int    `json:"proxyCounts"`     // proxy name -> count
}

// GetRecentEntries returns the most recent N entries across all trace files.
func (b *Browser) GetRecentEntries(count int) ([]SearchResult, error) {
	return b.Search(SearchOptions{
		MaxEntries: count,
	})
}

// GetEntriesForProxy returns entries for a specific proxy.
func (b *Browser) GetEntriesForProxy(proxyName string, limit int) ([]SearchResult, error) {
	return b.Search(SearchOptions{
		ProxyName:  proxyName,
		MaxEntries: limit,
	})
}

// GetEntriesInTimeRange returns entries within a time range.
func (b *Browser) GetEntriesInTimeRange(start, end time.Time, limit int) ([]SearchResult, error) {
	return b.Search(SearchOptions{
		StartTime:  start,
		EndTime:    end,
		MaxEntries: limit,
	})
}

// FindMessagesByContent searches for messages containing specific content.
// This performs a simple string search in the raw message JSON.
func (b *Browser) FindMessagesByContent(searchText string, limit int) ([]SearchResult, error) {
	files, err := ListTraceFiles()
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	searchLower := strings.ToLower(searchText)

	for _, file := range files {
		entries, err := ReadTraceFile(file.Name)
		if err != nil {
			continue
		}

		for i, entry := range entries {
			// Search in message content
			msgStr := strings.ToLower(string(entry.Message))
			if strings.Contains(msgStr, searchLower) {
				results = append(results, SearchResult{
					Entry:    entry,
					File:     file,
					Position: i,
				})

				if limit > 0 && len(results) >= limit {
					return results, nil
				}
			}
		}
	}

	return results, nil
}

// GetMessageDetails parses a message and returns its details.
func GetMessageDetails(entry TraceEntry) (*MessageDetails, error) {
	var msgArray []interface{}
	if err := json.Unmarshal(entry.Message, &msgArray); err != nil {
		return nil, fmt.Errorf("invalid message format: %w", err)
	}

	if len(msgArray) == 0 {
		return nil, fmt.Errorf("empty message")
	}

	details := &MessageDetails{
		Timestamp: entry.Timestamp,
		Direction: entry.Direction,
		Proxy:     entry.Proxy,
		Raw:       entry.Message,
	}

	// Parse message type (first element)
	if msgType, ok := msgArray[0].(float64); ok {
		details.MessageType = int(msgType)
		details.MessageTypeName = getMessageTypeName(details.MessageType)
	}

	// Parse based on message type
	switch details.MessageType {
	case 10: // LINK
		if len(msgArray) >= 2 {
			details.ObjectID, _ = msgArray[1].(string)
		}
	case 11: // INIT
		if len(msgArray) >= 3 {
			details.ObjectID, _ = msgArray[1].(string)
			details.Args = msgArray[2]
		}
	case 30: // INVOKE
		if len(msgArray) >= 4 {
			if reqID, ok := msgArray[1].(float64); ok {
				details.RequestID = int64(reqID)
			}
			details.Symbol, _ = msgArray[2].(string)
			details.Args = msgArray[3]
		}
	case 31: // INVOKE_REPLY
		if len(msgArray) >= 3 {
			if reqID, ok := msgArray[1].(float64); ok {
				details.RequestID = int64(reqID)
			}
			details.Args = msgArray[2]
		}
	case 40: // SIGNAL
		if len(msgArray) >= 3 {
			details.Symbol, _ = msgArray[1].(string)
			details.Args = msgArray[2]
		}
	case 50: // PROPERTY_CHANGE
		if len(msgArray) >= 3 {
			details.Symbol, _ = msgArray[1].(string)
			details.Args = msgArray[2]
		}
	}

	return details, nil
}

// MessageDetails contains parsed information about a trace message.
type MessageDetails struct {
	Timestamp       int64           `json:"timestamp"`
	Direction       string          `json:"direction"`
	Proxy           string          `json:"proxy"`
	MessageType     int             `json:"messageType"`
	MessageTypeName string          `json:"messageTypeName"`
	ObjectID        string          `json:"objectId,omitempty"`
	Symbol          string          `json:"symbol,omitempty"`
	RequestID       int64           `json:"requestId,omitempty"`
	Args            interface{}     `json:"args,omitempty"`
	Raw             json.RawMessage `json:"raw"`
}

// getMessageTypeName returns the name of a message type.
func getMessageTypeName(msgType int) string {
	switch msgType {
	case 10:
		return "LINK"
	case 11:
		return "INIT"
	case 12:
		return "UNLINK"
	case 20:
		return "SET_PROPERTY"
	case 30:
		return "INVOKE"
	case 31:
		return "INVOKE_REPLY"
	case 40:
		return "SIGNAL"
	case 50:
		return "PROPERTY_CHANGE"
	case 70:
		return "ERROR"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", msgType)
	}
}
