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
	v = viper.New()

	v.SetEnvPrefix("apigear")
	v.AutomaticEnv() // read in environment variables that match

	packageDir := helper.Join(ConfigDir, "templates")
	v.SetDefault(KeyTemplatesDir, packageDir)

	err = helper.MakeDir(packageDir)
	cobra.CheckErr(err)

	registryDir := helper.Join(home, ".apigear", "registry")

	v.SetDefault(KeyRegistryUrl, registryUrl)
	v.SetDefault(KeyRegistryDir, registryDir)
	v.SetDefault(KeyServerPort, 8085)
	v.SetDefault(KeyEditorCommand, "code")
	v.SetDefault(KeyUpdateChannel, "stable")
	v.SetDefault(KeyVersion, "0.0.0")
	v.SetDefault(KeyGitAuthToken, "")
	v.SetDefault(KeyCommit, "none")
	v.SetDefault(KeyDate, "unknown")
	v.SetDefault(KeyRegistryUrl, registryUrl)

	// Search config in home directory with name ".apigear" (without extension).

	ConfigFile = helper.Join(ConfigDir, "config.json")
	v.AddConfigPath(ConfigDir)
	v.SetConfigType("json")
	v.SetConfigName("config")

	if !helper.IsFile(ConfigFile) {
		err := helper.MakeDir(ConfigDir)
		cobra.CheckErr(err)
		err = helper.WriteFile(ConfigFile, []byte("{}"))
		cobra.CheckErr(err)
	}

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
