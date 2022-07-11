package prj

import (
	"apigear/pkg/prj"
	"fmt"

	"github.com/spf13/cobra"
)

// NewInitCommand returns a new cobra.Command for the "init" command.
func NewInitCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  `The init command allows you to initialize a new project.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			fmt.Printf("init project %s\n", dir)
			info, err := prj.InitProject(dir)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
			fmt.Printf("created project at: %s\n", info.Path)

		},
	}
	return cmd
}
