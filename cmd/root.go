package cmd

import (
	"fmt"
	"objectapi/cmd/conf"
	"objectapi/cmd/mon"
	"objectapi/cmd/sdk"
	"objectapi/cmd/sim"
	"objectapi/cmd/tpl"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool
var dryRun bool

func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "objectapi",
		Short: "objectapi is a tool to manage code generation templates",
		Long: `The ObjectAPI standard allows you to describe objects as API and generate SDKs in different languages. 
Additional API monaitoring and simulation is also supported.`,
	}
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.objectapi.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "dry-run")

	rootCmd.AddCommand(sdk.NewRootCommand())
	rootCmd.AddCommand(mon.NewRootCommand())
	rootCmd.AddCommand(conf.NewRootCommand())
	rootCmd.AddCommand(tpl.NewRootCommand())
	rootCmd.AddCommand(sim.NewRootCommand())
	rootCmd.AddCommand(NewCheckCommand())
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(NewYaml2JsonCommand())
	rootCmd.AddCommand(NewJson2YamlCommand())

	return rootCmd
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

		// Search config in home directory with name ".objectapi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".objectapi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if verbose {
		fmt.Fprintln(os.Stderr, "verbose output")
	}
}
