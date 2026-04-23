package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamExportCmd() *cobra.Command {
	opts := &session.ExportOptions{
		Bucket: config.SessionBucket,
	}
	var deviceID string

	cmd := &cobra.Command{
		Use:     "export",
		Short:   "export a recorded stream session",
		GroupID: "data",
		RunE: func(cmd *cobra.Command, _ []string) error {
			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose

			// Validate that either --session or --device is provided
			if opts.SessionID == "" && deviceID == "" {
				return fmt.Errorf("either --session or --device must be specified")
			}
			if opts.SessionID != "" && deviceID != "" {
				return fmt.Errorf("cannot specify both --session and --device")
			}

			// If device is specified, find the latest session for that device
			if deviceID != "" {
				var foundSession *session.Metadata
				if err := withSessionManager(cmd.Context(), opts.Bucket, func(mgr *session.SessionStore) error {
					sessions, err := mgr.List()
					if err != nil {
						return fmt.Errorf("list sessions: %w", err)
					}

					// Find the most recent session for this device
					var latestSession *session.Metadata
					for i := range sessions {
						if sessions[i].DeviceID == deviceID {
							if latestSession == nil || sessions[i].Start.After(latestSession.Start) {
								latestSession = &sessions[i]
							}
						}
					}

					if latestSession == nil {
						return fmt.Errorf("no sessions found for device %s", deviceID)
					}

					foundSession = latestSession
					opts.SessionID = latestSession.SessionID
					return nil
				}); err != nil {
					return err
				}

				// Print info about the found session
				cmd.Printf("searching latest device session, found: %s, recorded at: %s\n",
					foundSession.SessionID,
					foundSession.Start.Format("2006-01-02 15:04:05"))
			}

			file, err := resolveExportWriter(opts.OutputPath)
			if err != nil {
				return err
			}
			opts.Writer = file
			defer file.Close()

			// Get message count before export using withSessionManager
			var messageCount int
			if err := withSessionManager(cmd.Context(), opts.Bucket, func(mgr *session.SessionStore) error {
				meta, err := mgr.Info(opts.SessionID)
				if err != nil {
					return err
				}
				messageCount = meta.MessageCount
				return nil
			}); err != nil {
				return err
			}

			if err := session.Export(cmd.Context(), *opts); err != nil {
				return err
			}

			cmd.Printf("session %s exported to %s (%d messages)\n", opts.SessionID, opts.OutputPath, messageCount)
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.SessionID, "session", "", "Session identifier to export")
	cmd.Flags().StringVar(&deviceID, "device", "", "Device identifier (exports latest session)")
	cmd.Flags().StringVar(&opts.OutputPath, "output", "", "Destination JSONL file")
	cmd.MarkFlagsMutuallyExclusive("session", "device")
	if err := cmd.MarkFlagRequired("output"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}

func resolveExportWriter(path string) (*os.File, error) {
	if path == "" {
		return nil, fmt.Errorf("output path cannot be empty")
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
