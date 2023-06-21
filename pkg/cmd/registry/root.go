package registry

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "registry",
		Aliases: []string{"r"},
		Short:   "template registry management",
		Long:    `Discover and search templates from the remote registry.`,
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
	cmd.AddCommand(NewUpdateCommand())
	return cmd
}
