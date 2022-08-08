package sol

import (
	"github.com/apigear-io/cli/pkg/spec"
)

// RunSolution reads the solution file
// and starts the solution runner
// returns the dependencies of the solution
func RunSolution(file string) ([]string, error) {
	log.Debugf("run solution %s", file)
	doc, err := ReadSolutionDoc(file)
	if err != nil {
		log.Errorf("error reading solution: %s", err)
		return nil, err
	}
	return RunSolutionDocument(doc)
}

func RunSolutionDocument(doc *spec.SolutionDoc) ([]string, error) {
	runner := NewSolutionRunner(doc)
	return runner.Run()
}
