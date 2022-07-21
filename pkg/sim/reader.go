package sim

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func ReadScenario(file string) (*ScenarioDoc, error) {
	switch filepath.Ext(file) {
	case ".json":
		return ReadJsonScenario(file)
	case ".yaml":
		return ReadYamlScenario(file)
	default:
		return nil, fmt.Errorf("unsupported scenario file extension %s", file)
	}
}

func ReadJsonScenario(file string) (*ScenarioDoc, error) {
	scenario := &ScenarioDoc{}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Errorf("failed to read scenario file %s: %v", file, err)
		return nil, err
	}
	err = json.Unmarshal(data, scenario)
	if err != nil {
		log.Errorf("failed to unmarshal scenario file %s: %v", file, err)
		return nil, err
	}
	return scenario, err
}

func ReadYamlScenario(file string) (*ScenarioDoc, error) {
	scenario := &ScenarioDoc{}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Errorf("failed to read scenario file %s: %v", file, err)
		return nil, err
	}
	err = yaml.Unmarshal(data, scenario)
	if err != nil {
		log.Errorf("failed to unmarshal scenario file %s: %v", file, err)
		return nil, err
	}
	return scenario, err
}
