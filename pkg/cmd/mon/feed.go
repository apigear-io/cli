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
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				wg.Add(1)
				go func(fn string, emitter chan *mon.Event) {
					defer wg.Done()
					for i := 0; i < options.repeat; i++ {
						err := mon.ReadJsonEvents(fn, emitter)
						if err != nil {
							log.Error().Err(err).Msg("error reading events")
						}
					}
				}(options.script, emitter)
				wg.Add(1)
				go func(emitter chan *mon.Event, sleep time.Duration) {
					defer wg.Done()
					sender.SendEvents(emitter, sleep)
				}(emitter, options.sleep)
				wg.Wait()

			case ".js":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				wg.Add(1)
				go func(script string, emitter chan *mon.Event) {
					defer wg.Done()
					vm := mon.NewEventScript(emitter)
					err := vm.RunScriptFromFile(script)
					if err != nil {
						log.Error().Err(err).Msg("error running script")
					}
				}(options.script, emitter)
				sender.SendEvents(emitter, options.sleep)
				wg.Wait()
			case ".csv":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				wg.Add(1)
				go func(fn string, emitter chan *mon.Event) {
					defer wg.Done()
					err := mon.ReadCsvEvents(fn, emitter)
					if err != nil {
						log.Error().Err(err).Msg("error reading events")
					}
				}(options.script, emitter)
				sender.SendEvents(emitter, options.sleep)
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
