package cache

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "install [git-url] [version]",
		Short:   "install template into cache",
		Long:    `install template from git-url and version into cache`,
		Aliases: []string{"i"},
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			version := args[1]
			cmd.Printf("installing template from %s\n", url)
			fqn, err := repos.Cache.Install(url, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Printf("installed template %s\n", fqn)
		},
	}
	return cmd
}
