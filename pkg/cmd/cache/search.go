package cache

import (
	"github.com/apigear-io/cli/pkg/repos"
	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {

	// cmd represents the pkgSearch command
	var cmd = &cobra.Command{
		Use:     "search",
		Short:   "search templates from cache",
		Long:    `search templates by name from cache.`,
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("search results ...")
			pattern := ""
			if len(args) > 0 {
				pattern = args[0]
			}
			infos, err := repos.Cache.Search(pattern)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if len(infos) == 0 {
				cmd.Println("  no results found")
			} else {
				displayRepoInfos(infos)
			}
		},
	}
	return cmd
}
