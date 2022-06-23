package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the confOpen command
func NewInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Information about the configuration",
		Long:  `Information about the configuration.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Configuration information:")
			cmd.Printf("  Config file: %s\n", viper.ConfigFileUsed())
		},
	}
	return cmd
}
