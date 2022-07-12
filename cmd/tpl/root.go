package tpl

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "template",
		Aliases: []string{"t", "tpl"},
		Short:   "Manage code generation templates for SDK creation.",
		Long:    `SDK templates can be installed from git repositories and used to generate code. The templates are stored in a local folder.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(NewSearchCommand())
	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewOpenCommand())
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewRemoveCommand())
	cmd.AddCommand(NewUpdateCommand())
	return cmd
}
