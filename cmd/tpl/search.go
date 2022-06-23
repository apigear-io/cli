package tpl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {

	// cmd represents the pkgSearch command
	var cmd = &cobra.Command{
		Use:   "find",
		Short: "Find a template by name from template registry",
		Long:  `Find a template by name from template registry.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pkgSearch called")
		},
	}
	return cmd
}
