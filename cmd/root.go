package cmd

import (
	"apigear/cmd/conf"
	"apigear/cmd/mon"
	"apigear/cmd/prj"
	"apigear/cmd/sdk"
	"apigear/cmd/sim"
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
	}
	cobra.OnInitialize(config.InitConfig)

	cmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.apigear.yaml)")
	cmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	cmd.PersistentFlags().BoolVar(&config.DryRun, "dry-run", false, "dry-run")
	cmd.PersistentFlags().String("env", "development", "environment (development, production, staging)")
	cmd.AddCommand(sdk.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(conf.NewRootCommand())
	cmd.AddCommand(tpl.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(prj.NewRootCommand())
	cmd.AddCommand(NewCheckCommand())
	cmd.AddCommand(NewYaml2JsonCommand())
	cmd.AddCommand(NewJson2YamlCommand())
	cmd.AddCommand(NewDocsCommand())

	viper.Set("version", cmd.Version)

	return cmd
}
