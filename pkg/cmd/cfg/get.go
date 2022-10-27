package cfg

import (
	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display configuration values",
		Long:    `Display the value of a configuration variable`,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// print all settings
				cmd.Println("all settings:")
				for k, v := range cfg.AllSettings() {
					cmd.Printf("  %s: %s\n", k, v)
				}
			} else {
				// print setting by key
				key := args[0]
				if cfg.IsSet(key) {
					cmd.Printf("%s: %s\n", key, cfg.Get(key))
				} else {
					cmd.Printf("key '%s' was never set\n", key)
				}
			}
		},
	}
	return cmd
}
