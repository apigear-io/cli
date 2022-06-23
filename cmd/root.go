package cmd

import (
	"apigear/cmd/conf"
	"apigear/cmd/mon"
	"apigear/cmd/sdk"
	"apigear/cmd/sim"
	"apigear/cmd/tpl"
	"apigear/pkg/log"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool
var dryRun bool

func NewRootCommand() *cobra.Command {
	// cmd represents the base command when called without any subcommands
	cmd := &cobra.Command{
		Use:     "apigear",
		Short:   "apigear creates instrumented SDKs from an API description",
		Long:    `ApiGear allows you to describe interfaces and generate instrumented SDKs out of the descriptions.`,
		Version: "0.0.1",
	}
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.apigear.yaml)")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "dry-run")
	cmd.PersistentFlags().String("env", "development", "environment (development, production, staging)")
	viper.BindPFlag("env", cmd.PersistentFlags().Lookup("env"))
	viper.Set("version", cmd.Version)
	cmd.AddCommand(sdk.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(conf.NewRootCommand())
	cmd.AddCommand(tpl.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(NewCheckCommand())
	cmd.AddCommand(NewYaml2JsonCommand())
	cmd.AddCommand(NewJson2YamlCommand())
	cmd.AddCommand(NewDocsCommand())

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".apigear" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".apigear")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.Set("verbose", verbose)
	log.SetVerbose(verbose)
}
