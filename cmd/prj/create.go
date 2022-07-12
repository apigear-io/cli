package prj

import (
	"github.com/apigear-io/cli/pkg/log"

	"github.com/spf13/cobra"
)

type createOptions struct {
	Type string
	Name string
}

func NewCreateCommand() *cobra.Command {
	var o createOptions
	var cmd = &cobra.Command{
		Use:   "create doc-type doc-name",
		Short: "Create a new document inside current project",
		Long:  `The create command allows you to create a new document inside current project.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			o.Type = args[0]
			o.Name = args[1]
			log.Debug("create project %#v", o)
		},
	}
	return cmd
}
