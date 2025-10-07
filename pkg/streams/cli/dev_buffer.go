package cli

import (
	"github.com/spf13/cobra"
)

func newDeviceBufferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "buffer",
		Short:   "Manage device buffering",
		Aliases: []string{"buf"},
	}

	cmd.AddCommand(
		newDeviceBufferEnableCmd(),
		newDeviceBufferDisableCmd(),
		newDeviceBufferInfoCmd(),
		newDeviceBufferListCmd(),
	)
	return cmd
}
