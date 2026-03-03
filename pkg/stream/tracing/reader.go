package tracing

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Reader provides functionality to read trace files.
type Reader struct {
	filename string
	file     *os.File
	gzReader *gzip.Reader
	scanner  *bufio.Scanner
	closed   bool
}

// NewReader creates a new trace file reader.
// Supports both .jsonl and .jsonl.gz (gzip compressed) files.
func NewReader(filename string) (*Reader, error) {
	// Security: prevent path traversal
	base := filepath.Base(filename)
	if base != filename || strings.Contains(filename, "..") {
		return nil, fmt.Errorf("invalid filename: path traversal not allowed")
	}

	// Validate extension
	if !strings.HasSuffix(filename, ".jsonl") && !strings.HasSuffix(filename, ".jsonl.gz") {
		return nil, fmt.Errorf("invalid filename: must be .jsonl or .jsonl.gz")
	}

	path := filepath.Join(GetTraceDir(), filename)

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	reader := &Reader{
		filename: filename,
		file:     file,
	}

	// Handle gzip compression
	var scanReader io.Reader = file
	if strings.HasSuffix(filename, ".gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		reader.gzReader = gzReader
		scanReader = gzReader
	}

	// Create scanner with large buffer for big messages
	scanner := bufio.NewScanner(scanReader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024) // 1MB max line size

	reader.scanner = scanner

	return reader, nil
}

// ReadAll reads all entries from the trace file.
func (r *Reader) ReadAll() ([]TraceEntry, error) {
	if r.closed {
		return nil, fmt.Errorf("reader is closed")
	}

	var entries []TraceEntry
	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			continue
		}

		var entry TraceEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			// Skip invalid lines
			continue
		}

		entries = append(entries, entry)
	}

	if err := r.scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return entries, nil
}

// ReadFiltered reads entries matching the filter criteria.
func (r *Reader) ReadFiltered(filter FilterOptions) ([]TraceEntry, error) {
	if r.closed {
		return nil, fmt.Errorf("reader is closed")
	}

	var entries []TraceEntry
	count := 0

	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			continue
		}

		var entry TraceEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			// Skip invalid lines
			continue
		}

		// Apply filters
		if !matchesFilter(entry, filter) {
			continue
		}

		entries = append(entries, entry)
		count++

		// Check limit
		if filter.Limit > 0 && count >= filter.Limit {
			break
		}
	}

	if err := r.scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return entries, nil
}

// ForEach iterates over all entries and calls the callback for each one.
// Returns early if the callback returns an error.
func (r *Reader) ForEach(callback func(entry TraceEntry) error) error {
	if r.closed {
		return fmt.Errorf("reader is closed")
	}

	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			continue
		}

		var entry TraceEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			// Skip invalid lines
			continue
		}

		if err := callback(entry); err != nil {
			return err
		}
	}

	if err := r.scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// Close closes the reader and releases resources.
func (r *Reader) Close() error {
	if r.closed {
		return nil
	}

	r.closed = true

	if r.gzReader != nil {
		r.gzReader.Close()
	}
	if r.file != nil {
		return r.file.Close()
	}

	return nil
}

// matchesFilter checks if an entry matches the filter criteria.
func matchesFilter(entry TraceEntry, filter FilterOptions) bool {
	// Direction filter
	if filter.Direction != "" && entry.Direction != filter.Direction {
		return false
	}

	// Proxy name filter
	if len(filter.ProxyNames) > 0 {
		found := false
		for _, name := range filter.ProxyNames {
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
	if filter.StartTime > 0 && entry.Timestamp < filter.StartTime {
		return false
	}
	if filter.EndTime > 0 && entry.Timestamp > filter.EndTime {
		return false
	}

	return true
}

// ReadTraceFile is a convenience function to read all entries from a file.
func ReadTraceFile(filename string) ([]TraceEntry, error) {
	reader, err := NewReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return reader.ReadAll()
}

// ReadTraceFileFiltered is a convenience function to read filtered entries from a file.
func ReadTraceFileFiltered(filename string, filter FilterOptions) ([]TraceEntry, error) {
	reader, err := NewReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return reader.ReadFiltered(filter)
}
