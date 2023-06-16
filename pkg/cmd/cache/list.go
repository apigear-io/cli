package cache

import (
	"os"

	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	// cmd represents the pkgList command
	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list templates",
		Long:    `list templates. A template can be installed the install command.`,
		Run: func(cmd *cobra.Command, _ []string) {
			infos, err := repos.Cache.List()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(-1)
			}
			cmd.Println("list of templates from registry and cache")
			displayRepoInfos(infos)
		},
	}
	return cmd
}
