package sol

import (
	"fmt"
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
	for _, file := range inputs {
		log.Debug().Msgf("parse input %s", file)
		switch filepath.Ext(file) {
		case ".yaml", ".yml", ".json":
			p := model.NewDataParser(s)
			err := p.ParseFile(file)
			if err != nil {
				log.Error().Err(err).Msgf("input file: %s. skip", file)
				return fmt.Errorf("parse %s: %w", file, err)
			}
		case ".idl":
			p := idl.NewParser(s)
			err := p.ParseFile(file)
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

func checkInputs(inputs []string) error {
	log.Debug().Msgf("check inputs %v", inputs)
	for _, input := range inputs {
		switch helper.Ext(input) {
		case ".yaml", ".yml", ".json":
			err := checkFile(input)
			if err != nil {
				return err
			}
		case ".idl":
			err := spec.CheckIdlFile(input)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown type %s", input)
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
