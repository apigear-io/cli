package prj

import (
	"github.com/spf13/cobra"
)

// NewShareCommand returns a new cobra.Command for the "share" command.
func NewShareCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "share",
		Short: "share a project with your team (tbd)",
		Long:  `share a project and all files with your team to work together (tbd)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			cmd.Printf("share project %s\n", dir)
		},
	}
	return cmd
}
