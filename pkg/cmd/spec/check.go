package spec

import (
	"fmt"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewCheckCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "check",
		Aliases: []string{"c"},
		Short:   "Check document",
		Long:    `Check documents and report errors`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var file = args[0]
			switch filepath.Ext(file) {
			case ".json", ".yaml":
				result, err := spec.CheckFile(file)
				if err != nil {
					return err
				}
				if result.Valid() {
					fmt.Printf("valid: %s\n", file)
				} else {
					fmt.Printf("invalid: %s\n", file)
					for _, desc := range result.Errors() {
						fmt.Println(desc.String())
					}
				}
			case ".csv":
				err := spec.CheckCsvFile(file)
				if err != nil {
					return err
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			case ".ndjson":
				err := spec.CheckNdjsonFile(file)
				if err != nil {
					return err
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			case ".idl":
				err := spec.CheckIdlFile(file)
				if err != nil {
					return err
				} else {
					fmt.Printf("valid: %s\n", file)
				}
			default:
				fmt.Printf("unknown file type %s", file)
			}
			return nil
		},
	}
	return cmd
}
