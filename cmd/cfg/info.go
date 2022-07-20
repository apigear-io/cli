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
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("info:")
			cmd.Printf("  config file: %s\n", viper.ConfigFileUsed())
			cmd.Println("  config:")
			for k := range viper.AllSettings() {
				cmd.Printf("    %s: %v\n", k, viper.Get(k))
			}
		},
	}
	return cmd
}
