package prj

import (
	"github.com/apigear-io/lib/prj"

	"github.com/spf13/cobra"
)

// NewInfoCommand returns a new cobra.Command for the "info" command.
func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Show project information",
		Long:  `The info command allows you to show project information.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			cmd.Printf("# info %s\n", dir)
			info, err := prj.GetProjectInfo(dir)
			if err != nil {
				cmd.Printf("error: %s\n", err)
				return
			}
			cmd.Printf("path: %s\n", info.Path)
			cmd.Printf("name: %s\n", info.Name)
			for _, v := range info.Documents {
				cmd.Printf("  %s\n", v.Name)
			}

		},
	}
	return cmd
}
