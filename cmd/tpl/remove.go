package tpl

import (
	"github.com/apigear-io/lib/tpl"

	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "remove [template]",
		Aliases: []string{"rm"},
		Short:   "Remove an installed template",
		Long:    `Remove an installed template.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name = args[0]
			cmd.Printf("rm template %s\n", name)
			err := tpl.RemoveTemplate(name)
			if err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
	return cmd
}
