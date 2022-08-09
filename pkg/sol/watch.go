package sol

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/fsnotify/fsnotify"
)

// WatchSolutionFile watches the solution file and depending files
// for changes and runs the solution when a change is detected
func WatchSolutionFile(solPath string) {
	// run solution and get the dependencies to watch
	deps, err := RunSolutionFile(solPath)
	if err != nil {
		log.Warnf("Error running solution: %s", err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("file %s modified", event.Name)
					_, err := RunSolutionFile(solPath)
					if err != nil {
						log.Warnf("Error running solution: %s", err)
					}
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()
	err = watcher.Add(solPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, dep := range deps {
		err = watcher.Add(dep)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}

// WatchSolutionDocument watches the solution document and depending files
// for changes and runs the solution when a change is detected
func WatchSolutionDocument(doc *spec.SolutionDoc) {
	// run solution and get the dependencies to watch
	deps, err := RunSolutionDocument(doc)
	if err != nil {
		log.Warnf("Error running solution: %s", err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("file %s modified", event.Name)
					_, err := RunSolutionDocument(doc)
					if err != nil {
						log.Warnf("Error running solution: %s", err)
					}
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()
	for _, dep := range deps {
		// check if dep is a file and if add to watcher
		info, err := os.Stat(dep)
		if err != nil {
			log.Warnf("Error getting file info for %s: %s", dep, err)
			continue
		}
		if info.Mode().IsRegular() {
			log.Debugf("Adding file %s to watcher", dep)
			err = watcher.Add(dep)
			if err != nil {
				log.Fatal(err)
			}
		} else if info.IsDir() {
			err := filepath.Walk(dep, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Warnf("Error walking path %s: %s", path, err)
					return nil
				}
				if info.IsDir() {
					log.Debugf("Adding dir %s to watcher", path)
					err := watcher.Add(path)
					if err != nil {
						log.Fatal(err)
					}
				}
				return nil
			})
			if err != nil {
				log.Warnf("Error walking path %s: %s", dep, err)
			}
		}
	}
	<-done
}
