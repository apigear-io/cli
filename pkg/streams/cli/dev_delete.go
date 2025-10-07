package cli

import (
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceDeleteCmd() *cobra.Command {
	var deviceID string
	bucket := config.DeviceBucket

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Remove a device profile",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), bucket, func(mgr *store.DeviceStore) error {
				err := mgr.Delete(deviceID)
				if err != nil {
					return err
				}
				cmd.Printf("device %s deleted\n", deviceID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device-id", "", "Device identifier")
	cmd.Flags().StringVar(&bucket, "device-bucket", bucket, "Device metadata bucket")
	cmd.MarkFlagRequired("device-id")

	return cmd
}
