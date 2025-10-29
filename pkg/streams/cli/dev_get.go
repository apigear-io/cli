package cli

import (
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceGetCmd() *cobra.Command {
	var deviceID string
	bucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "get",
		Short:   "Fetch a device profile",
		Aliases: []string{"show"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), bucket, func(mgr *store.DeviceStore) error {
				info, err := mgr.Get(deviceID)
				if err != nil {
					return err
				}

				cmd.Printf("device: %s\n", deviceID)
				cmd.Printf("  description: %s\n", info.Description)
				cmd.Printf("  location:    %s\n", info.Location)
				cmd.Printf("  owner:       %s\n", info.Owner)
				cmd.Printf("  updated:     %s\n", info.Updated.Format(time.RFC3339))
				if info.BufferDuration != "" {
					cmd.Printf("  buffer:      %s\n", info.BufferDuration)
				}
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().StringVar(&bucket, "device-bucket", bucket, "Device metadata bucket")
	if err := cmd.MarkFlagRequired("device-id"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}
