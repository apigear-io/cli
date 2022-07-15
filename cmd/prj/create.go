package prj

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var project string
	var cmd = &cobra.Command{
		Use:   "create doc-type doc-name",
		Short: "Create a new document inside current project",
		Long:  `The create command allows you to create a new document inside current project.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			docType := args[0]
			name := args[1]
			docName := prj.MakeDocumentName(docType, name)
			target := filepath.Join(project, "apigear", docName)
			err := prj.CreateProjectDocument(docType, target)
			if err != nil {
				log.Errorf("error: %s\n", err)
				return
			}
			cmd.Printf("document %s created\n", target)
		},
	}
	cmd.Flags().StringVarP(&project, "project", "p", ".", "project directory")
	return cmd
}
