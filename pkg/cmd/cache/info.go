package cache

import (
	"github.com/apigear-io/cli/pkg/repos"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [name] [version]",
		Short: "display template information from the cache",
		Long:  `display template information from the cached for named templates. I no name is given all templates are listed.`,
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
				infos, err := repos.Cache.ListVersions(name)
				if err != nil {
					cmd.PrintErrln(err)
				}
				if len(infos) == 0 {
					cmd.Println("  no results found")
				} else {
					DisplayTemplateInfos(infos)
				}
			case 2:
				name := args[0]
				version := args[1]
				info, err := repos.Cache.Info(name, version)
				if err != nil {
					cmd.PrintErrln(err)
					return
				}
				DisplayTemplateInfo(info)
			default:
				cmd.PrintErrln("invalid number of arguments")
			}
		},
	}
	return cmd
}
