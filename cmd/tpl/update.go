package tpl

import (
	"github.com/apigear-io/lib/tpl"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "update [template]",
		Aliases: []string{"up"},
		Short:   "Update installed templates.",
		Long:    `Update installed templates.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name = args[0]
			err := tpl.UpdateTemplate(name)
			if err != nil {
				cmd.PrintErrln(err)
			}
			cmd.Printf("template %s updated\n", name)
		},
	}
	return cmd
}
