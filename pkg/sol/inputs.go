package sol

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"
)

// parseInputs parses the inputs from the layer.
// A input can be either a file or a directory.
// If the input is a directory, the files in the directory will be parsed.
func parseInputs(s *model.System, inputs []string) error {
	log.Debug().Msgf("parse inputs %v", inputs)
	idlParser := idl.NewParser(s)
	dataParser := model.NewDataParser(s)
	for _, file := range inputs {
		log.Debug().Msgf("parse input %s", file)
		switch filepath.Ext(file) {
		case ".yaml", ".yml", ".json":
			err := dataParser.ParseFile(file)
			if err != nil {
				log.Error().Err(err).Msgf("input file: %s. skip", file)
				return fmt.Errorf("parse %s: %w", file, err)
			}
		case ".idl":
			err := idlParser.ParseFile(file)
			if err != nil {
				log.Error().Err(err).Msgf("input: %s. skip", file)
				return err
			}
		default:
			log.Error().Msgf("unknown type %s. skip", file)
		}
	}
	err := s.ResolveAll()
	if err != nil {
		return fmt.Errorf("error resolving system: %w", err)
	}
	return nil
}

func expandInputs(rootDir string, inputs []string) ([]string, error) {
	result := make([]string, 0)
	for _, input := range inputs {
		input = helper.Join(rootDir, input)
		if helper.IsDir(input) {
			entries, err := os.ReadDir(input)
			if err != nil {
				return nil, err
			}
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				if hasExtension(entry.Name(), []string{"module.yaml", "module.yml", "module.json", ".idl"}) {
					result = append(result, helper.Join(input, entry.Name()))
				}
			}
		} else {
			result = append(result, input)
		}
	}
	return result, nil
}

func checkInputs(inputs []string) error {
	for _, input := range inputs {
		err := checkFile(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkFile(file string) error {
	result, err := spec.CheckFile(file)
	if err != nil {
		log.Warn().Msgf("check document %s: %s", file, err)
		return fmt.Errorf("check document %s: %s", file, err)
	}
	if !result.Valid() {
		log.Warn().Msgf("document %s is invalid", file)
		for _, desc := range result.Errors() {
			log.Warn().Msgf("\t%s", desc)
		}
		return fmt.Errorf("document %s is invalid", file)
	}
	return nil
}
