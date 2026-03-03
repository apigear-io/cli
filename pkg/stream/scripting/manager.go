package scripting

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Manager manages multiple running scripts.
type Manager struct {
	engines    map[string]*Engine
	enginesMu  sync.RWMutex
	scriptsDir string
	stats      StatsRecorder

	// For generating unique engine IDs
	nextID atomic.Int64

	// Stop channel and once
	stopCh   chan struct{}
	stopOnce sync.Once
}

// NewManager creates a new script manager.
func NewManager(scriptsDir string, stats StatsRecorder) *Manager {
	m := &Manager{
		engines:    make(map[string]*Engine),
		scriptsDir: scriptsDir,
		stats:      stats,
		stopCh:     make(chan struct{}),
	}

	// Ensure scripts directory exists
	if scriptsDir != "" {
		_ = os.MkdirAll(scriptsDir, 0755)
	}

	return m
}

// RunScript runs a script with the given name.
// Scripts run forever until manually stopped or until they call exit().
// If a script with the same name is already running, it will be stopped first (restart behavior).
func (m *Manager) RunScript(name, code string) (string, error) {
	// Check if a script with this name is already running and stop it (restart behavior)
	m.enginesMu.RLock()
	var enginesToStop []*Engine
	for _, existingEngine := range m.engines {
		if existingEngine.Name() == name && !existingEngine.IsStopped() {
			enginesToStop = append(enginesToStop, existingEngine)
		}
	}
	m.enginesMu.RUnlock()

	// Stop existing scripts outside the lock
	for _, engine := range enginesToStop {
		engine.Stop()
	}

	// Brief wait to allow cleanup to start
	if len(enginesToStop) > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	id := fmt.Sprintf("script-%d", m.nextID.Add(1))

	engine := NewEngine(id, name)
	engine.SetStats(m.stats)

	// Set up auto-cleanup when engine stops
	// Keep engine around for 30 seconds after stop to allow clients to fetch output
	engine.SetOnStopCallback(func() {
		go func() {
			time.Sleep(30 * time.Second)
			m.enginesMu.Lock()
			delete(m.engines, id)
			m.enginesMu.Unlock()
		}()
	})

	m.enginesMu.Lock()
	m.engines[id] = engine
	m.enginesMu.Unlock()

	// Run the script asynchronously
	go func() {
		err := engine.RunAsync(code)
		if err != nil {
			engine.writeOutput("error", fmt.Sprintf("Script error: %v", err))
			engine.Stop()
		}
		// Script runs forever until stopped or exit() is called
	}()

	return id, nil
}

// StopScript stops a running script by ID.
// This is idempotent - stopping an already-stopped script returns success.
func (m *Manager) StopScript(id string) error {
	m.enginesMu.RLock()
	engine, ok := m.engines[id]
	m.enginesMu.RUnlock()

	if !ok {
		// Script already stopped or doesn't exist - that's fine
		return nil
	}

	// Stop triggers the onStopCallback which handles cleanup
	engine.Stop()
	return nil
}

// GetRunningScripts returns information about all running scripts.
// Note: Only returns scripts that are still running. Stopped scripts are filtered out
// even if they remain in memory during the cleanup grace period.
func (m *Manager) GetRunningScripts() []ScriptInfo {
	m.enginesMu.RLock()
	defer m.enginesMu.RUnlock()

	result := make([]ScriptInfo, 0, len(m.engines))
	for id, engine := range m.engines {
		// Skip stopped engines (they remain in map for grace period but shouldn't show as running)
		if engine.IsStopped() {
			continue
		}

		scriptType := ScriptTypeClient
		if engine.GetBackendServer() != nil {
			scriptType = ScriptTypeBackend
		}
		result = append(result, ScriptInfo{
			ID:   id,
			Name: engine.Name(),
			Type: scriptType,
		})
	}
	return result
}

// GetEngine returns an engine by ID.
func (m *Manager) GetEngine(id string) *Engine {
	m.enginesMu.RLock()
	defer m.enginesMu.RUnlock()
	return m.engines[id]
}

// ScriptType indicates whether a script is a client or backend script.
type ScriptType string

const (
	// ScriptTypeClient is a client script that uses connect() to connect to backends.
	ScriptTypeClient ScriptType = "client"
	// ScriptTypeBackend is a backend script that uses createBackend() to create a server.
	ScriptTypeBackend ScriptType = "backend"
)

// ScriptInfo contains information about a running script.
type ScriptInfo struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Type ScriptType `json:"type"`
}

// ScriptFile contains information about a saved script file.
type ScriptFile struct {
	Name    string     `json:"name"`
	Type    ScriptType `json:"type"`
	ModTime int64      `json:"modTime"`
	Code    string     `json:"code,omitempty"`
}

// ListScripts returns all saved scripts.
func (m *Manager) ListScripts() ([]string, error) {
	if m.scriptsDir == "" {
		return nil, fmt.Errorf("scripts directory not configured")
	}

	entries, err := os.ReadDir(m.scriptsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var scripts []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".js") {
			scripts = append(scripts, strings.TrimSuffix(entry.Name(), ".js"))
		}
	}
	return scripts, nil
}

// LoadScript loads a script by name.
func (m *Manager) LoadScript(name string) (string, error) {
	if m.scriptsDir == "" {
		return "", fmt.Errorf("scripts directory not configured")
	}

	path := filepath.Join(m.scriptsDir, name+".js")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// LoadScriptWithInfo loads a script with its modification time.
func (m *Manager) LoadScriptWithInfo(name string) (*ScriptFile, error) {
	if m.scriptsDir == "" {
		return nil, fmt.Errorf("scripts directory not configured")
	}

	path := filepath.Join(m.scriptsDir, name+".js")

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &ScriptFile{
		Name:    name,
		ModTime: info.ModTime().UnixMilli(),
		Code:    string(data),
	}, nil
}

// GetScriptModTime returns the modification time of a script.
func (m *Manager) GetScriptModTime(name string) (int64, error) {
	if m.scriptsDir == "" {
		return 0, fmt.Errorf("scripts directory not configured")
	}

	path := filepath.Join(m.scriptsDir, name+".js")
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.ModTime().UnixMilli(), nil
}

// ErrConflict is returned when a save conflicts with a newer version.
var ErrConflict = fmt.Errorf("conflict: script was modified by another user")

// SaveScript saves a script with the given name.
func (m *Manager) SaveScript(name, code string) error {
	if m.scriptsDir == "" {
		return fmt.Errorf("scripts directory not configured")
	}

	// Validate name (alphanumeric, dashes, underscores only)
	for _, c := range name {
		isLower := c >= 'a' && c <= 'z'
		isUpper := c >= 'A' && c <= 'Z'
		isDigit := c >= '0' && c <= '9'
		isSpecial := c == '-' || c == '_'
		if !isLower && !isUpper && !isDigit && !isSpecial {
			return fmt.Errorf("invalid script name: %s", name)
		}
	}

	path := filepath.Join(m.scriptsDir, name+".js")
	return os.WriteFile(path, []byte(code), 0644)
}

// SaveScriptWithCheck saves a script, checking that it hasn't been modified since expectedModTime.
// If expectedModTime is 0, the check is skipped (for new scripts or force saves).
func (m *Manager) SaveScriptWithCheck(name, code string, expectedModTime int64) (int64, error) {
	if m.scriptsDir == "" {
		return 0, fmt.Errorf("scripts directory not configured")
	}

	// Validate name
	for _, c := range name {
		isLower := c >= 'a' && c <= 'z'
		isUpper := c >= 'A' && c <= 'Z'
		isDigit := c >= '0' && c <= '9'
		isSpecial := c == '-' || c == '_'
		if !isLower && !isUpper && !isDigit && !isSpecial {
			return 0, fmt.Errorf("invalid script name: %s", name)
		}
	}

	path := filepath.Join(m.scriptsDir, name+".js")

	// Check current mod time if expectedModTime is provided
	if expectedModTime > 0 {
		info, err := os.Stat(path)
		if err == nil {
			currentModTime := info.ModTime().UnixMilli()
			if currentModTime != expectedModTime {
				return currentModTime, ErrConflict
			}
		}
		// If file doesn't exist, that's fine - we're creating it
	}

	if err := os.WriteFile(path, []byte(code), 0644); err != nil {
		return 0, err
	}

	// Get new mod time
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.ModTime().UnixMilli(), nil
}

// DeleteScript deletes a script by name.
func (m *Manager) DeleteScript(name string) error {
	if m.scriptsDir == "" {
		return fmt.Errorf("scripts directory not configured")
	}

	path := filepath.Join(m.scriptsDir, name+".js")
	return os.Remove(path)
}

// Close stops all scripts and cleans up.
func (m *Manager) Close() {
	m.stopOnce.Do(func() {
		close(m.stopCh)

		// Collect engines to stop (to avoid deadlock with onStopCallback)
		m.enginesMu.Lock()
		enginesToStop := make([]*Engine, 0, len(m.engines))
		for _, engine := range m.engines {
			enginesToStop = append(enginesToStop, engine)
		}
		m.engines = make(map[string]*Engine)
		m.enginesMu.Unlock()

		// Stop engines outside the lock
		for _, engine := range enginesToStop {
			engine.Stop()
		}
	})
}

// LoadAndStart loads a script by name and runs it.
// This implements the BackendStarter interface for proxy backend mode.
// The script is expected to call createBackend(wsUrl) which starts the WebSocket server.
func (m *Manager) LoadAndStart(scriptName, listenAddr, _ string) error {
	// Check if already running
	m.enginesMu.RLock()
	for _, engine := range m.engines {
		if engine.Name() == scriptName && engine.GetBackendServer() != nil {
			m.enginesMu.RUnlock()
			return fmt.Errorf("backend script already running: %s", scriptName)
		}
	}
	m.enginesMu.RUnlock()

	// Load script
	code, err := m.LoadScript(scriptName)
	if err != nil {
		return fmt.Errorf("failed to load script %s: %w", scriptName, err)
	}

	// Create engine
	id := fmt.Sprintf("backend-%d", m.nextID.Add(1))
	engine := NewEngine(id, scriptName)
	engine.SetStats(m.stats)

	// Set up auto-cleanup when engine stops
	// Keep engine around for 30 seconds after stop to allow clients to fetch output
	engine.SetOnStopCallback(func() {
		go func() {
			time.Sleep(30 * time.Second)
			m.enginesMu.Lock()
			delete(m.engines, id)
			m.enginesMu.Unlock()
		}()
	})

	m.enginesMu.Lock()
	m.engines[id] = engine
	m.enginesMu.Unlock()

	// Run the script - it will call createBackend(wsUrl) which starts the server
	if err := engine.RunAsync(code); err != nil {
		m.enginesMu.Lock()
		delete(m.engines, id)
		m.enginesMu.Unlock()
		return fmt.Errorf("failed to run script %s: %w", scriptName, err)
	}

	return nil
}

// StopScriptByName stops a running script by name.
func (m *Manager) StopScriptByName(name string) error {
	m.enginesMu.RLock()
	var engineToStop *Engine
	var idToStop string
	for id, engine := range m.engines {
		if engine.Name() == name {
			engineToStop = engine
			idToStop = id
			break
		}
	}
	m.enginesMu.RUnlock()

	if engineToStop == nil {
		return fmt.Errorf("script not found: %s", name)
	}

	engineToStop.Stop()
	_ = idToStop // Used for logging in full implementation
	return nil
}

// GetEngineByName returns an engine by script name.
func (m *Manager) GetEngineByName(name string) *Engine {
	m.enginesMu.RLock()
	defer m.enginesMu.RUnlock()

	for _, engine := range m.engines {
		if engine.Name() == name {
			return engine
		}
	}
	return nil
}
