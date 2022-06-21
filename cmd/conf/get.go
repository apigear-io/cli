package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the confGet command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "prints a configuration value",
	Args:  cobra.MaximumNArgs(1),
	Long:  `prints the value of a configuration parameter or all configuration parameters if no key is given`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// print all settings
			for k, v := range viper.AllSettings() {
				cmd.Printf("%s: %s\n", k, v)
			}
		} else {
			// print setting by key
			key := args[0]
			cmd.Printf("%s: %s\n", key, viper.Get(key))
		}
	},
}
