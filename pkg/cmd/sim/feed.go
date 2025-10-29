package sim

import (
	"path/filepath"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/olnk"
	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		addr   string
		script string
		sleep  time.Duration
		repeat int
		batch  int
	}
	var options = &ClientOptions{}
	// cmd represents the simCli command
	var cmd = &cobra.Command{
		Use:     "feed",
		Aliases: []string{"f"},
		Short:   "Feed simulation from command line",
		Long:    `Feed simulation calls using JSON documents from command line`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.script = args[0]
			log.Info().Str("script", options.script).Str("addr", options.addr).Int("repeat", options.repeat).Dur("sleep", options.sleep).Msg("feed simulation")
			feeder := olnk.NewFeeder()
			err := feeder.Connect(cmd.Context(), options.addr)
			if err != nil {
				return err
			}
			defer func() {
				if closeErr := feeder.Close(); closeErr != nil {
					log.Error().Err(closeErr).Msg("failed to close feeder")
				}
			}()
			switch filepath.Ext(options.script) {
			case ".ndjson":
				items, err := helper.ScanFile(options.script)
				if err != nil {
					return err
				}
				ctrl := helper.NewSenderControl[[]byte](options.repeat, options.sleep, options.batch)
				return ctrl.Run(items, feeder.Feed)
			}
			<-cmd.Context().Done()
			log.Info().Msg("done")
			return nil
		},
	}
	cmd.Flags().DurationVarP(&options.sleep, "sleep", "", 100, "sleep duration between messages")
	cmd.Flags().StringVarP(&options.addr, "addr", "", "ws://127.0.0.1:4333/ws", "address of the simulation server")
	cmd.Flags().IntVarP(&options.repeat, "repeat", "", 1, "number of times to repeat the script")
	cmd.Flags().IntVarP(&options.batch, "batch", "", 1, "number of messages to send in a batch")
	return cmd
}
