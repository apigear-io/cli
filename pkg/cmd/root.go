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
)

func Must(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("fatal error")
	}
}

func NewRootCommand() *cobra.Command {
	cobra.OnInitialize(config.InitConfig)
	// cmd represents the base command when called without any subcommands
	cmd := &cobra.Command{
		Use:   "apigear",
		Short: "apigear creates instrumented SDKs from an API description",
		Long:  `ApiGear allows you to describe interfaces and generate instrumented SDKs out of the descriptions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.apigear/config.json)")
	cmd.AddCommand(gen.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(cfg.NewRootCommand())
	cmd.AddCommand(tpl.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(spec.NewRootCommand())
	cmd.AddCommand(prj.NewRootCommand())
	cmd.AddCommand(x.NewRootCommand())
	cmd.AddCommand(NewUpdateCommand())
	cmd.AddCommand(NewVersionCommand())

	return cmd
}
