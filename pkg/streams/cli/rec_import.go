package cli

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import",
		Short:   "import a recorded stream session from JSONL",
		GroupID: "data",
		RunE: func(cmd *cobra.Command, _ []string) error {
			log.Info().Msg("importing recorded session")
			log.Error().Msg("not implemented yet")
			return session.Import(cmd.Context(), session.ImportOptions{})
		},
	}

	return cmd
}
