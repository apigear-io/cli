package prj

import (
	"apigear/pkg/prj"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// NewPackCommand returns a new cobra.Command for the "pack" command.
func NewPackCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pack",
		Short: "Pack a project",
		Long:  `The pack command allows you to pack a project.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			fmt.Printf("pack project %s\n", source)
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}
			base := filepath.Base(source)
			target := filepath.Join(cwd, fmt.Sprintf("%s.tgz", base))

			target, err = prj.PackProject(source, target)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("project %s packed to %s\n", source, target)
		},
	}
	return cmd
}
