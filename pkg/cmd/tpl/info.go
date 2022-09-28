package tpl

import (
	"strconv"

	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func DisplayTemplateInfos(infos []tpl.TemplateInfo) {
	cells := make([][]string, len(infos)+1)
	cells[0] = []string{"name", "url", "cached", "registry"}
	for i, info := range infos {
		cells[i+1] = []string{
			info.Name,
			info.Git,
			strconv.FormatBool(info.InCache),
			strconv.FormatBool(info.InRegistry),
		}
	}

	pterm.DefaultTable.WithHasHeader().WithData(cells).Render()
}

func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [name]",
		Short: "display template information",
		Long:  `display template information for named templates. I no name is given all templates are listed.`,
		Run: func(cmd *cobra.Command, args []string) {
			infos := []tpl.TemplateInfo{}
			if len(args) == 0 {
				list, err := tpl.ListTemplates()
				if err != nil {
					cmd.PrintErrln(err)
				}
				infos = list
			} else {
				for _, name := range args {
					info, err := tpl.GetLocalTemplateInfo(name)
					if err != nil {
						cmd.PrintErrln(err)
						return
					}
					infos = append(infos, info)
				}
			}
			DisplayTemplateInfos(infos)
		},
	}
	return cmd
}
