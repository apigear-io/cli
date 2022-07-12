package mon

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		url    string
		script string
	}
	var options = &ClientOptions{}
	var cmd = &cobra.Command{
		Use:   "feed",
		Short: "Feed a script to a monitor",
		Long:  `Feeds API calls from various sources to the monitor to be displayed. This is mainly to playback recorded API calls.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.script = args[0]
			log.Debug("run script ", options.script)
			switch filepath.Ext(options.script) {
			case ".json", ".ndjson":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				go func(fn string, emitter chan *mon.Event) {
					err := mon.ReadJsonEvents(fn, emitter)
					if err != nil {
						log.Error(err)
					}
				}(options.script, emitter)
				sender.SendEvents(emitter)
			case ".js":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				go func(script string, emitter chan *mon.Event) {
					vm := mon.NewEventScript(emitter)
					err := vm.RunScriptFromFile(options.script)
					if err != nil {
						log.Error(err)
					}
				}(options.script, emitter)
				sender.SendEvents(emitter)
			case ".csv":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				go func(fn string, emitter chan *mon.Event) {
					err := mon.ReadCsvEvents(fn, emitter)
					if err != nil {
						log.Error(err)
					}
				}(options.script, emitter)
				sender.SendEvents(emitter)
			default:
				log.Error("unknown file type: ", options.script)
			}
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://127.0.0.1:5555/monitor/123/", "monitor server address")

	return cmd
}
