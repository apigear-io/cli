package tpl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "tpl",
		Short: "manage code generation templates",
		Long:  `SDK templates can be installed from git repositories and used to generate code. The templates are stored in a local folder.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tpl called")
		},
	}
	cmd.AddCommand(NewSearchCommand())
	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewOpenCommand())
	cmd.AddCommand(NewListCommand())
	return cmd
}
