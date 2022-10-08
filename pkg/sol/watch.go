package sol

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func (t *task) startWatch() (chan<- bool, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	t.watcher = watcher
	err = t.watchDeps()
	if err != nil {
		return nil, err
	}
	go t.watchForever()
	return t.done, nil
}

func (t *task) stopWatch() {
	t.done <- true
}

func (t *task) watchDeps() error {
	if t.watcher == nil {
		return fmt.Errorf("watcher is not initialized")
	}
	for _, dep := range t.deps {
		// check if dep is a file and if add to watcher
		info, err := os.Stat(dep)
		if err != nil {
			log.Warn().Msgf("file info for %s: %s", dep, err)
			continue
		}
		if info.Mode().IsRegular() {
			log.Debug().Msgf("add file %s to watcher", dep)
			err = t.watcher.Add(dep)
			if err != nil {
				log.Error().Err(err).Msg("failed to add file to watcher")
			}
		} else if info.IsDir() {
			err := filepath.Walk(dep, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Warn().Msgf("walk path %s: %s", path, err)
					return nil
				}
				if info.IsDir() {
					log.Debug().Msgf("add dir %s to watcher", path)
					err := t.watcher.Add(path)
					if err != nil {
						log.Error().Err(err).Msg("failed to add dir to watcher")
					}
				}
				return nil
			})
			if err != nil {
				log.Warn().Msgf("walk path %s: %s", dep, err)
			}
		}
	}
	return nil
}

func (t *task) watchForever() {
	if t.watcher == nil {
		log.Warn().Msgf("watcher is not initialized")
	}
	for {
		select {
		case event, ok := <-t.watcher.Events:
			if !ok {
				return
			}
			log.Debug().Msgf("event: %s", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Debug().Msgf("modified file: %s", event.Name)
				err := t.run()
				if err != nil {
					log.Warn().Msgf("error running watched task: %s", err)
				}
			}
		case err, ok := <-t.watcher.Errors:
			if !ok {
				log.Debug().Msgf("watcher error: %s", err)
				return
			}
			log.Warn().Msgf("watch error: %s", err)
		case <-t.done:
			t.watcher.Close()
			return
			// do nothing
		}
	}
}
