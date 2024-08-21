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
		Short: "Lint a template",
		Long:  `Lint a template`,
		Run: func(cmd *cobra.Command, args []string) {
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
