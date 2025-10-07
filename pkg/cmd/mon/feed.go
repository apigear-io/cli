package mon

import (
	"fmt"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		url      string        // monitor server url
		script   string        // script to run
		repeat   int           // -1 for infinite
		interval time.Duration // sleep between each event
		deviceId string        // device id to use
		batch    int           // number of events to send in a batch
	}
	var options = &ClientOptions{}
	var cmd = &cobra.Command{
		Use:   "feed",
		Short: "Feed a script to a monitor",
		Long:  `Feeds API calls from various sources to the monitor to be displayed. This is mainly to playback recorded API calls.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			options.script = args[0]
			log.Debug().Msgf("run script %s", options.script)
			var events []mon.Event
			var err error
			switch helper.Ext(options.script) {
			case ".json", ".ndjson":
				events, err = mon.ReadJsonEvents(options.script)
				log.Debug().Msgf("read %d events", len(events))
				if err != nil {
					return fmt.Errorf("error reading events: %w", err)
				}
			case ".js":
				vm := mon.NewEventScript()
				events, err = vm.RunScriptFromFile(options.script)
				if err != nil {
					return fmt.Errorf("error running script: %w", err)
				}
			case ".csv":
				events, err = mon.ReadCsvEvents(options.script)
				if err != nil {
					return fmt.Errorf("error reading events: %w", err)
				}
			default:
				return fmt.Errorf("unsupported script type: %s", options.script)
			}
			if len(events) == 0 {
				return fmt.Errorf("no events to send")
			}
			url := strings.Join([]string{strings.TrimRight(options.url, "/"), "monitor", options.deviceId}, "/")
			log.Info().Msgf("sending %d events to %s", len(events), url)
			sender := helper.NewHTTPSender(url)
			ctrl := helper.NewSenderControl[mon.Event](options.repeat, options.interval, options.batch)
			err = ctrl.Run(events, func(event mon.Event) error {
				if event.Device == "" {
					event.Device = options.deviceId
				}
				// send as an array of events

				payload := [1]mon.Event{event}
				log.Info().Msgf("send event %s %s %s", event.Device, event.Type.String(), event.Symbol)
				return sender.SendValue(payload)
			})
			if err != nil {
				log.Warn().Msgf("error sending events: %s", err)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://localhost:5555", "monitor server address")
	// repeat is -1 for infinite
	cmd.Flags().IntVar(&options.repeat, "repeat", 1, "number of times to repeat the script")
	// sleep is in milliseconds
	cmd.Flags().DurationVar(&options.interval, "interval", 100*time.Millisecond, "interval between each event")
	// deviceId to use
	cmd.Flags().StringVar(&options.deviceId, "device", "123", "device id to use")
	// batch size
	cmd.Flags().IntVar(&options.batch, "batch", 1, "number of events to send in a batch")

	return cmd
}
