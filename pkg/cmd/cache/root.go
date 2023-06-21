package cache

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "cache",
		Aliases: []string{"c"},
		Short:   "templayte cache management",
		Long:    `manage sdk templates cache. A template is a git repository that contains a sdk template.`,
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
	cmd.AddCommand(NewImportCommand())
	cmd.AddCommand(NewCleanCommand())
	return cmd
}
