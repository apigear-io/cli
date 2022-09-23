package config

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigFile string
var ConfigDir string
var DryRun bool = false

// initConfig reads in config file and ENV variables if set.
func init() {
	debug := os.Getenv("DEBUG") == "1"
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Find home directory.
		// Search config in home directory with name ".apigear" (without extension).
		home, err := os.UserHomeDir()
		ConfigDir = helper.Join(home, ".apigear")
		ConfigFile = viper.ConfigFileUsed()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".apigear")
		viper.SetConfigFile(ConfigFile)
		if debug {
			fmt.Printf("config path dir: %s\n", home)
		}
	}
	if debug {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config file: %v\n", err)
	}

	viper.SetEnvPrefix("apigear")
	viper.AutomaticEnv() // read in environment variables that match
	initPackageDir()
	initRegistryDir()
}

func initPackageDir() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	packageDir := helper.Join(home, ".apigear", "templates")
	viper.SetDefault(KeyPackageDir, packageDir)
	err = os.MkdirAll(packageDir, 0755)
	if err != nil {
		panic(err)
	}
}

func initRegistryDir() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	registryDir := helper.Join(home, ".apigear", "registry")
	viper.SetDefault(KeyRegistryDir, registryDir)
}
