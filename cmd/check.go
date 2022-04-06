package cmd

import (
	"fmt"
	"log"
	"objectapi/pkg/spec"

	"github.com/spf13/cobra"
)

func NewCheckCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "check",
		Short: "check document",
		Long:  `check documents and report errors`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var file = args[0]
			result, err := spec.CheckFile(file)
			if err != nil {
				panic(err)
			}
			if result.Valid() {
				fmt.Printf("valid: %s\n", file)
			} else {
				log.Printf("invalid: %s\n", file)
				for _, desc := range result.Errors() {
					log.Println(desc.String())
				}
			}
		},
	}
	return cmd
}
