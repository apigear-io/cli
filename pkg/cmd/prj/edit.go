package prj

import (
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewEditCommand opens the project in a configured editor
func NewEditCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit a project in the default editor (vscode)",
		Long:  `Edit a project in the default editor (e.g.Visual Studio Code).`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]
			cmd.Printf("launch vscode with %s\n", dir)
			return prj.OpenEditor(dir)
		},
	}
	return cmd
}
