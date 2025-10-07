package cli

import (
	"errors"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newDeviceBufferDisableCmd() *cobra.Command {
	var deviceID string
	deviceBucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "disable",
		Short:   "Disable buffering for a device",
		Aliases: []string{"off"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if deviceID == "" {
				return errors.New("device-id is required")
			}

			return withDeviceStore(cmd.Context(), deviceBucket, func(mgr *store.DeviceStore) error {
				info, err := mgr.Get(deviceID)
				if err != nil {
					return err
				}
				info.BufferDuration = ""

				if err := mgr.Upsert(deviceID, info); err != nil {
					return err
				}

				log.Info().Str("device", deviceID).Msg("buffer disabled")
				cmd.Printf("buffer disabled for %s\n", deviceID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().StringVar(&deviceBucket, "device-bucket", deviceBucket, "Device metadata bucket")
	cmd.MarkFlagRequired("device-id")
	return cmd
}
