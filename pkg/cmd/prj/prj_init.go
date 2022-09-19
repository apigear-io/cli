package prj

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewInitCommand returns a new cobra.Command for the "init" command.
func NewInitCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  `Initialize a project with a default project files`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			log.Debugf("init project %s", dir)
			info, err := prj.InitProject(dir)
			if err != nil {
				cmd.Printf("error: %s\n", err)
				return
			}
			cmd.Printf("project initialized at: %s\n", info.Path)

		},
	}
	return cmd
}
