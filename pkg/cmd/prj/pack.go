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
	var dir string
	var cmd = &cobra.Command{
		Use:   "pack",
		Short: "pack project",
		Long:  `pack the project and all files into a archive file for sharing`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := filepath.Abs(dir)
			if err != nil {
				return err
			}
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			cmd.Printf("pack project %s\n", dir)
			base := filepath.Base(dir)
			target := helper.Join(cwd, "..", fmt.Sprintf("%s.tgz", base))

			target, err = prj.PackProject(dir, target)
			if err != nil {
				return err
			}
			cmd.Printf("project %s packed to %s\n", dir, target)
			return nil
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "project directory to pack")
	cmd.MarkFlagRequired("dir")
	return cmd
}
