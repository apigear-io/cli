package cache

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "remove [name@version]",
		Aliases: []string{"rm"},
		Short:   "remove installed template",
		Long:    `remove installed template by name.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fqn := args[0]
			err := repos.Cache.Remove(fqn)
			if err != nil {
				cmd.PrintErrln(err)
			} else {
				cmd.Printf("template %s removed \n", fqn)
			}
		},
	}
	return cmd
}
