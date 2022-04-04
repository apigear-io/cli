package conf

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the conf command
	cmd := &cobra.Command{
		Use:   "conf",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("conf called")
		},
	}
	cmd.AddCommand(openCmd)
	cmd.AddCommand(getCmd)
	return cmd
}

func init() {
}
