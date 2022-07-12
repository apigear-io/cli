package tpl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {

	// cmd represents the pkgSearch command
	var cmd = &cobra.Command{
		Use:   "search",
		Short: "Search templates by name.",
		Long:  `Search templates by name using wildcards.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pkgSearch called")
		},
	}
	return cmd
}
