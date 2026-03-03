package scripting

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/dop251/goja"
)

// TraceEntry represents a single trace log entry.
type TraceEntry struct {
	Timestamp int64           `json:"ts"`
	Direction string          `json:"dir"`
	Message   json.RawMessage `json:"msg"`
}

// TraceReader provides JavaScript API for reading trace files.
type TraceReader struct {
	filename string
	entries  []TraceEntry
	position int
	engine   *Engine
	mu       sync.Mutex
}

// TraceDir holds the trace directory path (set by config).
// Deprecated: Use tracing.GetTraceDir() instead.
var TraceDir = "./data/traces"

// NewTraceReader creates a new trace reader for the given file.
func NewTraceReader(filename string, engine *Engine) (*TraceReader, error) {
	// Validate filename
	if !strings.HasSuffix(filename, ".jsonl") && !strings.HasSuffix(filename, ".jsonl.gz") {
		return nil, fmt.Errorf("invalid filename: must be .jsonl or .jsonl.gz")
	}

	// Security: prevent path traversal
	base := filepath.Base(filename)
	if base != filename || strings.Contains(filename, "..") {
		return nil, fmt.Errorf("invalid filename: path traversal not allowed")
	}

	path := filepath.Join(TraceDir, filename)

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Handle gzip compression
	var reader io.Reader = file
	if strings.HasSuffix(filename, ".gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	}

	// Parse entries
	entries := make([]TraceEntry, 0, 1000)
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var entry TraceEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue // Skip invalid lines
		}

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("no valid entries found in file")
	}

	return &TraceReader{
		filename: filename,
		entries:  entries,
		position: 0,
		engine:   engine,
	}, nil
}

// ToValue converts the TraceReader to a JavaScript object.
func (tr *TraceReader) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// filename - the trace file name
	_ = obj.Set("filename", tr.filename)

	// length - total number of entries
	_ = obj.Set("length", len(tr.entries))

	// position - current iterator position
	_ = obj.DefineAccessorProperty("position", vm.ToValue(func(call goja.FunctionCall) goja.Value {
		tr.mu.Lock()
		defer tr.mu.Unlock()
		return vm.ToValue(tr.position)
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)

	// reset() - reset iterator to beginning
	_ = obj.Set("reset", func(call goja.FunctionCall) goja.Value {
		tr.mu.Lock()
		tr.position = 0
		tr.mu.Unlock()
		return goja.Undefined()
	})

	// hasNext() - check if there are more entries
	_ = obj.Set("hasNext", func(call goja.FunctionCall) goja.Value {
		tr.mu.Lock()
		defer tr.mu.Unlock()
		return vm.ToValue(tr.position < len(tr.entries))
	})

	// next() - get next entry and advance position, returns null at end
	_ = obj.Set("next", func(call goja.FunctionCall) goja.Value {
		tr.mu.Lock()
		defer tr.mu.Unlock()

		if tr.position >= len(tr.entries) {
			return goja.Null()
		}

		entry := tr.entries[tr.position]
		tr.position++

		return tr.entryToValue(vm, entry)
	})

	// get(index) - get entry at specific index without advancing
	_ = obj.Set("get", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("get requires index argument"))
		}
		index := int(call.Arguments[0].ToInteger())

		tr.mu.Lock()
		defer tr.mu.Unlock()

		if index < 0 || index >= len(tr.entries) {
			return goja.Null()
		}

		return tr.entryToValue(vm, tr.entries[index])
	})

	// seek(index) - move iterator to specific position
	_ = obj.Set("seek", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("seek requires index argument"))
		}
		index := int(call.Arguments[0].ToInteger())

		tr.mu.Lock()
		defer tr.mu.Unlock()

		if index < 0 {
			index = 0
		}
		if index > len(tr.entries) {
			index = len(tr.entries)
		}
		tr.position = index

		return goja.Undefined()
	})

	// forEach(callback) - iterate over all entries synchronously
	_ = obj.Set("forEach", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("forEach requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("argument must be a function"))
		}

		tr.mu.Lock()
		entries := tr.entries
		tr.mu.Unlock()

		for i, entry := range entries {
			entryVal := tr.entryToValue(vm, entry)
			_, err := callback(goja.Undefined(), entryVal, vm.ToValue(i))
			if err != nil {
				tr.engine.writeOutput("error", fmt.Sprintf("forEach callback error: %v", err))
				break
			}
		}

		return goja.Undefined()
	})

	// filter(predicate) - return entries matching predicate
	_ = obj.Set("filter", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("filter requires predicate argument"))
		}
		predicate, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("argument must be a function"))
		}

		tr.mu.Lock()
		entries := tr.entries
		tr.mu.Unlock()

		result := make([]goja.Value, 0)
		for i, entry := range entries {
			entryVal := tr.entryToValue(vm, entry)
			ret, err := predicate(goja.Undefined(), entryVal, vm.ToValue(i))
			if err != nil {
				tr.engine.writeOutput("error", fmt.Sprintf("filter predicate error: %v", err))
				break
			}
			if ret.ToBoolean() {
				result = append(result, entryVal)
			}
		}

		return vm.ToValue(result)
	})

	// playback(options) - play all entries with timing
	// options: { speed: 1.0, onMessage: fn, onComplete: fn, direction: "SEND"|"RECV"|"" }
	_ = obj.Set("playback", func(call goja.FunctionCall) goja.Value {
		options := make(map[string]interface{})
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) && !goja.IsNull(call.Arguments[0]) {
			if err := vm.ExportTo(call.Arguments[0], &options); err != nil {
				panic(vm.NewTypeError("invalid options object"))
			}
		}

		speed := 1.0
		if s, ok := options["speed"].(float64); ok && s > 0 {
			speed = s
		}

		var onMessage goja.Callable
		if fn, ok := options["onMessage"]; ok {
			onMessage, _ = goja.AssertFunction(vm.ToValue(fn))
		}

		var onComplete goja.Callable
		if fn, ok := options["onComplete"]; ok {
			onComplete, _ = goja.AssertFunction(vm.ToValue(fn))
		}

		dirFilter := ""
		if d, ok := options["direction"].(string); ok {
			dirFilter = d
		}

		stopCh := make(chan struct{})

		go tr.runPlayback(vm, speed, dirFilter, onMessage, onComplete, stopCh)

		// Return stop function
		return vm.ToValue(func() {
			close(stopCh)
		})
	})

	// entries() - return all entries as array (for small files)
	_ = obj.Set("entries", func(call goja.FunctionCall) goja.Value {
		tr.mu.Lock()
		entries := tr.entries
		tr.mu.Unlock()

		result := make([]goja.Value, len(entries))
		for i, entry := range entries {
			result[i] = tr.entryToValue(vm, entry)
		}

		return vm.ToValue(result)
	})

	return obj
}

func (tr *TraceReader) entryToValue(vm *goja.Runtime, entry TraceEntry) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("ts", entry.Timestamp)
	_ = obj.Set("dir", entry.Direction)

	// Parse message as JSON
	var msg interface{}
	if err := json.Unmarshal(entry.Message, &msg); err != nil {
		_ = obj.Set("msg", string(entry.Message))
	} else {
		_ = obj.Set("msg", msg)
	}

	return obj
}

func (tr *TraceReader) runPlayback(vm *goja.Runtime, speed float64, dirFilter string, onMessage, onComplete goja.Callable, stopCh chan struct{}) {
	tr.mu.Lock()
	entries := tr.entries
	tr.mu.Unlock()

	for i := 0; i < len(entries); i++ {
		select {
		case <-tr.engine.ctx.Done():
			return
		case <-stopCh:
			return
		default:
		}

		entry := entries[i]

		// Apply direction filter
		if dirFilter != "" && entry.Direction != dirFilter {
			continue
		}

		// Calculate delay based on timestamps
		if i > 0 {
			prevTs := entries[i-1].Timestamp
			delayMs := float64(entry.Timestamp-prevTs) / speed
			if delayMs > 0 {
				select {
				case <-time.After(time.Duration(delayMs) * time.Millisecond):
				case <-tr.engine.ctx.Done():
					return
				case <-stopCh:
					return
				}
			}
		}

		// Call onMessage callback
		if onMessage != nil {
			tr.engine.ScheduleCallback(func(vm *goja.Runtime) {
				entryVal := tr.entryToValue(vm, entry)
				_, err := onMessage(goja.Undefined(), entryVal, vm.ToValue(i))
				if err != nil {
					tr.engine.writeOutput("error", fmt.Sprintf("playback onMessage error: %v", err))
				}
			})
		}
	}

	// Call onComplete callback
	if onComplete != nil {
		tr.engine.ScheduleCallback(func(vm *goja.Runtime) {
			_, err := onComplete(goja.Undefined())
			if err != nil {
				tr.engine.writeOutput("error", fmt.Sprintf("playback onComplete error: %v", err))
			}
		})
	}
}

// registerTraceReader registers the openTrace function in the VM.
func (e *Engine) registerTraceReader(vm *goja.Runtime) {
	// openTrace(filename) - opens a trace file and returns a TraceReader
	_ = vm.Set("openTrace", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("openTrace requires filename argument"))
		}
		filename := call.Arguments[0].String()

		reader, err := NewTraceReader(filename, e)
		if err != nil {
			panic(vm.NewGoError(err))
		}

		return reader.ToValue(vm)
	})

	// listTraces() - returns array of available trace files
	_ = vm.Set("listTraces", func(call goja.FunctionCall) goja.Value {
		files, err := listTraceFiles()
		if err != nil {
			panic(vm.NewGoError(err))
		}

		result := make([]string, len(files))
		for i, f := range files {
			result[i] = f.Name()
		}

		return vm.ToValue(result)
	})
}

// listTraceFiles returns a list of trace files in the trace directory.
func listTraceFiles() ([]os.DirEntry, error) {
	entries, err := os.ReadDir(TraceDir)
	if err != nil {
		return nil, err
	}

	var result []os.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".jsonl") || strings.HasSuffix(name, ".jsonl.gz") {
			result = append(result, entry)
		}
	}

	return result, nil
}
