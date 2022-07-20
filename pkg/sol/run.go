package sol

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
)

// RunSolution reads the solution file
// and starts the solution runner
func RunSolution(file string) error {
	log.Infof("run solution %s", file)
	doc, err := ReadSolutionDoc(file)
	if err != nil {
		log.Errorf("error reading solution: %s", err)
		return err
	}
	rootDir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return err
	}
	proc := NewSolutionRunner(rootDir, doc)
	proc.Run()
	return nil
}
