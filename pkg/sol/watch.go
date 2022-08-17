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
	t.watchForever()
	return t.done, nil
}

func (t *task) stopWatch() {
	t.done <- true
	if t.watcher != nil {
		t.watcher.Close()
	}
}

func (t *task) watchDeps() error {
	if t.watcher == nil {
		return fmt.Errorf("watcher is not initialized")
	}
	for _, dep := range t.deps {
		// check if dep is a file and if add to watcher
		info, err := os.Stat(dep)
		if err != nil {
			log.Warnf("file info for %s: %s", dep, err)
			continue
		}
		if info.Mode().IsRegular() {
			log.Debugf("add file %s to watcher", dep)
			err = t.watcher.Add(dep)
			if err != nil {
				log.Fatal(err)
			}
		} else if info.IsDir() {
			err := filepath.Walk(dep, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Warnf("walk path %s: %s", path, err)
					return nil
				}
				if info.IsDir() {
					log.Debugf("add dir %s to watcher", path)
					err := t.watcher.Add(path)
					if err != nil {
						log.Fatal(err)
					}
				}
				return nil
			})
			if err != nil {
				log.Warnf("walk path %s: %s", dep, err)
			}
		}
	}
	return nil
}

func (t *task) watchForever() {
	if t.watcher == nil {
		log.Warnf("watcher is not initialized")
	}
	for {
		select {
		case event := <-t.watcher.Events:
			log.Debugf("event: %s", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Debugf("modified file: %s", event.Name)
				err := t.run()
				if err != nil {
					log.Warnf("error running watched task: %s", err)
				}
			}
		case err := <-t.watcher.Errors:
			log.Errorf("watch error: %s", err)
		case <-t.done:
			return
		}
	}
}
