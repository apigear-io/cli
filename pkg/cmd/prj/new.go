package prj

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewProjectCommand returns a new cobra.Command for the "init" command.
func NewProjectCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "create new project",
		Long:  `create new project with default project files`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug().Msgf("create project in %s", dir)
			info, err := prj.InitProject(dir)
			if err != nil {
				return err
			}
			cmd.Printf("project created at: %s\n", info.Path)
			return nil
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "project directory to create")
	cmd.MarkFlagRequired("dir")
	return cmd
}
