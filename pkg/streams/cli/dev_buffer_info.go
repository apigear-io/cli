package cli

import (
	"errors"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceBufferInfoCmd() *cobra.Command {
	var deviceID string
	deviceBucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "info",
		Short:   "Show buffering status for a device",
		Aliases: []string{"show"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if deviceID == "" {
				return errors.New("device-id is required")
			}

			return withDeviceStore(cmd.Context(), deviceBucket, func(mgr *store.DeviceStore) error {
				info, err := mgr.Get(deviceID)
				if err != nil {
					return err
				}

				buffer := info.BufferDuration
				if buffer == "" {
					buffer = "disabled"
				}

				cmd.Printf("device: %s\n", deviceID)
				cmd.Printf("buffer: %s\n", buffer)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().StringVar(&deviceBucket, "device-bucket", deviceBucket, "Device metadata bucket")
	if err := cmd.MarkFlagRequired("device-id"); err != nil {
		cobra.CheckErr(err)
	}
	return cmd
}
