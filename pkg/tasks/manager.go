package tasks

import (
	"context"
	"errors"
	"sync"
)

// ErrTaskNotFound is returned when a task is not found
var ErrTaskNotFound = errors.New("task not found")

// TaskManager allows you to create tasks and run them
type TaskManager struct {
	sync.RWMutex
	tasks map[string]*TaskItem
}

// New creates a new task manager
func New() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*TaskItem),
	}
}

// Register creates a new task
func (tm *TaskManager) Register(name string, tf TaskFunc) *TaskItem {
	task := NewTaskItem(name, tf)
	tm.AddTask(task)
	return task
}

// AddTask adds a task to the task manager
func (tm *TaskManager) AddTask(task *TaskItem) {
	if task == nil {
		return
	}
	if tm.Has(task.name) {
		return
	}
	tm.Lock()
	defer tm.Unlock()
	tm.tasks[task.name] = task
}

// RmTask removes a task from the task manager
func (tm *TaskManager) RmTask(name string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	tm.Lock()
	defer tm.Unlock()
	delete(tm.tasks, name)
	return nil
}

// Get returns a task
func (tm *TaskManager) Get(name string) *TaskItem {
	tm.RLock()
	defer tm.RUnlock()
	task, ok := tm.tasks[name]
	if !ok {
		return nil
	}
	return task
}

// Run runs a task
func (tm *TaskManager) Run(ctx context.Context, name string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	task.Run(ctx)
	return nil
}

// Watch watches a task
func (tm *TaskManager) Watch(ctx context.Context, name string, dependencies ...string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	task.Watch(ctx, dependencies...)
	return nil
}

// Names returns the names of all the tasks
func (tm *TaskManager) Names() []string {
	tm.RLock()
	defer tm.RUnlock()
	var names []string
	for name := range tm.tasks {
		names = append(names, name)
	}
	return names
}

// Has returns true if the task exists
func (tm *TaskManager) Has(name string) bool {
	tm.RLock()
	defer tm.RUnlock()
	_, ok := tm.tasks[name]
	return ok
}

// Cancel cancels a task
func (tm *TaskManager) Cancel(name string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	task.Cancel()
	return nil
}

// CancelAll cancels all the tasks
func (tm *TaskManager) CancelAll() {
	tm.RLock()
	defer tm.RUnlock()
	for _, task := range tm.tasks {
		task.Cancel()
	}
}
