package tpl

import (
	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {

	// cmd represents the pkgSearch command
	var cmd = &cobra.Command{
		Use:   "search",
		Short: "Search templates by name.",
		Long:  `Search templates by name using wildcards.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Please visit https://apigear.io/templates/ for a list of templates.")
		},
	}
	return cmd
}
