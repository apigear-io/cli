package prj

import (
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var prjDir string
	var cmd = &cobra.Command{
		Use:   "create doc-type doc-name",
		Short: "Create a new document inside current project",
		Long:  `Create a new document inside current project from a template.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			docType := args[0]
			name := args[1]
			target, err := prj.CreateProjectDocument(prjDir, docType, name)
			if err != nil {
				return err
			}
			cmd.Printf("document %s created\n", target)
			return nil
		},
	}
	cmd.Flags().StringVarP(&prjDir, "project", "p", ".", "project directory")
	return cmd
}
