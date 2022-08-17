package gen

import (
	"github.com/apigear-io/cli/pkg/spec"
	"gopkg.in/yaml.v3"
	"os"
)

// ReadRulesDoc reads rules from a file.
func ReadRulesDoc(filename string) (spec.RulesDoc, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return spec.RulesDoc{}, err
	}
	var file spec.RulesDoc
	err = yaml.Unmarshal(content, &file)
	if err != nil {
		return spec.RulesDoc{}, err
	}
	return file, nil
}
