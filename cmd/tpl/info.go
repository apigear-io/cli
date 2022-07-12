package tpl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewOpenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Shows information about a template.",
		Long:  `Shows the information and local path of the named template.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("info called")
		},
	}
	return cmd
}
