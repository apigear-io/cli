package prj

import (
	"fmt"
	"os"
)

type DemoType string

const (
	DemoModule   DemoType = "module"
	DemoSolution DemoType = "solution"
	DemoScenario DemoType = "scenario"
)

func writeDemo(target string, content []byte) error {
	if _, err := os.Stat(target); err == nil {
		return fmt.Errorf("file %s already exists", target)
	}
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}
