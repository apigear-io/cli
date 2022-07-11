package prj

import (
	"apigear/pkg/log"
	"apigear/pkg/prj"
	"os"

	"github.com/spf13/cobra"
)

// NewImportCommand returns a new cobra.Command for the "import" command.
func NewImportCommand() *cobra.Command {
	var target string
	var cmd = &cobra.Command{
		Use:   "import source --target target",
		Short: "Import a project from a directory",
		Long:  `The import command allows you to import a project from a directory.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			log.Debug("import project %s to %s", source, target)
			info, err := prj.ImportProject(source, target)
			if err != nil {
				log.Error("error: %s", err)
				os.Exit(1)
			}
			log.Infof("project %s imported to %s", source, info.Path)
		},
	}
	cmd.Flags().StringVarP(&target, "target", "t", "", "target directory")
	cmd.MarkFlagRequired("target")
	return cmd
}
