package sim

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	// cmd represents the simCli command
	var cmd = &cobra.Command{
		Use:   "simCli",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("simCli called")
		},
	}
	return cmd
}
