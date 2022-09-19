package mon

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the mon command
	cmd := &cobra.Command{
		Use:     "monitor",
		Aliases: []string{"mon", "m"},
		Short:   "Display monitor API calls",
		Long:    `Display monitored API calls using a monitoring server. SDKs typically create trace points and forward all API traffic to this monitoring service if configured.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewServerCommand())
	return cmd
}
