package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "update [template]",
		Aliases: []string{"up"},
		Short:   "List and update installed templates.",
		Long:    `List and update installed templates.`,
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
