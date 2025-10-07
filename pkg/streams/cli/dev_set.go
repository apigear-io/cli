package cli

import (
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceSetCmd() *cobra.Command {
	var (
		info      store.DeviceInfo
		deviceID  string
		bufferDur time.Duration
	)
	bucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "set",
		Short:   "Create or update a device profile",
		Aliases: []string{"update"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), bucket, func(mgr *store.DeviceStore) error {
				if bufferDur > 0 {
					info.BufferDuration = bufferDur.String()
				}

				if err := mgr.Upsert(deviceID, info); err != nil {
					return err
				}

				cmd.Printf("device %s updated\n", deviceID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().StringVar(&bucket, "device-bucket", bucket, "Device metadata bucket")
	cmd.Flags().StringVar(&info.Description, "description", "", "Device description")
	cmd.Flags().StringVar(&info.Location, "location", "", "Device location")
	cmd.Flags().StringVar(&info.Owner, "owner", "", "Device owner")
	cmd.Flags().DurationVar(&bufferDur, "buffer", 0, "Optional rolling buffer window (e.g. 5m)")
	cmd.MarkFlagRequired("device-id")

	return cmd
}
