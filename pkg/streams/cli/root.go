package cli

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var rootOpts = struct {
	server  string
	verbose bool
}{
	server:  nats.DefaultURL,
	verbose: false,
}

func Execute() {
	cmd := NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "streams",
		Short: "Message capture and playback utilities for NATS",
		Long:  "streams captures live NATS traffic, manages device metadata, and replays recorded sessions for analysis.",
	}
	cmd.PersistentFlags().StringVar(&rootOpts.server, "server", nats.DefaultURL, "NATS server URL")
	cmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "Enable verbose output")

	cmd.AddCommand(
		newDataCmd(),
		newRecordingsCmd(),
		newDeviceCmd(),
		newServeCmd(),
	)
	return cmd
}
