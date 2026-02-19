package tracing

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Writer provides functionality to write trace entries to JSONL files.
// Supports automatic file rotation using lumberjack.
type Writer struct {
	filename string
	logger   *lumberjack.Logger
	mu       sync.Mutex
	closed   bool
}

// WriterConfig configures trace file writing and rotation.
type WriterConfig struct {
	Filename   string // Base filename (e.g., "proxy-name.jsonl")
	MaxSizeMB  int    // Max size in MB before rotation (default: 10)
	MaxBackups int    // Max number of old files to keep (default: 5)
	MaxAgeDays int    // Max age in days to keep files (default: 7)
	Compress   bool   // Compress rotated files with gzip
}

// NewWriter creates a new trace writer with rotation support.
func NewWriter(config WriterConfig) (*Writer, error) {
	if config.Filename == "" {
		return nil, fmt.Errorf("filename is required")
	}

	// Set defaults
	if config.MaxSizeMB == 0 {
		config.MaxSizeMB = 10
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 5
	}
	if config.MaxAgeDays == 0 {
		config.MaxAgeDays = 7
	}

	// Ensure trace directory exists
	traceDir := GetTraceDir()
	if err := os.MkdirAll(traceDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create trace directory: %w", err)
	}

	// Full path to trace file
	path := filepath.Join(traceDir, config.Filename)

	// Create lumberjack logger for rotation
	logger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    config.MaxSizeMB,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAgeDays,
		Compress:   config.Compress,
	}

	return &Writer{
		filename: config.Filename,
		logger:   logger,
	}, nil
}

// Write writes a single trace entry.
func (w *Writer) Write(entry TraceEntry) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return fmt.Errorf("writer is closed")
	}

	// Marshal to JSON
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal entry: %w", err)
	}

	// Write with newline
	_, err = w.logger.Write(append(data, '\n'))
	return err
}

// WriteMessage writes a trace entry from message components.
func (w *Writer) WriteMessage(direction, proxyName string, message []byte) error {
	entry := TraceEntry{
		Timestamp: time.Now().UnixMilli(),
		Direction: direction,
		Proxy:     proxyName,
		Message:   json.RawMessage(message),
	}
	return w.Write(entry)
}

// WriteBatch writes multiple entries at once.
func (w *Writer) WriteBatch(entries []TraceEntry) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return fmt.Errorf("writer is closed")
	}

	for _, entry := range entries {
		data, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal entry: %w", err)
		}

		if _, err := w.logger.Write(append(data, '\n')); err != nil {
			return err
		}
	}

	return nil
}

// Rotate forces a rotation of the current log file.
func (w *Writer) Rotate() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return fmt.Errorf("writer is closed")
	}

	return w.logger.Rotate()
}

// Close closes the writer and flushes any buffered data.
func (w *Writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return nil
	}

	w.closed = true
	return w.logger.Close()
}

// GetWriter returns the underlying io.Writer.
func (w *Writer) GetWriter() io.Writer {
	return w.logger
}

// AppendToFile appends entries to an existing trace file without rotation.
func AppendToFile(filename string, entries []TraceEntry) error {
	path := filepath.Join(GetTraceDir(), filename)

	// Open file in append mode
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Write each entry
	for _, entry := range entries {
		data, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal entry: %w", err)
		}

		if _, err := file.Write(append(data, '\n')); err != nil {
			return fmt.Errorf("failed to write entry: %w", err)
		}
	}

	return nil
}

// WriteToFile writes entries to a new trace file (overwrites if exists).
func WriteToFile(filename string, entries []TraceEntry) error {
	path := filepath.Join(GetTraceDir(), filename)

	// Create file (truncate if exists)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write each entry
	for _, entry := range entries {
		data, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal entry: %w", err)
		}

		if _, err := file.Write(append(data, '\n')); err != nil {
			return fmt.Errorf("failed to write entry: %w", err)
		}
	}

	return nil
}

// CopyTraceTo copies a trace file to a new location.
func CopyTraceTo(srcFilename, dstPath string) error {
	srcPath := filepath.Join(GetTraceDir(), srcFilename)

	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
