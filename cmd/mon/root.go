package mon

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the mon command
	cmd := &cobra.Command{
		Use:     "monitor",
		Aliases: []string{"mon", "m"},
		Short:   "Monitor API calls from client libraries",
		Long: `Typical a SDK contains trace calls which are send to the API monitor. 
		This monitor then can be used to display and analyze the API calls.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewServerCommand())
	return cmd
}
