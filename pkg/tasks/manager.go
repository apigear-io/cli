package tasks

import (
	"context"
	"errors"
	"sync"
)

// ErrTaskNotFound is returned when a task is not found
var ErrTaskNotFound = errors.New("task not found")

// TaskManager provides a simple registry for managing multiple tasks
type TaskManager struct {
	mu      sync.RWMutex
	tasks   map[string]*taskEntry
	hooks   []func(*TaskEvent)
	hooksMu sync.RWMutex
}

type taskEntry struct {
	task *Task
	fn   TaskFunc
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*taskEntry),
		hooks: make([]func(*TaskEvent), 0),
	}
}

// AddHook adds an event hook function
func (tm *TaskManager) AddHook(fn func(*TaskEvent)) {
	tm.hooksMu.Lock()
	defer tm.hooksMu.Unlock()
	tm.hooks = append(tm.hooks, fn)
}

// fireHook fires event hooks (if any exist)
func (tm *TaskManager) fireHook(name string, state TaskState) {
	tm.hooksMu.RLock()
	hooks := make([]func(*TaskEvent), len(tm.hooks))
	copy(hooks, tm.hooks)
	tm.hooksMu.RUnlock()

	if len(hooks) > 0 {
		event := &TaskEvent{
			Name:  name,
			State: state,
			Meta:  map[string]interface{}{},
		}
		for _, hook := range hooks {
			hook(event)
		}
	}
}

// Register creates and registers a task (meta is ignored for simplicity)
func (tm *TaskManager) Register(name string, meta map[string]interface{}, fn TaskFunc) *Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Remove existing task if present
	if entry, exists := tm.tasks[name]; exists {
		entry.task.Cancel()
		entry.task.CancelWatch()
		tm.fireHook(name, TaskStateRemoved)
	}

	task := NewTask()
	tm.tasks[name] = &taskEntry{
		task: task,
		fn:   fn,
	}
	tm.fireHook(name, TaskStateAdded)
	return task
}

// Run runs a registered task once
func (tm *TaskManager) Run(ctx context.Context, name string) error {
	tm.mu.RLock()
	entry, exists := tm.tasks[name]
	tm.mu.RUnlock()

	if !exists {
		return ErrTaskNotFound
	}

	tm.fireHook(name, TaskStateRunning)
	err := entry.task.Run(ctx, entry.fn)
	if err != nil {
		tm.fireHook(name, TaskStateFailed)
		return err
	}
	tm.fireHook(name, TaskStateFinished)
	return nil
}

// Watch runs a registered task and watches files for changes
func (tm *TaskManager) Watch(ctx context.Context, name string, dependencies ...string) error {
	tm.mu.RLock()
	entry, exists := tm.tasks[name]
	tm.mu.RUnlock()

	if !exists {
		return ErrTaskNotFound
	}

	tm.fireHook(name, TaskStateWatching)
	go func() {
		if err := entry.task.Watch(ctx, entry.fn, dependencies...); err != nil {
			log.Error().Err(err).Str("task", name).Msg("watch failed")
			tm.fireHook(name, TaskStateFailed)
		}
	}()
	return nil
}

// Cancel cancels a registered task
func (tm *TaskManager) Cancel(name string) error {
	tm.mu.RLock()
	entry, exists := tm.tasks[name]
	tm.mu.RUnlock()

	if !exists {
		return ErrTaskNotFound
	}

	entry.task.Cancel()
	entry.task.CancelWatch()
	tm.fireHook(name, TaskStateStopped)
	return nil
}

// CancelAll cancels all registered tasks
func (tm *TaskManager) CancelAll() {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	for _, entry := range tm.tasks {
		entry.task.Cancel()
		entry.task.CancelWatch()
	}
}

// Has returns true if the task exists
func (tm *TaskManager) Has(name string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	_, exists := tm.tasks[name]
	return exists
}

// Names returns the names of all registered tasks
func (tm *TaskManager) Names() []string {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	names := make([]string, 0, len(tm.tasks))
	for name := range tm.tasks {
		names = append(names, name)
	}
	return names
}
