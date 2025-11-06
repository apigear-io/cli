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
		Use:     "buffer-off",
		Short:   "disable buffering for a device",
		Aliases: []string{"buff-off"},
		GroupID: "buffer",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if deviceID == "" {
				return errors.New("device is required")
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

	cmd.Flags().StringVar(&deviceID, "device", "", "Device identifier")
	cmd.Flags().StringVar(&deviceBucket, "device-bucket", deviceBucket, "Device metadata bucket")
	if err := cmd.MarkFlagRequired("device"); err != nil {
		cobra.CheckErr(err)
	}
	return cmd
}
