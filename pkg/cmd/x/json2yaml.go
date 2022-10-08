package x

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func Json2Yaml(input string) error {
	matches, err := filepath.Glob(input)
	if err != nil {
		return err
	}
	for _, file := range matches {
		ext := filepath.Ext(file)
		if ext != ".json" {
			return fmt.Errorf("%s is not a json file", file)
		}
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		yamlData, err := spec.JsonToYaml(data)
		if err != nil {
			return err
		}
		// replace the extension from yaml(yml) to json
		yamlFile := file[:len(file)-len(ext)] + ".yaml"
		err = os.WriteFile(yamlFile, yamlData, 0644)
		if err != nil {
			return err
		}
	}
	return nil

}

func NewJson2YamlCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "json2yaml",
		Aliases: []string{"j2y"},
		Short:   "convert json doc to yaml doc",
		Long:    `convert one or many json documents to yaml documents`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := Json2Yaml(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("convert json to yaml")
			}
		},
	}
	return cmd
}
