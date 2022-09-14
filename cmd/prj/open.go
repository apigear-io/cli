package prj

import (
	"os"

	"github.com/apigear-io/lib/prj"

	"github.com/spf13/cobra"
)

// NewOpenCommand opens the project directory in a ApiGear Studio
func NewOpenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "open project-path",
		Short: "Open a project in studio",
		Long:  `The open command allows you to open a project in studio.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			cmd.Printf("open project %s\n", dir)
			err := prj.OpenStudio(dir)
			if err != nil {
				cmd.Printf("error: %s\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
