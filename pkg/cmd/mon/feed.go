package mon

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

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
			wg := &sync.WaitGroup{}
			switch filepath.Ext(options.script) {
			case ".json", ".ndjson":
				sender := mon.NewEventSender(options.url)
				for i := 0; i < options.repeat; i++ {
					events, err := mon.ReadJsonEvents(options.script)
					if err != nil {
						log.Error().Err(err).Msg("error reading events")
					}
					sender.SendEvents(events, options.sleep)
				}
			case ".js":
				sender := mon.NewEventSender(options.url)
				vm := mon.NewEventScript()
				events, err := vm.RunScriptFromFile(options.script)
				if err != nil {
					log.Error().Err(err).Msg("error running script")
				}
				sender.SendEvents(events, options.sleep)
			case ".csv":
				sender := mon.NewEventSender(options.url)
				events, err := mon.ReadCsvEvents(options.script)
				if err != nil {
					log.Error().Err(err).Msg("error reading events")
				}
				sender.SendEvents(events, options.sleep)
				wg.Wait()
			default:
				return fmt.Errorf("unsupported script type: %s", options.script)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://127.0.0.1:5555/monitor/123/", "monitor server address")
	// repeat is -1 for infinite
	cmd.Flags().IntVar(&options.repeat, "repeat", 1, "number of times to repeat the script")
	// sleep is in milliseconds
	cmd.Flags().DurationVar(&options.sleep, "sleep", 0, "sleep between each event")

	return cmd
}
