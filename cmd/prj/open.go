package prj

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewOpenCommand returns a new cobra.Command for the "open" command.
func NewOpenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "open project-path",
		Short: "Open a project in studio",
		Long:  `The open command allows you to open a project in studio.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			fmt.Printf("open project %s\n", dir)
			err := prj.OpenStudio(dir)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
