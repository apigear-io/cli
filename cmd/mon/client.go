package mon

import (
	"apigear/pkg/log"
	"apigear/pkg/mon"
	"path"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		url    string
		script string
	}
	var options = &ClientOptions{}
	var cmd = &cobra.Command{
		Use:   "client",
		Short: "API monitor client",
		Long:  `The monitor client allows you to send api events the monitor server for testing purposes.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.script = args[0]
			log.Debug("run script ", options.script)
			switch path.Ext(options.script) {
			case ".json", ".ndjson":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				go mon.ReadJsonEvents(options.script, emitter)
				sender.SendEvents(emitter)
			case ".js":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				vm := mon.NewEventScript(emitter)
				vm.RunScript(options.script)
				sender.SendEvents(emitter)
			case ".csv":
				emitter := make(chan *mon.Event)
				sender := mon.NewEventSender(options.url)
				go mon.ReadCsvEvents(options.script, emitter)
				sender.SendEvents(emitter)
			default:
				log.Error("unknown file type: ", options.script)
			}
		},
	}
	cmd.Flags().StringVar(&options.url, "url", "http://127.0.0.1:5555/monitor/123/", "monitor server address")

	return cmd
}
