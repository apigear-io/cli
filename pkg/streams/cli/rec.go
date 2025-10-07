package cli

import "github.com/spf13/cobra"

func newRecordingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "recordings",
		Short:   "Manage session recordings",
		Aliases: []string{"rec", "record"},
	}

	cmd.AddCommand(
		newRecordingsStartCmd(),
		newRecordingsStopCmd(),
		newRecordingsStatusCmd(),
		newRecordingsListCmd(),
		newRecordingsShowCmd(),
		newRecordingsDeleteCmd(),
		newRecordingsPlayCmd(),
		newRecordingsExportCmd(),
	)

	return cmd
}
