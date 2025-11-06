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
		Use:     "device-rm",
		Short:   "remove a device profile",
		Aliases: []string{"dev-rm"},
		GroupID: "device",
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

	cmd.Flags().StringVar(&deviceID, "device", "", "Device identifier")
	cmd.Flags().StringVar(&bucket, "device-bucket", bucket, "Device metadata bucket")
	if err := cmd.MarkFlagRequired("device"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}
