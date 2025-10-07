package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newRecordingsExportCmd() *cobra.Command {
	opts := &session.ExportOptions{
		Bucket: config.SessionBucket,
	}

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export a recorded session to JSONL",
		RunE: func(cmd *cobra.Command, _ []string) error {
			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose

			file, err := resolveExportWriter(opts.OutputPath)
			if err != nil {
				return err
			}
			if file != nil {
				defer file.Close()
				opts.Writer = file
			} else {
				opts.Writer = os.Stdout
			}

			if err := session.Export(cmd.Context(), *opts); err != nil {
				return err
			}

			if file != nil {
				cmd.Printf("session %s exported to %s\n", opts.SessionID, opts.OutputPath)
			} else {
				cmd.Printf("session %s exported to stdout\n", opts.SessionID)
			}
			return nil
		},
	}

	opts.OutputPath = "-"

	cmd.Flags().StringVar(&opts.SessionID, "session-id", "", "Session identifier to export")
	cmd.Flags().StringVar(&opts.Bucket, "session-bucket", opts.Bucket, "Key-value bucket containing session metadata")
	cmd.Flags().StringVar(&opts.OutputPath, "output", opts.OutputPath, "Destination JSONL file (use '-' for stdout)")
	cmd.MarkFlagRequired("session-id")

	return cmd
}

func resolveExportWriter(path string) (*os.File, error) {
	if path == "" || path == "-" {
		return nil, nil
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create export dir: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create export file: %w", err)
	}

	return file, nil
}
