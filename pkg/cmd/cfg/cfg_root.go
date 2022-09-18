package cfg

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the conf command
	cmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg"},
		Short:   "commands related to application configuration",
		Long:    `The config command allows you to manage application configurations.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewInfoCmd())
	cmd.AddCommand(NewGetCmd())
	return cmd
}

func init() {
}
