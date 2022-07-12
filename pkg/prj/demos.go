package prj

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/vfs"
)

func writeDemoModule(path string) error {
	log.Infof("Write Demo Module %s", path)
	// return if path exists
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file %s already exists", path)
	}
	return os.WriteFile(path, []byte(vfs.DemoModule), 0644)
}

func writeDemoSolution(path string) error {
	log.Infof("Write Demo Solution %s", path)
	// check if path exists
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file %s already exists", path)
	}
	log.Infof("Write Demo Solution %s", path)
	return os.WriteFile(path, []byte(vfs.DemoSolution), 0644)
}
