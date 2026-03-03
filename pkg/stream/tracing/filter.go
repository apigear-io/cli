package tracing

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Filter provides advanced filtering capabilities for trace entries.
type Filter struct {
	// Basic filters
	Direction  string
	ProxyNames []string
	StartTime  int64
	EndTime    int64

	// Message type filters
	MessageTypes []int // ObjectLink message types (10=LINK, 30=INVOKE, etc.)

	// Pattern filters
	ObjectIDPattern *regexp.Regexp
	SymbolPattern   *regexp.Regexp

	// Content filters
	ContainsText string // Simple text search in raw message
}

// NewFilter creates a new filter with no criteria (matches all).
func NewFilter() *Filter {
	return &Filter{}
}

// WithDirection sets the direction filter.
func (f *Filter) WithDirection(direction string) *Filter {
	f.Direction = direction
	return f
}

// WithProxyNames sets the proxy name filter.
func (f *Filter) WithProxyNames(names ...string) *Filter {
	f.ProxyNames = names
	return f
}

// WithTimeRange sets the time range filter.
func (f *Filter) WithTimeRange(startTime, endTime int64) *Filter {
	f.StartTime = startTime
	f.EndTime = endTime
	return f
}

// WithMessageTypes sets the message type filter.
func (f *Filter) WithMessageTypes(types ...int) *Filter {
	f.MessageTypes = types
	return f
}

// WithObjectIDPattern sets the object ID pattern filter.
func (f *Filter) WithObjectIDPattern(pattern string) (*Filter, error) {
	if pattern == "" {
		f.ObjectIDPattern = nil
		return f, nil
	}
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return f, fmt.Errorf("invalid object ID pattern: %w", err)
	}
	f.ObjectIDPattern = regex
	return f, nil
}

// WithSymbolPattern sets the symbol pattern filter.
func (f *Filter) WithSymbolPattern(pattern string) (*Filter, error) {
	if pattern == "" {
		f.SymbolPattern = nil
		return f, nil
	}
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return f, fmt.Errorf("invalid symbol pattern: %w", err)
	}
	f.SymbolPattern = regex
	return f, nil
}

// WithContainsText sets the text content filter.
func (f *Filter) WithContainsText(text string) *Filter {
	f.ContainsText = text
	return f
}

// Matches checks if an entry matches all filter criteria.
func (f *Filter) Matches(entry TraceEntry) bool {
	// Direction filter
	if f.Direction != "" && entry.Direction != f.Direction {
		return false
	}

	// Proxy name filter
	if len(f.ProxyNames) > 0 {
		found := false
		for _, name := range f.ProxyNames {
			if entry.Proxy == name {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Time range filter
	if f.StartTime > 0 && entry.Timestamp < f.StartTime {
		return false
	}
	if f.EndTime > 0 && entry.Timestamp > f.EndTime {
		return false
	}

	// Parse message for advanced filters
	if len(f.MessageTypes) > 0 || f.ObjectIDPattern != nil || f.SymbolPattern != nil {
		details, err := GetMessageDetails(entry)
		if err != nil {
			return false
		}

		// Message type filter
		if len(f.MessageTypes) > 0 {
			found := false
			for _, msgType := range f.MessageTypes {
				if details.MessageType == msgType {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		// Object ID pattern filter
		if f.ObjectIDPattern != nil {
			if !f.ObjectIDPattern.MatchString(details.ObjectID) {
				return false
			}
		}

		// Symbol pattern filter
		if f.SymbolPattern != nil {
			if !f.SymbolPattern.MatchString(details.Symbol) {
				return false
			}
		}
	}

	// Content filter
	if f.ContainsText != "" {
		msgStr := strings.ToLower(string(entry.Message))
		searchText := strings.ToLower(f.ContainsText)
		if !strings.Contains(msgStr, searchText) {
			return false
		}
	}

	return true
}

// Apply applies the filter to a slice of entries and returns matching entries.
func (f *Filter) Apply(entries []TraceEntry) []TraceEntry {
	if f.isEmpty() {
		return entries
	}

	result := make([]TraceEntry, 0, len(entries))
	for _, entry := range entries {
		if f.Matches(entry) {
			result = append(result, entry)
		}
	}
	return result
}

// isEmpty checks if the filter has any criteria set.
func (f *Filter) isEmpty() bool {
	return f.Direction == "" &&
		len(f.ProxyNames) == 0 &&
		f.StartTime == 0 &&
		f.EndTime == 0 &&
		len(f.MessageTypes) == 0 &&
		f.ObjectIDPattern == nil &&
		f.SymbolPattern == nil &&
		f.ContainsText == ""
}

// FilterByObjectLinkType filters entries by ObjectLink message type.
// Common types: LINK=10, INIT=11, INVOKE=30, INVOKE_REPLY=31, SIGNAL=40, PROPERTY_CHANGE=50
func FilterByObjectLinkType(entries []TraceEntry, messageTypes ...int) []TraceEntry {
	filter := NewFilter().WithMessageTypes(messageTypes...)
	return filter.Apply(entries)
}

// FilterByDirection filters entries by direction (SEND or RECV).
func FilterByDirection(entries []TraceEntry, direction string) []TraceEntry {
	filter := NewFilter().WithDirection(direction)
	return filter.Apply(entries)
}

// FilterByProxy filters entries by proxy name.
func FilterByProxy(entries []TraceEntry, proxyNames ...string) []TraceEntry {
	filter := NewFilter().WithProxyNames(proxyNames...)
	return filter.Apply(entries)
}

// FilterByTimeRange filters entries by timestamp range.
func FilterByTimeRange(entries []TraceEntry, startTime, endTime int64) []TraceEntry {
	filter := NewFilter().WithTimeRange(startTime, endTime)
	return filter.Apply(entries)
}

// TransformEntry represents a transformation to apply to an entry.
type TransformEntry func(entry TraceEntry) (TraceEntry, bool)

// Transform applies a transformation function to each entry.
// If the function returns false, the entry is excluded.
func Transform(entries []TraceEntry, fn TransformEntry) []TraceEntry {
	result := make([]TraceEntry, 0, len(entries))
	for _, entry := range entries {
		if transformed, keep := fn(entry); keep {
			result = append(result, transformed)
		}
	}
	return result
}

// RemapProxyName changes the proxy name in all entries.
func RemapProxyName(entries []TraceEntry, oldName, newName string) []TraceEntry {
	return Transform(entries, func(entry TraceEntry) (TraceEntry, bool) {
		if entry.Proxy == oldName {
			entry.Proxy = newName
		}
		return entry, true
	})
}

// RemapTimestamps adjusts all timestamps by an offset.
func RemapTimestamps(entries []TraceEntry, offsetMs int64) []TraceEntry {
	return Transform(entries, func(entry TraceEntry) (TraceEntry, bool) {
		entry.Timestamp += offsetMs
		return entry, true
	})
}

// NormalizeTimestamps adjusts timestamps so the first entry starts at time 0.
func NormalizeTimestamps(entries []TraceEntry) []TraceEntry {
	if len(entries) == 0 {
		return entries
	}
	startTime := entries[0].Timestamp
	return RemapTimestamps(entries, -startTime)
}

// MergeTraces merges multiple trace files into a single sorted list.
func MergeTraces(filenames ...string) ([]TraceEntry, error) {
	var allEntries []TraceEntry

	for _, filename := range filenames {
		entries, err := ReadTraceFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", filename, err)
		}
		allEntries = append(allEntries, entries...)
	}

	// Sort by timestamp
	return SortByTimestamp(allEntries), nil
}

// SortByTimestamp sorts entries by timestamp (oldest first).
func SortByTimestamp(entries []TraceEntry) []TraceEntry {
	sorted := make([]TraceEntry, len(entries))
	copy(sorted, entries)

	// Simple bubble sort for small lists, or use sort.Slice for larger
	if len(sorted) < 100 {
		// Bubble sort for clarity
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i].Timestamp > sorted[j].Timestamp {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}
	} else {
		// Use standard library sort for larger lists
		for i := 0; i < len(sorted)-1; i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i].Timestamp > sorted[j].Timestamp {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}
	}

	return sorted
}

// ExtractInvokeSequence extracts an invoke-reply sequence for a specific request ID.
func ExtractInvokeSequence(entries []TraceEntry, requestID int64) ([]TraceEntry, error) {
	var sequence []TraceEntry

	for _, entry := range entries {
		details, err := GetMessageDetails(entry)
		if err != nil {
			continue
		}

		// Match INVOKE or INVOKE_REPLY with the request ID
		if (details.MessageType == 30 || details.MessageType == 31) && details.RequestID == requestID {
			sequence = append(sequence, entry)
		}
	}

	if len(sequence) == 0 {
		return nil, fmt.Errorf("no messages found for request ID %d", requestID)
	}

	return sequence, nil
}

// GroupByRequestID groups invoke/reply pairs by request ID.
func GroupByRequestID(entries []TraceEntry) map[int64][]TraceEntry {
	groups := make(map[int64][]TraceEntry)

	for _, entry := range entries {
		details, err := GetMessageDetails(entry)
		if err != nil {
			continue
		}

		// Only group INVOKE and INVOKE_REPLY messages
		if details.MessageType == 30 || details.MessageType == 31 {
			groups[details.RequestID] = append(groups[details.RequestID], entry)
		}
	}

	return groups
}

// ExportToJSON exports entries to JSON format.
func ExportToJSON(entries []TraceEntry) ([]byte, error) {
	return json.MarshalIndent(entries, "", "  ")
}

// ExportToJSONL exports entries to JSONL format (one JSON object per line).
func ExportToJSONL(entries []TraceEntry) ([]byte, error) {
	var result strings.Builder
	for _, entry := range entries {
		data, err := json.Marshal(entry)
		if err != nil {
			return nil, err
		}
		result.Write(data)
		result.WriteString("\n")
	}
	return []byte(result.String()), nil
}
