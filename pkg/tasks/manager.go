package tasks

import (
	"context"
	"errors"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/sasha-s/go-deadlock"
)

// ErrTaskNotFound is returned when a task is not found
var ErrTaskNotFound = errors.New("task not found")

// TaskManager allows you to create tasks and run them
type TaskManager struct {
	deadlock.RWMutex
	helper.Hook[TaskEvent]
	tasks map[string]*TaskItem
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*TaskItem),
		Hook:  helper.Hook[TaskEvent]{},
	}
}

// Register creates a new task
func (tm *TaskManager) Register(name string, meta map[string]interface{}, tf TaskFunc) *TaskItem {
	if tm.Has(name) {
		err := tm.RmTask(name)
		if err != nil {
			log.Warn().Err(err).Msg("error removing task")
		}
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
	tm.FireHook(NewTaskEvent(task, TaskStateAdded))
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
	tm.FireHook(NewTaskEvent(task, TaskStateRemoved))
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
	tm.FireHook(NewTaskEvent(task, TaskStateRunning))
	err := task.Run(ctx)
	if err != nil {
		log.Error().Err(err).Str("task", name).Msg("failed to run task")
		tm.FireHook(NewTaskEvent(task, TaskStateFailed))
		return err
	}
	tm.FireHook(NewTaskEvent(task, TaskStateFinished))
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
	tm.FireHook(NewTaskEvent(task, TaskStateWatching))
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
	task.CancelWatch()
	task.Cancel()
	tm.FireHook(NewTaskEvent(task, TaskStateStopped))
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
