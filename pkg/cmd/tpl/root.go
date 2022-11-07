package tpl

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "template",
		Aliases: []string{"t", "tpl"},
		Short:   "manage sdk templates",
		Long:    `sdk templates are git repositories that contain a sdk template.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Usage()
			if err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
	cmd.AddCommand(NewSearchCommand())
	cmd.AddCommand(NewInstallCommand())
	cmd.AddCommand(NewInfoCommand())
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewRemoveCommand())
	cmd.AddCommand(NewUpdateCommand())
	cmd.AddCommand(NewUpgradeCommand())
	cmd.AddCommand(NewImportCommand())
	return cmd
}
