package cfg

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/spf13/cobra"
)

// openCmd represents the confOpen command
func NewInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info",
		Aliases: []string{"i"},
		Short:   "Display the config information",
		Long:    `Display the config information and the location of the config file`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("info:")
			cmd.Printf("  config file: %s\n", config.ConfigFileUsed())
			cmd.Println("  config:")
			for k, v := range config.AllSettings() {
				cmd.Printf("    %s: %v\n", k, v)
			}
		},
	}
	return cmd
}
