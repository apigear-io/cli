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
		Short:   "List all installed packages",
		Long:    `List all installed packages. A package can be installed using a git url or a local directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			infos, err := tpl.ListTemplates()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(-1)
			}
			cmd.Println("Installed Template Packages:")
			for _, info := range infos {
				cmd.Printf("  * %s\n", info.Name)
			}
		},
	}
	return cmd
}
