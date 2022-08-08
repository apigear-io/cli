package sol

import (
	"github.com/fsnotify/fsnotify"
)

// WatchSolution watches the solution file and depending files
// for changes and runs the solution when a change is detected
func WatchSolution(solPath string) {
	// run solution and get the dependencies to watch
	deps, err := RunSolution(solPath)
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
					_, err := RunSolution(solPath)
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
