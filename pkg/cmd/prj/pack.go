package prj

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/prj"

	"github.com/spf13/cobra"
)

// NewPackCommand returns a new cobra.Command for the "pack" command.
func NewPackCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pack",
		Short: "Pack a project",
		Long:  `Pack the project and all files into a archive file`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			cmd.Printf("pack project %s\n", source)
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			base := filepath.Base(source)
			target := helper.Join(cwd, fmt.Sprintf("%s.tgz", base))

			target, err = prj.PackProject(source, target)
			if err != nil {
				return err
			}
			cmd.Printf("project %s packed to %s\n", source, target)
			return nil
		},
	}
	return cmd
}
