package config

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	registryUrl = "https://github.com/apigear-io/template-registry.git"
	ConfigFile  string
	ConfigDir   string
)

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	debug := os.Getenv("DEBUG") == "1"
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
		ConfigDir = helper.Dir(ConfigFile)
	} else {
		// Find home directory.
		// Search config in home directory with name ".apigear" (without extension).
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		ConfigDir = helper.Join(home, ".apigear")
		ConfigFile = helper.Join(ConfigDir, "config.json")
		cobra.CheckErr(err)
		viper.AddConfigPath(ConfigDir)
		viper.SetConfigType("json")
		viper.SetConfigName("config")
		if debug {
			fmt.Printf("config path dir: %s\n", home)
		}
	}
	if !helper.IsFile(ConfigFile) {
		err := helper.MakeDir(ConfigDir)
		if err != nil {
			fmt.Printf("failed to create config dir: %s\n", ConfigDir)
		}
		err = helper.WriteFile(ConfigFile, []byte("{}"))
		if err != nil {
			fmt.Printf("failed to create config file: %s\n", ConfigFile)
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
	initTemplates()
	initRegistry()
}

func initTemplates() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	packageDir := helper.Join(home, ".apigear", "templates")
	viper.SetDefault(KeyTemplatesDir, packageDir)
	err = os.MkdirAll(packageDir, 0755)
	if err != nil {
		panic(err)
	}
}

func initRegistry() {
	viper.SetDefault(KeyRegistryUrl, registryUrl)
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	registryDir := helper.Join(home, ".apigear", "registry")
	viper.SetDefault(KeyRegistryDir, registryDir)
}
