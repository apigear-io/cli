package x

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

func idl2yaml(input string) error {
	matches, err := filepath.Glob(input)
	if err != nil {
		return err
	}
	for _, file := range matches {
		log.Debug().Msgf("Converting IDL file: %s", file)
		ext := filepath.Ext(file)
		if ext != ".idl" {
			return fmt.Errorf("%s is not an IDL file", file)
		}
		sys := model.NewSystem("NO_NAME")
		log.Debug().Msgf("Parsing IDL file: %s", file)
		parser := idl.NewParser(sys)
		err = parser.ParseFile(file)
		if err != nil {
			return fmt.Errorf("parse IDL file: %w", err)
		}
		log.Debug().Msgf("Validating system after parsing IDL file: %s", file)
		err = sys.Validate()
		if err != nil {
			return fmt.Errorf("validate system: %w", err)
		}
		for _, module := range sys.Modules {
			log.Debug().Msgf("Converting module %s to YAML", module.Name)
			data, err := yaml.Marshal(module)
			if err != nil {
				return fmt.Errorf("marshal module to YAML: %w", err)
			}
			newFile := strings.TrimSuffix(file, ext) + ".yaml"
			log.Debug().Msgf("Writing YAML file: %s", newFile)
			err = os.WriteFile(newFile, data, 0644)
			if err != nil {
				return fmt.Errorf("write YAML file: %w", err)
			}
			fmt.Printf("Converted %s to %s\n", file, newFile)
		}
	}
	return nil
}

func NewIdl2YamlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "idl2yaml [file]",
		Short: "Convert IDL file to YAML",
		Long:  `Convert an IDL file to a YAML representation.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no file specified")
			}
			file := args[0]
			return idl2yaml(file)
		},
	}
	return cmd
}
