package tracing

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	traceDir     string
	traceDirOnce sync.Once
)

// TraceEntry represents a single trace log entry in JSONL format.
type TraceEntry struct {
	Timestamp int64           `json:"ts"`     // Unix timestamp in milliseconds
	Direction string          `json:"dir"`    // "SEND" or "RECV"
	Proxy     string          `json:"proxy"`  // Proxy name
	Message   json.RawMessage `json:"msg"`    // Raw message as JSON
}

// TraceFileInfo contains metadata about a trace file.
type TraceFileInfo struct {
	Name      string    `json:"name"`      // File name (e.g., "proxy-name.jsonl")
	Path      string    `json:"path"`      // Full file path
	Size      int64     `json:"size"`      // File size in bytes
	ModTime   time.Time `json:"modTime"`   // Last modification time
	ProxyName string    `json:"proxyName"` // Extracted proxy name
}

// TraceStats holds summary statistics about trace files.
type TraceStats struct {
	FileCount  int     `json:"fileCount"`  // Number of trace files
	TotalBytes int64   `json:"totalBytes"` // Total size in bytes
	TotalMB    float64 `json:"totalMB"`    // Total size in MB
	TraceDir   string  `json:"traceDir"`   // Trace directory path
}

// FilterOptions specifies criteria for filtering trace entries.
type FilterOptions struct {
	Direction  string   // Filter by direction: "SEND", "RECV", or empty for all
	ProxyNames []string // Filter by proxy names (empty = all)
	StartTime  int64    // Filter entries >= start time (Unix ms)
	EndTime    int64    // Filter entries <= end time (Unix ms)
	Limit      int      // Maximum number of entries to return (0 = unlimited)
}

// DefaultTraceDir returns the default trace directory.
func DefaultTraceDir() string {
	return "./data/traces"
}

// SetTraceDir sets the trace directory. Should be called early in startup.
func SetTraceDir(dir string) error {
	traceDirOnce.Do(func() {
		if dir != "" {
			traceDir = dir
		} else if envDir := os.Getenv("APIGEAR_TRACE_DIR"); envDir != "" {
			traceDir = envDir
		} else {
			traceDir = DefaultTraceDir()
		}

		// Ensure directory exists
		if err := os.MkdirAll(traceDir, 0755); err != nil {
			// Fall back to current directory
			traceDir = "."
		}
	})
	return nil
}

// GetTraceDir returns the configured trace directory.
func GetTraceDir() string {
	if traceDir == "" {
		SetTraceDir("")
	}
	return traceDir
}

// ListTraceFiles returns all .jsonl and .jsonl.gz files in the trace directory,
// sorted by modification time (newest first).
func ListTraceFiles() ([]TraceFileInfo, error) {
	dir := GetTraceDir()

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []TraceFileInfo{}, nil
		}
		return nil, err
	}

	var files []TraceFileInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// Match both .jsonl and .jsonl.gz (rotated/compressed files)
		name := entry.Name()
		if !strings.HasSuffix(name, ".jsonl") && !strings.HasSuffix(name, ".jsonl.gz") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		// Extract proxy name from filename (format: proxyname.jsonl or proxyname-timestamp.jsonl.gz)
		proxyName := extractProxyName(name)

		files = append(files, TraceFileInfo{
			Name:      name,
			Path:      filepath.Join(dir, name),
			Size:      info.Size(),
			ModTime:   info.ModTime(),
			ProxyName: proxyName,
		})
	}

	// Sort by modification time, newest first
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.After(files[j].ModTime)
	})

	return files, nil
}

// DeleteTraceFile deletes a trace file by name.
func DeleteTraceFile(name string) error {
	// Security: only allow deleting trace files in the trace directory
	if !strings.HasSuffix(name, ".jsonl") && !strings.HasSuffix(name, ".jsonl.gz") {
		return os.ErrInvalid
	}
	if strings.Contains(name, "/") || strings.Contains(name, "\\") || strings.Contains(name, "..") {
		return os.ErrInvalid
	}

	path := filepath.Join(GetTraceDir(), name)
	return os.Remove(path)
}

// GetTraceStats returns summary statistics about trace files.
func GetTraceStats() TraceStats {
	files, err := ListTraceFiles()
	if err != nil {
		return TraceStats{TraceDir: GetTraceDir()}
	}

	var totalBytes int64
	for _, f := range files {
		totalBytes += f.Size
	}

	return TraceStats{
		FileCount:  len(files),
		TotalBytes: totalBytes,
		TotalMB:    float64(totalBytes) / (1024 * 1024),
		TraceDir:   GetTraceDir(),
	}
}

// extractProxyName extracts the proxy name from a trace filename.
// Handles formats like: "proxy.jsonl", "proxy-2024-01-01T00-00-00.000.jsonl.gz"
func extractProxyName(filename string) string {
	// Remove extensions
	baseName := strings.TrimSuffix(strings.TrimSuffix(filename, ".gz"), ".jsonl")

	// For rotated files like "server-2024-01-01T00-00-00.000"
	if idx := strings.Index(baseName, "-20"); idx > 0 {
		return baseName[:idx]
	}

	return baseName
}
