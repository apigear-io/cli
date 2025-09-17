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

func writeDemo(target string, content []byte) (err error) {
	if _, statErr := os.Stat(target); statErr == nil {
		return fmt.Errorf("file %s already exists", target)
	}
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	_, err = f.Write(content)
	return err
}
