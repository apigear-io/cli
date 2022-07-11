package prj

import (
	"apigear/pkg/prj"
	"fmt"

	"github.com/spf13/cobra"
)

// NewRecentCommand returns a new cobra.Command for the "recent" command.
func NewRecentCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "recent",
		Short: "Show recent projects",
		Long:  `The recent command allows you to show recent projects.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("recent projects:")
			for _, info := range prj.RecentProjectInfos() {
				fmt.Printf("  %s\n", info.Name)
			}
		},
	}
	return cmd
}
