package tpl

import (
	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/spf13/cobra"
)

func NewLintCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "lint",
		Short: "Lint a template directory",
		Long:  `Lint a template directory, it will fail if the templates can not be compiled`,
		Run: func(cmd *cobra.Command, args []string) {
			// trying to create a generator, it will fail
			// if the templates in the dir are not valid
			_, err := gen.New(gen.Options{
				TemplatesDir: dir,
				System:       model.NewSystem("test"),
				Features:     []string{"all"},
				Force:        true,
			})
			if err != nil {
				cmd.Printf("template dir '%s' is not valid: %s\n", dir, err)
			} else {
				cmd.Printf("template dir '%s' is valid\n", dir)
			}
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	cmd.MarkFlagRequired("dir")
	return cmd
}
