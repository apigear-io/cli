package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {

	// cmd represents the pkgSearch command
	var cmd = &cobra.Command{
		Use:     "search",
		Short:   "search templates",
		Long:    `search templates by name.`,
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("search registry:")
			pattern := ""
			if len(args) > 0 {
				pattern = args[0]
			}
			result, err := tpl.SearchRegistry(pattern)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if len(result) == 0 {
				cmd.Println("  no results found")
			} else {
				for _, info := range result {
					cmd.Printf("  * %s\t%s\n", info.Name, info.Git)
				}
			}
		},
	}
	return cmd
}
