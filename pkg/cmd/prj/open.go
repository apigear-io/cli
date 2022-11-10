package prj

import (
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewOpenCommand opens the project directory in a ApiGear Studio
func NewOpenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "open project-path",
		Short: "Open a project in studio",
		Long:  `Open the given project in the desktop studio, if installed`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]
			cmd.Printf("open project %s\n", dir)
			return prj.OpenStudio(dir)
		},
	}
	return cmd
}
