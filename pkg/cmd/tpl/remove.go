package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "remove [name]",
		Aliases: []string{"rm"},
		Short:   "remove installed template",
		Long:    `remove installed template by name.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name = args[0]
			err := tpl.RemoveTemplate(name)
			if err != nil {
				cmd.PrintErrln(err)
			} else {
				cmd.Printf("template %s removed \n", name)
			}
		},
	}
	return cmd
}
