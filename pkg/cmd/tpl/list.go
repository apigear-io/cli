package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	// cmd represents the pkgList command
	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list installed templates",
		Long:    `list installed templates. A template can be installed the install command.`,
		Run: func(cmd *cobra.Command, _ []string) {
			infos, err := tpl.ListTemplates()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(-1)
			}
			cmd.Println("list installed templates:")
			for _, info := range infos {
				cmd.Printf("  * %s\n", info.Name)
			}
		},
	}
	return cmd
}
