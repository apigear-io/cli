package cmd

import (
	"apigear/cmd/cfg"
	"apigear/cmd/mon"
	"apigear/cmd/prj"
	"apigear/cmd/sdk"
	"apigear/cmd/sim"
	"apigear/cmd/tools"
	"apigear/cmd/tpl"
	"apigear/pkg/config"
	"apigear/pkg/log"

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
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cobra.OnInitialize(config.InitConfig)

	cmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.apigear.yaml)")
	cmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	cmd.PersistentFlags().BoolVar(&config.DryRun, "dry-run", false, "dry-run")
	cmd.AddCommand(sdk.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(cfg.NewRootCommand())
	cmd.AddCommand(tpl.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(prj.NewRootCommand())
	cmd.AddCommand(tools.NewRootCommand())

	viper.Set("version", cmd.Version)

	return cmd
}
