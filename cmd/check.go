package cmd

import (
	"fmt"
	"objectapi/pkg/logger"
	"objectapi/pkg/spec"

	"github.com/spf13/cobra"
)

var log = logger.Get()

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
				fmt.Println("valid")
			} else {
				log.Info("invalid")
				for _, desc := range result.Errors() {
					log.Info(desc.String())
				}
			}
		},
	}
	return cmd
}
