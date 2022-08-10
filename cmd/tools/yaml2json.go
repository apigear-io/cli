package tools

import (
	"log"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewYaml2JsonCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "yaml2json",
		Aliases: []string{"y2j"},
		Short:   "convert yaml doc to json doc",
		Long:    `convert one or many yaml documents to json documents`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var input = args[0]
			matches, err := filepath.Glob(input)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range matches {
				ext := filepath.Ext(file)
				if ext != ".yaml" && ext != ".yml" {
					log.Fatalf("file %s is not a yaml file", file)
				}
				data, err := os.ReadFile(file)
				if err != nil {
					log.Fatal(err)
				}
				var jsonData []byte
				jsonData, err = spec.YamlToJson(data)
				if err != nil {
					log.Fatal(err)
				}
				// replace the extension from yaml(yml) to json
				jsonFile := file[:len(file)-len(ext)] + ".json"
				err = os.WriteFile(jsonFile, jsonData, 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	}
	return cmd
}