package cmd

import (
	"github.com/apigear-io/cli/pkg/cmd/cfg"
	"github.com/apigear-io/cli/pkg/cmd/gen"
	"github.com/apigear-io/cli/pkg/cmd/mon"
	"github.com/apigear-io/cli/pkg/cmd/prj"
	"github.com/apigear-io/cli/pkg/cmd/sim"
	"github.com/apigear-io/cli/pkg/cmd/spec"
	"github.com/apigear-io/cli/pkg/cmd/tpl"
	"github.com/apigear-io/cli/pkg/cmd/x"
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/log"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewRootCommand() *cobra.Command {
	// cmd represents the base command when called without any subcommands
	cmd := &cobra.Command{
		Use:     "apigear",
		Short:   "apigear creates instrumented SDKs from an API description",
		Long:    `ApiGear allows you to describe interfaces and generate instrumented SDKs out of the descriptions.`,
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cobra.OnInitialize(config.InitConfig)
	cmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.apigear.yaml)")
	cmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	cmd.PersistentFlags().BoolVar(&config.DryRun, "dry-run", false, "dry-run")
	cmd.AddCommand(gen.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(cfg.NewRootCommand())
	cmd.AddCommand(tpl.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(spec.NewRootCommand())
	cmd.AddCommand(prj.NewRootCommand())
	cmd.AddCommand(x.NewRootCommand())

	viper.Set("version", cmd.Version)

	return cmd
}
