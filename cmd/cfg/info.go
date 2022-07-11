package cfg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the confOpen command
func NewInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Information about configuration",
		Long:  `Information about configuration.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("info:")
			cmd.Printf("  config file: %s\n", viper.ConfigFileUsed())
		},
	}
	return cmd
}
