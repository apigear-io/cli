package logging

import (
	"encoding/json"
	"sync"
	"time"
)

// LogLevel represents the severity of a log entry
type LogLevel string

const (
	LevelDebug LogLevel = "DEBUG"
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// Logger is an in-memory logger with a circular buffer
type Logger struct {
	mu      sync.RWMutex
	entries []LogEntry
	maxSize int
	index   int // Current write position
}

// NewLogger creates a new logger with the specified buffer size
func NewLogger(maxSize int) *Logger {
	if maxSize <= 0 {
		maxSize = 1000
	}
	return &Logger{
		entries: make([]LogEntry, 0, maxSize),
		maxSize: maxSize,
	}
}

// log adds a log entry
func (l *Logger) log(level LogLevel, message string, fields map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}

	// Circular buffer logic
	if len(l.entries) < l.maxSize {
		l.entries = append(l.entries, entry)
	} else {
		l.entries[l.index] = entry
		l.index = (l.index + 1) % l.maxSize
	}
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.log(LevelDebug, message, fields)
}

// Info logs an info message
func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.log(LevelWarn, message, fields)
}

// Error logs an error message
func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.log(LevelError, message, fields)
}

// GetEntries returns all log entries in chronological order
func (l *Logger) GetEntries() []LogEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if len(l.entries) < l.maxSize {
		// Buffer not full yet, return in order
		result := make([]LogEntry, len(l.entries))
		copy(result, l.entries)
		return result
	}

	// Buffer is full, need to reorder from index
	result := make([]LogEntry, l.maxSize)
	copy(result, l.entries[l.index:])
	copy(result[l.maxSize-l.index:], l.entries[:l.index])
	return result
}

// GetEntriesFiltered returns log entries filtered by level and search term
func (l *Logger) GetEntriesFiltered(level LogLevel, search string) []LogEntry {
	entries := l.GetEntries()
	if level == "" && search == "" {
		return entries
	}

	filtered := make([]LogEntry, 0)
	for _, entry := range entries {
		// Level filter
		if level != "" && entry.Level != level {
			continue
		}

		// Search filter (case-insensitive substring match in message or fields)
		if search != "" {
			if !contains(entry.Message, search) && !containsInFields(entry.Fields, search) {
				continue
			}
		}

		filtered = append(filtered, entry)
	}

	return filtered
}

// Clear removes all log entries
func (l *Logger) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = make([]LogEntry, 0, l.maxSize)
	l.index = 0
}

// Helper functions

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && indexOf(s, substr) >= 0))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func containsInFields(fields map[string]interface{}, search string) bool {
	if fields == nil {
		return false
	}

	// Convert fields to JSON string and search
	data, err := json.Marshal(fields)
	if err != nil {
		return false
	}
	return contains(string(data), search)
}

// Global logger instance
var globalLogger *Logger

func init() {
	globalLogger = NewLogger(1000)
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	return globalLogger
}

// Convenience functions for global logger

// Debug logs a debug message to the global logger
func Debug(message string, fields map[string]interface{}) {
	globalLogger.Debug(message, fields)
}

// Info logs an info message to the global logger
func Info(message string, fields map[string]interface{}) {
	globalLogger.Info(message, fields)
}

// Warn logs a warning message to the global logger
func Warn(message string, fields map[string]interface{}) {
	globalLogger.Warn(message, fields)
}

// Error logs an error message to the global logger
func Error(message string, fields map[string]interface{}) {
	globalLogger.Error(message, fields)
}
