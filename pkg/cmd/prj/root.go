package prj

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the mon command
	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"prj"},
		Short:   "Manage apigear projects",
		Long:    `Projects consist of API descriptions, SDK configuration, simulation documents and other files`,
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
