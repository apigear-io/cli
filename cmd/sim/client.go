package sim

import (
	"apigear/pkg/log"
	"apigear/pkg/net/rpc"
	"path"
	"time"

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
		Use:   "client",
		Short: "running simu calls from command line",
		Long: `Simlation clients runs simulation calls from command line. 
These calls can be used to test the service.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.script = args[0]
			log.Debug("run script ", options.script)
			switch path.Ext(options.script) {
			case ".ndjson":
				emitter := make(chan rpc.RpcMessage)
				writer := ConsoleHandler{}
				sender := rpc.NewRpcSender(writer)
				err := sender.Dial(options.addr)
				if err != nil {
					log.Fatalf("failed to connect to %s: %v", options.addr, err)
				}
				go rpc.ReadJsonMessagesFromFile(options.script, options.sleepDuration, emitter)
				sender.SendMessages(emitter)
			}
		},
	}
	cmd.Flags().DurationVarP(&options.sleepDuration, "sleep", "", 200, "sleep duration between messages")
	cmd.Flags().StringVarP(&options.addr, "addr", "", "localhost:5555", "address of the simulation server")
	return cmd
}
