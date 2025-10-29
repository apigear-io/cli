package cli

import (
	"errors"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newDeviceBufferEnableCmd() *cobra.Command {
	var (
		deviceID string
		window   time.Duration
	)
	deviceBucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "enable",
		Short:   "Enable rolling buffering for a device",
		Aliases: []string{"on"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if deviceID == "" {
				return errors.New("device-id is required")
			}
			if window <= 0 {
				return errors.New("window must be positive")
			}

			return withDeviceStore(cmd.Context(), deviceBucket, func(mgr *store.DeviceStore) error {
				info, err := mgr.Get(deviceID)
				if err != nil {
					if !errors.Is(err, jetstream.ErrKeyNotFound) {
						return err
					}
					info = store.DeviceInfo{}
				}
				info.BufferDuration = window.String()

				if err := mgr.Upsert(deviceID, info); err != nil {
					return err
				}

				log.Info().Str("device", deviceID).Dur("window", window).Msg("buffer enabled")
				cmd.Printf("buffer enabled for %s (%s)\n", deviceID, window)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().DurationVar(&window, "window", 0, "Rolling buffer window (e.g. 5m)")
	cmd.Flags().StringVar(&deviceBucket, "device-bucket", deviceBucket, "Device metadata bucket")
	if err := cmd.MarkFlagRequired("device-id"); err != nil {
		cobra.CheckErr(err)
	}
	if err := cmd.MarkFlagRequired("window"); err != nil {
		cobra.CheckErr(err)
	}
	return cmd
}
