package gen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"
	"gopkg.in/yaml.v3"
)

// ReadRulesDoc reads rules from a file.
func ReadRulesDoc(filename string) (spec.RulesDoc, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return spec.RulesDoc{}, err
	}
	err = CheckRulesJson(filename, bytes)
	if err != nil {
		return spec.RulesDoc{}, err
	}
	var file spec.RulesDoc
	err = yaml.Unmarshal(bytes, &file)
	if err != nil {
		return spec.RulesDoc{}, err
	}
	return file, nil
}

func CheckRulesJson(file string, bytes []byte) error {
	var err error
	if filepath.Ext(file) == ".yaml" || filepath.Ext(file) == ".yml" {
		bytes, err = spec.YamlToJson(bytes)
		if err != nil {
			return err
		}
	}
	result, err := spec.CheckJson(spec.DocumentTypeRules, bytes)
	if err != nil {
		log.Warn().Msgf("check document %s: %s", file, err)
		return err
	}
	if !result.Valid() {
		log.Warn().Msgf("document %s is invalid", file)
		for _, desc := range result.Errors() {
			log.Warn().Msg(desc.String())
			err = fmt.Errorf("%s", desc.String())
		}
	}
	return err
}
