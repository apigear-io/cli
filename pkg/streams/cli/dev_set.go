package cli

import (
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/spf13/cobra"
)

func newDeviceSetCmd() *cobra.Command {
	var (
		info     store.DeviceInfo
		deviceID string
	)

	cmd := &cobra.Command{
		Use:     "device-set",
		Short:   "create or update a device profile",
		Aliases: []string{"update"},
		GroupID: "device",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withDeviceStore(cmd.Context(), config.DeviceBucket, func(mgr *store.DeviceStore) error {
				if err := mgr.Upsert(deviceID, info); err != nil {
					return err
				}

				cmd.Printf("device %s updated\n", deviceID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device", "", "Device identifier")
	cmd.Flags().StringVar(&info.Description, "description", "", "Device description")
	cmd.Flags().StringVar(&info.Location, "location", "", "Device location")
	cmd.Flags().StringVar(&info.Owner, "owner", "", "Device owner")
	if err := cmd.MarkFlagRequired("device"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}
