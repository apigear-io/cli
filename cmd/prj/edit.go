package prj

import (
	"os"

	"github.com/apigear-io/lib/prj"

	"github.com/spf13/cobra"
)

// NewEditCommand opens the project in a configured editor
func NewEditCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit a project in the default editor (vscode)",
		Long:  `The edit command allows you to edit a project in the default editor.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			cmd.Printf("launch vscode with %s\n", dir)
			err := prj.OpenEditor(dir)
			if err != nil {
				cmd.Printf("error: %s\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
