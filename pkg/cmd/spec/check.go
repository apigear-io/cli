package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewCheckCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "check",
		Aliases: []string{"c", "lint"},
		Short:   "Check document",
		Long:    `Check documents and report errors`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var file = args[0]
			spec, err := spec.CheckFile(file)
			if err != nil {
				return err
			}
			if spec.Valid() {
				fmt.Printf("valid: %s\n", file)
			} else {
				for _, desc := range spec.Errors {
					fmt.Printf("file: %s \n", file)
					fmt.Println(desc.String())
				}
			}
			return nil
		},
	}
	return cmd
}
