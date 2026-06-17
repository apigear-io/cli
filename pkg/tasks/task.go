package tasks

import (
	"context"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/fsnotify/fsnotify"
)

// TaskFunc is the function type of the task to run
type TaskFunc func(ctx context.Context) error

// Task represents a simple runnable task with optional file watching
type Task struct {
	cancel      context.CancelFunc
	watchCancel context.CancelFunc
}

// NewTask creates a new task
func NewTask() *Task {
	return &Task{}
}

// Run runs the task function once
func (t *Task) Run(ctx context.Context, fn TaskFunc) error {
	if t.cancel != nil {
		t.cancel()
	}
	ctx, t.cancel = context.WithCancel(ctx)
	return fn(ctx)
}

// Watch watches files and re-runs the task function when they change
func (t *Task) Watch(ctx context.Context, fn TaskFunc, dependencies ...string) error {
	// Run once initially
	if err := t.Run(ctx, fn); err != nil {
		log.Error().Err(err).Msg("initial task run failed")
	}

	if t.watchCancel != nil {
		t.watchCancel()
	}
	ctx, t.watchCancel = context.WithCancel(ctx)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error().Err(err).Msg("error creating watcher")
		return err
	}
	defer func() {
		if err := watcher.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close watcher")
		}
	}()

	for _, dep := range dependencies {
		// check if the file exists
		if _, err := os.Stat(dep); os.IsNotExist(err) {
			log.Debug().Str("file", dep).Msg("file does not exist")
			continue
		}
		log.Info().Str("file", dep).Msg("watching file")
		err := watcher.Add(dep)
		if err != nil {
			log.Debug().Err(err).Str("file", dep).Msg("error watching file")
		}
		// check if the dependency is a directory
		if helper.IsDir(dep) {
			err = filepath.WalkDir(dep, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					log.Error().Err(err).Str("dir", dep).Msg("error walking directory")
					return err
				}
				if d.IsDir() {
					err = watcher.Add(path)
					if err != nil {
						log.Warn().Err(err).Str("path", path).Msg("error watching directory")
					}
				}
				return nil
			})
			if err != nil {
				log.Warn().Err(err).Str("dir", dep).Msg("error walking directory")
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Debug().Str("file", event.Name).Msg("file modified")
				if err := t.Run(ctx, fn); err != nil {
					log.Error().Err(err).Msg("task run failed")
				}
			}
		case err := <-watcher.Errors:
			log.Error().Err(err).Msg("watcher error")
		}
	}
}

// Cancel cancels the running task
func (t *Task) Cancel() {
	if t.cancel != nil {
		t.cancel()
	}
}

// CancelWatch cancels the watch operation
func (t *Task) CancelWatch() {
	if t.watchCancel != nil {
		t.watchCancel()
	}
}
