package tpl

import (
	"apigear/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "update [template]",
		Aliases: []string{"up"},
		Short:   "update an installed template",
		Long:    `Update an installed template.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name = args[0]
			cmd.Printf("update template %s\n", name)
			err := tpl.UpdateTemplate(name)
			if err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
	return cmd
}
