package sim

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the sim command
	var cmd = &cobra.Command{
		Use:     "simulate",
		Aliases: []string{"sim", "s"},
		Short:   "Simulates a solution API calls",
		Long:    `A simulation server can run one or more scenarios which simulation API calls.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewServerCommand())
	return cmd
}
