package cli

import "github.com/spf13/cobra"

func newDeviceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "device",
		Short:   "Manage device metadata and buffering",
		Aliases: []string{"dev"},
	}

	cmd.AddCommand(
		newDeviceSetCmd(),
		newDeviceGetCmd(),
		newDeviceListCmd(),
		newDeviceDeleteCmd(),
		newDeviceBufferCmd(),
	)
	return cmd
}
