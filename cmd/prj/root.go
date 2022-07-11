package prj

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the mon command
	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"prj"},
		Short:   "Manages project creation and management",
		Long:    `The project command allows you to create a new project and manage it.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("prj called")
		},
	}
	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewEditCommand())
	cmd.AddCommand(NewImportCommand())
	cmd.AddCommand(NewInfoCommand())
	cmd.AddCommand(NewInitCommand())
	cmd.AddCommand(NewOpenCommand())
	cmd.AddCommand(NewPackCommand())
	cmd.AddCommand(NewRecentCommand())
	cmd.AddCommand(NewShareCommand())
	return cmd
}
