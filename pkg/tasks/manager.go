package tasks

import (
	"context"
	"errors"
	"sync"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/objectlink-core-go/log"
)

// ErrTaskNotFound is returned when a task is not found
var ErrTaskNotFound = errors.New("task not found")

// TaskManager allows you to create tasks and run them
type TaskManager struct {
	sync.RWMutex
	tasks map[string]*TaskItem
	*helper.EventEmitter[*TaskEvent]
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:        make(map[string]*TaskItem),
		EventEmitter: helper.NewEventEmitter[*TaskEvent](),
	}
}

// Register creates a new task
func (tm *TaskManager) Register(name string, meta map[string]interface{}, tf TaskFunc) *TaskItem {
	if tm.Has(name) {
		tm.RmTask(name)
	}
	task := NewTaskItem(name, meta, tf)
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
	tm.Emit(NewTaskEvent(task, TaskStateAdded))
}

// RmTask removes a task from the task manager
func (tm *TaskManager) RmTask(name string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	task.Cancel()
	tm.Lock()
	defer tm.Unlock()
	delete(tm.tasks, name)
	tm.Emit(NewTaskEvent(task, TaskStateRemoved))
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
	tm.Emit(NewTaskEvent(task, TaskStateRunning))
	err := task.Run(ctx)
	if err != nil {
		log.Error().Err(err).Str("task", name).Msg("failed to run task")
		tm.Emit(NewTaskEvent(task, TaskStateFailed))
		return err
	}
	tm.Emit(NewTaskEvent(task, TaskStateFinished))
	return nil
}

// Watch watches a task
func (tm *TaskManager) Watch(ctx context.Context, name string, dependencies ...string) error {
	task := tm.Get(name)
	if task == nil {
		return ErrTaskNotFound
	}
	err := task.Run(ctx)
	if err != nil {
		log.Error().Err(err).Str("task", name).Msg("failed to run task")
	}
	go task.Watch(ctx, dependencies...)
	tm.Emit(NewTaskEvent(task, TaskStateWatching))
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
	tm.Emit(NewTaskEvent(task, TaskStateStopped))
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
