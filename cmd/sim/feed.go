package sim

import (
	"path/filepath"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net/rpc"

	"github.com/spf13/cobra"
)

type ConsoleHandler struct{}

// Very similar to message handler
func (c ConsoleHandler) HandleMessage(msg rpc.RpcMessage) error {
	log.Debugf("received message: %+v\n", msg)
	return nil
}

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		addr          string
		script        string
		sleepDuration time.Duration
	}
	var options = &ClientOptions{}
	// cmd represents the simCli command
	var cmd = &cobra.Command{
		Use:   "feed",
		Short: "feed simulation from command line",
		Long:  `The feed command allows you to simulate calls from command line.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.script = args[0]
			log.Debug("run script ", options.script)
			switch filepath.Ext(options.script) {
			case ".ndjson":
				emitter := make(chan rpc.RpcMessage)
				writer := ConsoleHandler{}
				sender := rpc.NewRpcSender(writer)
				err := sender.Dial(options.addr)
				if err != nil {
					log.Fatalf("failed to connect to %s: %v", options.addr, err)
				}
				go rpc.ReadJsonMessagesFromFile(options.script, options.sleepDuration, emitter)
				go sender.ReadPump()
				sender.SendMessages(emitter)
			}
		},
	}
	cmd.Flags().DurationVarP(&options.sleepDuration, "sleep", "", 200, "sleep duration between messages")
	cmd.Flags().StringVarP(&options.addr, "addr", "", "localhost:5555", "address of the simulation server")
	return cmd
}
