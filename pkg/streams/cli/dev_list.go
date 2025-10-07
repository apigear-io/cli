package cli

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceListCmd() *cobra.Command {
	bucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List device profiles",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), bucket, func(mgr *store.DeviceStore) error {
				entries, err := mgr.List()
				if err != nil {
					return err
				}

				if len(entries) == 0 {
					cmd.Println("no devices found")
					return nil
				}

				fmt.Fprintf(cmd.OutOrStdout(), "%-20s  %-20s  %-20s  %-20s  %-8s  %s\n", "DEVICE", "DESCRIPTION", "LOCATION", "OWNER", "BUFFER", "UPDATED")
				for _, entry := range entries {
					fmt.Fprintf(cmd.OutOrStdout(), "%-20s  %-20s  %-20s  %-20s  %-8s  %s\n",
						entry.DeviceID,
						entry.Info.Description,
						entry.Info.Location,
						entry.Info.Owner,
						entry.Info.BufferDuration,
						entry.Info.Updated.Format(time.RFC3339),
					)
				}
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&bucket, "device-bucket", bucket, "Device metadata bucket")
	return cmd
}
