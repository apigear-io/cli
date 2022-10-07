package config

import (
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	registryUrl = "https://github.com/apigear-io/template-registry.git"
	ConfigFile  string
	ConfigDir   string
	Verbose     bool
)

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	ConfigDir = helper.Join(home, ".apigear")

	viper.SetEnvPrefix("apigear")
	viper.AutomaticEnv() // read in environment variables that match

	packageDir := helper.Join(ConfigDir, "templates")
	viper.SetDefault(KeyTemplatesDir, packageDir)
	err = helper.MakeDir(packageDir)
	cobra.CheckErr(err)

	viper.SetDefault(KeyRegistryUrl, registryUrl)
	registryDir := helper.Join(home, ".apigear", "registry")

	viper.SetDefault(KeyRegistryDir, registryDir)
	// Search config in home directory with name ".apigear" (without extension).

	ConfigFile = helper.Join(ConfigDir, "config.json")
	viper.AddConfigPath(ConfigDir)
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	if !helper.IsFile(ConfigFile) {
		err := helper.MakeDir(ConfigDir)
		cobra.CheckErr(err)
		err = helper.WriteFile(ConfigFile, []byte("{}"))
		cobra.CheckErr(err)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
