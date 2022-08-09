package tools

import (
	"log"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewJson2YamlCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "json2yaml",
		Aliases: []string{"j2y"},
		Short:   "convert json doc to yaml doc",
		Long:    `convert one or many json documents to yaml documents`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var file = args[0]
			ext := filepath.Ext(file)
			if ext != ".json" {
				log.Fatalf("file %s is not a yaml file", file)
			}
			data, err := os.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			yamlData, err := spec.YamlToJson(data)
			if err != nil {
				log.Fatal(err)
			}
			// replace the extension from yaml(yml) to json
			yamlFile := file[:len(file)-len(ext)] + ".yaml"
			err = os.WriteFile(yamlFile, yamlData, 0644)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
}
