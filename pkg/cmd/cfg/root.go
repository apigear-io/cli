package cfg

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the conf command
	cmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg", "c"},
		Short:   "Display the config vars",
		Long:    `Display and edit the configuration variables`,
	}
	cmd.AddCommand(NewInfoCmd())
	cmd.AddCommand(NewGetCmd())
	cmd.AddCommand(NewEnvCommand())
	return cmd
}
