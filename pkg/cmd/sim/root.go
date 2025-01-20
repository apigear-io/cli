package sim

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the sim command
	var cmd = &cobra.Command{
		Use:     "simulate",
		Aliases: []string{"sim", "s", "simu"},
		Short:   "Simulate API calls",
		Long:    `Simulate api calls using either a dynamic JS script or a static YAML document`,
	}
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewRunCommand())
	return cmd
}
