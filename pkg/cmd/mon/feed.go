package mon

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/runtime/monitoring"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		url    string        // monitor server url
		script string        // script to run
		repeat int           // -1 for infinite
		sleep  time.Duration // sleep between each event
	}
	var options = &ClientOptions{}
	var cmd = &cobra.Command{
		Use:   "feed",
		Short: "Feed a script to a monitor",
		Long:  `Feeds API calls from various sources to the monitor to be displayed. This is mainly to playback recorded API calls.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			options.script = args[0]
			logging.Debug().Msgf("run script %s", options.script)
			var events []monitoring.Event
			var err error
			switch foundation.Ext(options.script) {
			case ".json", ".ndjson":
				events, err = monitoring.ReadJsonEvents(options.script)
				logging.Debug().Msgf("read %d events", len(events))
				if err != nil {
					return fmt.Errorf("error reading events: %w", err)
				}
			case ".js":
				vm := monitoring.NewEventScript()
				events, err = vm.RunScriptFromFile(options.script)
				if err != nil {
					return fmt.Errorf("error running script: %w", err)
				}
			case ".csv":
				events, err = monitoring.ReadCsvEvents(options.script)
				if err != nil {
					return fmt.Errorf("error reading events: %w", err)
				}
			default:
				return fmt.Errorf("unsupported script type: %s", options.script)
			}
			if len(events) == 0 {
				return fmt.Errorf("no events to send")
			}
			sender := foundation.NewHTTPSender(options.url)
			ctrl := foundation.NewSenderControl[monitoring.Event](options.repeat, options.sleep)
			err = ctrl.Run(events, func(event monitoring.Event) error {
				if event.Source == "" {
					event.Source = "123"
				}
				// send as an array of events
				payload := [1]monitoring.Event{event}

				return sender.SendValue(payload)
			})
			if err != nil {
				logging.Warn().Msgf("error sending events: %s", err)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://localhost:5555/monitor/123", "monitor server address")
	// repeat is -1 for infinite
	cmd.Flags().IntVar(&options.repeat, "repeat", 1, "number of times to repeat the script")
	// sleep is in milliseconds
	cmd.Flags().DurationVar(&options.sleep, "sleep", 0, "sleep between each event")

	return cmd
}
