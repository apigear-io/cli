package stim

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the sim command
	var cmd = &cobra.Command{
		Use:     "stimulate",
		Aliases: []string{"stim"},
		Short:   "Stimulate API calls to services",
		Long:    `Stimulate API calls using either a dynamic JS script to services`,
	}
	cmd.AddCommand(NewRunCommand())
	return cmd
}
