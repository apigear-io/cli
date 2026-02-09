package tasks

import "fmt"

type TaskState int

const (
	// TaskStateIdle is the state when a task is idle
	TaskStateIdle TaskState = iota
	// TaskStateAdded is the state when a task is added
	TaskStateAdded
	// TaskStateRemoved is the state when a task is removed
	TaskStateRemoved
	// TaskStateWatching is the state when a task is watching
	TaskStateWatching
	// TaskStateRunning is the state when a task is running
	TaskStateRunning
	// TaskStateFinished is the state when a task is finished
	TaskStateFinished
	// TaskStateStopped is the state when a task is stopped
	TaskStateStopped
	// TaskStateFailed is the state when a task is failed
	TaskStateFailed
)

type TaskEvent struct {
	Name  string                 `json:"name"`
	State TaskState              `json:"state"`
	Meta  map[string]interface{} `json:"meta"`
}

// String returns the string representation of the task event
func (e *TaskEvent) String() string {
	return fmt.Sprintf("task %s: %s -> %v", e.Name, e.State, e.Meta)
}

func NewTaskEvent(item *TaskItem, state TaskState) *TaskEvent {
	return &TaskEvent{
		Name:  item.name,
		State: state,
		Meta:  item.meta,
	}
}

// String returns the string representation of the task state
func (s TaskState) String() string {
	switch s {
	case TaskStateIdle:
		return "idle"
	case TaskStateAdded:
		return "added"
	case TaskStateRemoved:
		return "removed"
	case TaskStateWatching:
		return "watching"
	case TaskStateRunning:
		return "running"
	case TaskStateStopped:
		return "stopped"
	case TaskStateFinished:
		return "finished"
	case TaskStateFailed:
		return "failed"
	default:
		return "unknown"
	}
}
