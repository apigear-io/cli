package x

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "embed" // for embedding the template

	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/spf13/cobra"
)

//go:embed module.idl.tpl
var moduleIdlTemplate string

func Yaml2Idl(input string) error {
	matches, err := filepath.Glob(input)
	if err != nil {
		return err
	}
	for _, file := range matches {
		ext := filepath.Ext(file)
		if ext != ".yaml" && ext != ".yml" {
			return fmt.Errorf("%s is not a yaml file", file)
		}
		system := model.NewSystem("NO_NAME")
		p := model.NewDataParser(system)
		err = p.ParseFile(file)
		if err != nil {
			return err
		}
		err := system.Validate()
		if err != nil {
			return fmt.Errorf("validate system: %w", err)
		}
		if len(system.Modules) == 0 {
			return fmt.Errorf("no modules found in %s", file)
		}
		if len(system.Modules) > 1 {
			return fmt.Errorf("multiple modules found in %s, only one module is supported", file)
		}
		module := system.Modules[0]
		ctx := model.ModuleScope{
			System: system,
			Module: module,
		}
		out, err := gen.RenderString(moduleIdlTemplate, ctx)
		if err != nil {
			return fmt.Errorf("render module idl: %w", err)
		}
		newFile := strings.TrimSuffix(file, ext) + ".idl"
		err = os.WriteFile(newFile, []byte(out), 0644)
		if err != nil {
			return fmt.Errorf("write idl file: %w", err)
		}
	}
	return nil
}

func NewYaml2IdlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "yaml2idl <file>",
		Aliases: []string{"y2i"},
		Short:   "convert yaml api doc to idl doc",
		Long:    `convert one or many yaml api documents to idl documents`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := Yaml2Idl(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("convert yaml to idl")
			}
		},
	}
	return cmd
}
