package prj

import (
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewInfoCommand returns a new cobra.Command for the "info" command.
func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "display project information",
		Long:  `display detailed project information`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]
			cmd.Printf("# info %s\n", dir)
			info, err := prj.GetProjectInfo(dir)
			if err != nil {
				return err
			}
			cmd.Printf("path: %s\n", info.Path)
			cmd.Printf("name: %s\n", info.Name)
			for _, v := range info.Documents {
				cmd.Printf("  %s\n", v.Name)
			}
			return nil
		},
	}
	return cmd
}
