package cli

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceBufferListCmd() *cobra.Command {
	deviceBucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List buffered devices",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), deviceBucket, func(mgr *store.DeviceStore) error {
				entries, err := mgr.List()
				if err != nil {
					return err
				}
				if len(entries) == 0 {
					cmd.Println("no devices found")
					return nil
				}

				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%-20s  %-8s\n", "DEVICE", "BUFFER"); err != nil {
					return err
				}
				for _, entry := range entries {
					if entry.Info.BufferDuration == "" {
						continue
					}
					if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%-20s  %-8s\n", entry.DeviceID, entry.Info.BufferDuration); err != nil {
						return err
					}
				}
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceBucket, "device-bucket", deviceBucket, "Device metadata bucket")
	return cmd
}
