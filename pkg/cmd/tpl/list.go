package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/pterm/pterm"

	"github.com/spf13/cobra"
)

func displayRepoInfos(infos []*git.RepoInfo) {
	cells := make([][]string, len(infos)+1)

	cells[0] = []string{"name", "local", "latest", "url", "state"}
	for i, info := range infos {
		state := "unknown"
		if info.InCache {
			state = "cached"
		} else if info.InRegistry {
			state = "registry"
		}

		cells[i+1] = []string{
			info.Name,
			info.Tag,
			info.Latest.Name,
			info.Git,
			state,
		}
	}
	err := pterm.DefaultTable.WithHasHeader().WithData(cells).Render()
	if err != nil {
		pterm.Error.Println(err)
	}
}

func NewListCommand() *cobra.Command {
	// cmd represents the pkgList command
	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "list templates",
		Long:    `list templates. A template can be installed the install command.`,
		Run: func(cmd *cobra.Command, _ []string) {
			infos, err := tpl.ListTemplates()
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
