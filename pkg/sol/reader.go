package sol

import (
	"io/ioutil"
	"objectapi/pkg/spec"

	"gopkg.in/yaml.v2"
)

func ReadSolutionDoc(file string) (spec.SolutionDoc, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return spec.SolutionDoc{}, err
	}
	var sol spec.SolutionDoc
	err = yaml.Unmarshal(data, &sol)
	if err != nil {
		return spec.SolutionDoc{}, err
	}
	return sol, nil
}
