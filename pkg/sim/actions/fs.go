package actions

import (
	"os"

	"github.com/apigear-io/cli/pkg/spec"
	"gopkg.in/yaml.v3"
)

// ReadScenario reads a scenario from file.
func ReadScenario(fn string) (*spec.ScenarioDoc, error) {
	bytes, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	doc := &spec.ScenarioDoc{}
	err = yaml.Unmarshal(bytes, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
