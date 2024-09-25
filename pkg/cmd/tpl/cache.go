package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewCacheCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "cache",
		Short: "list templates in the local cache",
		Run: func(cmd *cobra.Command, _ []string) {
			infos, err := repos.Cache.List()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(-1)
			}
			cmd.Println("list of templates from the local cache")
			cmd.Println()
			DisplayTemplateInfos(infos)
		},
	}
	return cmd
}
