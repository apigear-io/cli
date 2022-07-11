package prj

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewShareCommand returns a new cobra.Command for the "share" command.
func NewShareCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "share",
		Short: "Share a project",
		Long:  `The share command allows you to share a project.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			fmt.Printf("share project %s\n", dir)
		},
	}
	return cmd
}
