package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigFile string
var Verbose bool
var DryRun bool = false

// initConfig reads in config file and ENV variables if set.
func InitConfig() {	
	debug := os.Getenv("DEBUG") == "1"
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".apigear" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".apigear")
		viper.SetConfigFile(filepath.Join(home, ".apigear", "config.json"))
		if debug {
			fmt.Printf("config path dir: %s\n", home)
		}
	}
	if debug {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Debugf("no config file: %s", err)
	}

	viper.SetEnvPrefix("apigear")
	viper.AutomaticEnv() // read in environment variables that match
	log.Config(Verbose, debug)
}
