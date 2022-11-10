package prj

import (
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewRecentCommand returns a new cobra.Command for the "recent" command.
func NewRecentCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "recent",
		Short: "Display recent projects",
		Long:  `Display recently used projects and their locations`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.Println("recent projects:")
			for _, info := range prj.RecentProjectInfos() {
				cmd.Printf("  %s\n", info.Name)
			}
			return nil
		},
	}
	return cmd
}
