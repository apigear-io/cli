package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamImportCmd() *cobra.Command {
	opts := &session.ImportOptions{
		SessionBucket: config.SessionBucket,
		DeviceBucket:  config.DeviceBucket,
	}

	cmd := &cobra.Command{
		Use:     "import",
		Short:   "import a recorded stream session from JSONL",
		GroupID: "data",
		RunE: func(cmd *cobra.Command, _ []string) error {
			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose

			file, err := resolveImportReader(opts.InputPath)
			if err != nil {
				return err
			}
			var closeFn func() error
			if file != nil {
				opts.Reader = file
				closeFn = file.Close
			} else {
				opts.Reader = os.Stdin
			}

			if err := session.Import(cmd.Context(), *opts); err != nil {
				if closeFn != nil {
					if closeErr := closeFn(); closeErr != nil {
						return errors.Join(err, closeErr)
					}
				}
				return err
			}

			if closeFn != nil {
				if err := closeFn(); err != nil {
					return err
				}
			}

			if file != nil {
				cmd.Printf("session imported from %s\n", opts.InputPath)
			} else {
				cmd.Println("session imported from stdin")
			}
			return nil
		},
	}

	opts.InputPath = "-"
	opts.DeviceID = "123"

	cmd.Flags().StringVar(&opts.InputPath, "input", opts.InputPath, "Source JSONL file (use '-' for stdin)")
	cmd.Flags().StringVar(&opts.DeviceID, "device", opts.DeviceID, "Device identifier (auto-created if doesn't exist)")

	return cmd
}

func resolveImportReader(path string) (*os.File, error) {
	if path == "" || path == "-" {
		return nil, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open import file: %w", err)
	}

	return file, nil
}
