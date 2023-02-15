package tasks

import (
	"context"
	"os"
	"sync"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/fsnotify/fsnotify"
)

// TaskFunc is the function type of the task to run
type TaskFunc func(ctx context.Context) error

// TaskItem is the task item stored in the TaskManager
type TaskItem struct {
	sync.RWMutex
	name     string
	taskFunc TaskFunc
	cancel   context.CancelFunc
}

// NewTaskItem creates a new task item
func NewTaskItem(name string, tf TaskFunc) *TaskItem {
	return &TaskItem{
		name:     name,
		taskFunc: tf,
	}
}

// Run runs the task once
func (t *TaskItem) Run(ctx context.Context) {
	log.Debug().Msgf("run task: %s", t.name)
	if t.cancel != nil {
		// cancel the previous task
		t.cancel()
	}
	ctx, t.cancel = context.WithCancel(ctx)
	err := t.taskFunc(ctx)
	if err != nil {
		log.Error().Err(err).Str("task", t.name).Msg("failed to run task")
	}
	// clear the cancel function
	t.cancel = nil
}

// Watch watches all the dependencies of the task and runs the task
// it uses fsnotify to watch the files
func (t *TaskItem) Watch(ctx context.Context, dependencies ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error().Msgf("error creating watcher: %s", err)
		watcher.Close()
		return
	}
	defer watcher.Close()

	for _, dep := range dependencies {
		// check if the file exists
		if _, err := os.Stat(dep); os.IsNotExist(err) {
			log.Debug().Msgf("file %s does not exist", dep)
			continue
		}
		err := watcher.Add(dep)
		if err != nil {
			log.Debug().Msgf("error watching file %s: %s", dep, err)
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Debug().Msgf("modified file: %s", event.Name)
				t.Run(ctx)
			}
		case err := <-watcher.Errors:
			log.Error().Msgf("error watching file: %s", err)
		}
	}
}

// Cancel cancels the task
func (t *TaskItem) Cancel() {
	if t.cancel == nil {
		return
	}
	t.cancel()
}
