package sol

import (
	"github.com/fsnotify/fsnotify"
)

// WatchSolution watches the solution file and depending files
// for changes and runs the solution when a change is detected
func WatchSolution(file string) {
	err := RunSolution(file)
	if err != nil {
		log.Fatal(err)
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
					err := RunSolution(file)
					if err != nil {
						log.Errorf("failed to run solution: %s", err)
					}
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()
	err = watcher.Add(file)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
