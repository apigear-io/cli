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
	cmd := NewStreamCmd()
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream",
		Short: "manage message streams",
		Long:  "manage device streams, captured live data , manages device metadata, and replays recorded sessions.",
	}
	cmd.PersistentFlags().StringVar(&rootOpts.server, "server", nats.DefaultURL, "NATS server URL")
	cmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "Enable verbose output")
	cmd.AddGroup(&cobra.Group{ID: "record", Title: "stream recording"})
	cmd.AddGroup(&cobra.Group{ID: "session", Title: "recording sessions"})
	cmd.AddGroup(&cobra.Group{ID: "data", Title: "data"})
	cmd.AddGroup(&cobra.Group{ID: "device", Title: "devices"})
	cmd.AddGroup(&cobra.Group{ID: "buffer", Title: "device buffers"})
	cmd.AddCommand(
		newStreamRecordCmd(),
		newStreamStateCmd(),
		newStreamPlayCmd(),
		newStreamStopCmd(),
		newStreamListCmd(),
		newStreamShowCmd(),
		newStreamRemoveCmd(),
		newStreamExportCmd(),
		newStreamImportCmd(),
		newStreamTailCmd(),
		newStreamPublishCmd(),
		newStreamGenerateCmd(),
		newDeviceSetCmd(),
		newDeviceGetCmd(),
		newDeviceListCmd(),
		newDeviceDeleteCmd(),
		newDeviceBufferEnableCmd(),
		newDeviceBufferDisableCmd(),
		newDeviceBufferInfoCmd(),
		newDeviceBufferListCmd(),
	)

	return cmd
}
