package prj

import (
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/orchestration/project"

	"github.com/spf13/cobra"
)

func Must(err error) {
	if err != nil {
		logging.Fatal().Msgf("error: %s", err)
	}
}

// NewImportCommand returns a new cobra.Command for the "import" command.
func NewImportCommand() *cobra.Command {
	var target string
	var cmd = &cobra.Command{
		Use:   "import source --target target",
		Short: "import a project",
		Long:  `import a remote project from a repository to the local file system`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			logging.Debug().Msgf("import project %s to %s", source, target)
			info, err := project.ImportProject(source, target)
			if err != nil {
				return err
			}
			cmd.Printf("project %s imported to %s\n", source, info.Path)
			return nil
		},
	}
	cmd.Flags().StringVarP(&target, "target", "t", "", "target directory")
	Must(cmd.MarkFlagRequired("target"))
	return cmd
}
