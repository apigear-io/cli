package cli

import (
	"time"

	"github.com/apigear-io/cli/pkg/streams/msgio"
	"github.com/spf13/cobra"
)

func newDataGenerateCmd() *cobra.Command {
	opts := &msgio.GenerateOptions{
		Count: 1000,
		Seed:  time.Now().UnixNano(),
	}

	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "Generate JSONL monitor data from a template",
		Long:    "Render a Go template repeatedly with faker-backed helpers to build large JSONL files for testing.",
		Aliases: []string{"gen"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return msgio.Generate(*opts)
		},
	}

	cmd.Flags().StringVarP(&opts.TemplatePath, "template", "t", "", "Template file describing a single JSON object")
	cmd.Flags().StringVarP(&opts.OutputPath, "output", "o", "", "Destination JSONL file (defaults to stdout)")
	cmd.Flags().IntVarP(&opts.Count, "count", "c", opts.Count, "Number of JSON objects to generate")
	cmd.Flags().Int64Var(&opts.Seed, "seed", opts.Seed, "Random seed for faker data")
	cmd.MarkFlagRequired("template")

	return cmd
}
