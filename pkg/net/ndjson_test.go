package net

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNDJSONScanner(t *testing.T) {
	t.Run("creates scanner with default values", func(t *testing.T) {
		scanner := NewNDJSONScanner(0, 1)
		assert.NotNil(t, scanner)
		assert.Equal(t, time.Duration(0), scanner.sleep)
		assert.Equal(t, 1, scanner.repeat)
	})

	t.Run("creates scanner with custom values", func(t *testing.T) {
		sleep := 100 * time.Millisecond
		repeat := 5
		scanner := NewNDJSONScanner(sleep, repeat)
		assert.NotNil(t, scanner)
		assert.Equal(t, sleep, scanner.sleep)
		assert.Equal(t, repeat, scanner.repeat)
	})
}

func TestNDJSONScannerScan(t *testing.T) {
	t.Run("scans single line", func(t *testing.T) {
		input := strings.NewReader("line1\n")
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err := scanner.Scan(input, output)
		assert.NoError(t, err)
		assert.Equal(t, "line1", output.String())
	})

	t.Run("scans multiple lines", func(t *testing.T) {
		input := strings.NewReader("line1\nline2\nline3\n")
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err := scanner.Scan(input, output)
		assert.NoError(t, err)
		assert.Equal(t, "line1line2line3", output.String())
	})

	t.Run("handles empty input", func(t *testing.T) {
		input := strings.NewReader("")
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err := scanner.Scan(input, output)
		assert.NoError(t, err)
		assert.Empty(t, output.String())
	})

	t.Run("repeats scan multiple times", func(t *testing.T) {
		input := strings.NewReader("line1\nline2\n")
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 3)

		err := scanner.Scan(input, output)
		assert.NoError(t, err)
		// With repeat=3, only scans once since reader is exhausted
		assert.Equal(t, "line1line2", output.String())
	})

	t.Run("handles lines without trailing newline", func(t *testing.T) {
		input := strings.NewReader("line1\nline2")
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err := scanner.Scan(input, output)
		assert.NoError(t, err)
		assert.Equal(t, "line1line2", output.String())
	})

	t.Run("respects sleep duration", func(t *testing.T) {
		input := strings.NewReader("line1\nline2\n")
		output := &bytes.Buffer{}
		sleep := 10 * time.Millisecond
		scanner := NewNDJSONScanner(sleep, 1)

		start := time.Now()
		err := scanner.Scan(input, output)
		elapsed := time.Since(start)

		assert.NoError(t, err)
		assert.Equal(t, "line1line2", output.String())
		// Should take at least 2 * sleep (one per line)
		assert.GreaterOrEqual(t, elapsed, 2*sleep)
	})
}

func TestNDJSONScannerScanFile(t *testing.T) {
	t.Run("scans file successfully", func(t *testing.T) {
		// Create a temporary file
		tmpfile := t.TempDir() + "/test.ndjson"
		content := "line1\nline2\nline3\n"
		err := writeFile(tmpfile, content)
		require.NoError(t, err)

		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err = scanner.ScanFile(tmpfile, output)
		assert.NoError(t, err)
		assert.Equal(t, "line1line2line3", output.String())
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err := scanner.ScanFile("/nonexistent/file.ndjson", output)
		assert.Error(t, err)
	})

	t.Run("handles empty file", func(t *testing.T) {
		tmpfile := t.TempDir() + "/empty.ndjson"
		err := writeFile(tmpfile, "")
		require.NoError(t, err)

		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err = scanner.ScanFile(tmpfile, output)
		assert.NoError(t, err)
		assert.Empty(t, output.String())
	})

	t.Run("scans file with JSON lines", func(t *testing.T) {
		tmpfile := t.TempDir() + "/json.ndjson"
		content := `{"type":"call","symbol":"test1"}
{"type":"signal","symbol":"test2"}
{"type":"state","symbol":"test3"}
`
		err := writeFile(tmpfile, content)
		require.NoError(t, err)

		output := &bytes.Buffer{}
		scanner := NewNDJSONScanner(0, 1)

		err = scanner.ScanFile(tmpfile, output)
		assert.NoError(t, err)
		// Each line written without newline separator
		assert.Contains(t, output.String(), `{"type":"call","symbol":"test1"}`)
		assert.Contains(t, output.String(), `{"type":"signal","symbol":"test2"}`)
		assert.Contains(t, output.String(), `{"type":"state","symbol":"test3"}`)
	})
}

// Helper function to write test files
func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
