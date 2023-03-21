package mon

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"

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
			sender := helper.NewHTTPSender(options.url)
			ctrl := helper.NewSenderControl[mon.Event](options.repeat, options.sleep)
			ids := helper.MakeIdGenerator("M")
			ctrl.Run(events, func(event mon.Event) error {
				if event.Timestamp.IsZero() {
					event.Timestamp = time.Now()
				}
				if event.Id == "" {
					event.Id = ids()
				}
				if event.Source == "" {
					event.Source = "123"
				}
				data, _ := json.Marshal(event)
				log.Info().Msgf("-> %s", string(data))
				// send as an array of events
				payload := [1]mon.Event{event}

				return sender.SendValue(payload)
			})
			return nil
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://127.0.0.1:5555/monitor/123", "monitor server address")
	// repeat is -1 for infinite
	cmd.Flags().IntVar(&options.repeat, "repeat", 1, "number of times to repeat the script")
	// sleep is in milliseconds
	cmd.Flags().DurationVar(&options.sleep, "sleep", 0, "sleep between each event")

	return cmd
}
