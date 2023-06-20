package tasks

import (
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/fsnotify/fsnotify"
)

// TaskFunc is the function type of the task to run
type TaskFunc func(ctx context.Context) error

// TaskItem is the task item stored in the TaskManager
type TaskItem struct {
	sync.RWMutex
	name     string
	meta     map[string]interface{}
	taskFunc TaskFunc
	cancel   context.CancelFunc
}

// NewTaskItem creates a new task item
func NewTaskItem(name string, meta map[string]interface{}, tf TaskFunc) *TaskItem {
	return &TaskItem{
		name:     name,
		meta:     meta,
		taskFunc: tf,
	}
}

// Run runs the task once
// TODO: add error handling
func (t *TaskItem) Run(ctx context.Context) error {
	log.Debug().Msgf("run task: %s", t.name)
	if t.cancel != nil {
		// cancel the previous task
		t.cancel()
	}
	ctx, t.cancel = context.WithCancel(ctx)
	err := t.taskFunc(ctx)
	// handle the error
	if err != nil {
		t.UpdateMeta(map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	return nil
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
		log.Info().Msgf("watching file %s", dep)
		err := watcher.Add(dep)
		if err != nil {
			log.Debug().Msgf("error watching file %s: %s", dep, err)
		}
		// check if the dependency is a directory
		if helper.IsDir(dep) {
			filepath.WalkDir(dep, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					log.Error().Err(err).Msgf("error walking directory %s", dep)
					return err
				}
				if d.IsDir() {
					watcher.Add(path)
				}
				return nil
			})
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Debug().Msgf("modified file: %s", event.Name)
				err := t.Run(ctx)
				if err != nil {
					log.Error().Err(err).Msg("failed to run task")
				}
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

// UpdateMeta updates the meta data of the task
func (t *TaskItem) UpdateMeta(meta map[string]interface{}) {
	t.Lock()
	defer t.Unlock()
	for k, v := range meta {
		t.meta[k] = v
	}
}
