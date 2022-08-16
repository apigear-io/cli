package actions

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ReadScenario reads a scenario from file.
func ReadScenario(fn string) (*ScenarioDoc, error) {
	bytes, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	doc := &ScenarioDoc{}
	err = yaml.Unmarshal(bytes, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
