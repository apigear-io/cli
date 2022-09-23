package x

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func Yaml2Json(input string) error {
	matches, err := filepath.Glob(input)
	if err != nil {
		return err
	}
	for _, file := range matches {
		ext := filepath.Ext(file)
		if ext != ".yaml" && ext != ".yml" {
			return fmt.Errorf("%s is not a yaml file", file)
		}
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		var jsonData []byte
		jsonData, err = spec.YamlToJson(data)
		if err != nil {
			return err
		}
		// replace the extension from yaml(yml) to json
		jsonFile := file[:len(file)-len(ext)] + ".json"
		err = os.WriteFile(jsonFile, jsonData, 0644)
		if err != nil {
			return err
		}
		log.Info().Msgf("converted %s to %s", file, jsonFile)
	}
	return nil
}

func NewYaml2JsonCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "yaml2json <file>",
		Aliases: []string{"y2j"},
		Short:   "convert yaml doc to json doc",
		Long:    `convert one or many yaml documents to json documents`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.OnReport(func(l *log.ReportEvent) {
				fmt.Println(l.Message)
			})
			err := Yaml2Json(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("failed to convert yaml to json")
			}
		},
	}
	return cmd
}
