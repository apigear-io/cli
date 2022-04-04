package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "show program version",
		Long:  `show program version and exit`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("print version document")
		},
	}
	return cmd
}
