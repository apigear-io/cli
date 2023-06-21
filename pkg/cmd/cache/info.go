package cache

import (
	"github.com/apigear-io/cli/pkg/repos"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [name@version]",
		Short: "display template information from the cache",
		Long:  `display template information from the cached for named templates. I no name is given all templates are listed.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				infos, err := repos.Cache.List()
				if err != nil {
					cmd.PrintErrln(err)
				}
				if len(infos) == 0 {
					cmd.Println("  no results found")
				}
				DisplayTemplateInfos(infos)
			case 1:
				name := args[0]
				if repos.IsRepoID(name) {
					// display info for a specific template version
					info, err := repos.Cache.Info(name)
					if err != nil {
						cmd.PrintErrln(err)
						return
					}
					DisplayTemplateInfo(info)
				} else {
					// display info for all versions of a template
					infos, err := repos.Cache.ListVersions(name)
					if err != nil {
						cmd.PrintErrln(err)
					}
					if len(infos) == 0 {
						cmd.Println("  no results found")
					} else {
						DisplayTemplateInfos(infos)
					}
				}
			}
		},
	}
	return cmd
}
