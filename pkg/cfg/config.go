package cfg

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	KeyRecent        = "recent"
	KeyServerPort    = "server_port"
	KeyEditorCommand = "editor_command"
	KeyUpdateChannel = "update_channel"
	KeyTemplatesDir  = "templates_dir"
	KeyRegistryDir   = "registry_dir"
	KeyRegistryUrl   = "registry_url"
	KeyVersion       = "version"
	KeyCommit        = "commit"
	KeyDate          = "date"
)

const (
	registryUrl = "https://github.com/apigear-io/template-registry.git"
)

var (
	v *viper.Viper
)

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	cfgDir := helper.Join(home, ".apigear")
	vip, err := NewConfig(cfgDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	v = vip
}

func NewConfig(cfgDir string) (*viper.Viper, error) {
	nv := viper.New()

	nv.SetEnvPrefix("apigear")
	nv.AutomaticEnv() // read in environment variables that match

	packageDir := helper.Join(cfgDir, "templates")
	nv.SetDefault(KeyTemplatesDir, packageDir)

	err := helper.MakeDir(packageDir)
	cobra.CheckErr(err)

	registryDir := helper.Join(cfgDir, "registry")

	nv.SetDefault(KeyRegistryUrl, registryUrl)
	nv.SetDefault(KeyRegistryDir, registryDir)
	nv.SetDefault(KeyServerPort, 4333)
	nv.SetDefault(KeyEditorCommand, "code")
	nv.SetDefault(KeyUpdateChannel, "stable")
	nv.SetDefault(KeyVersion, "0.0.0")
	// public repo token for github to avoid rate limit
	nv.SetDefault(KeyCommit, "none")
	nv.SetDefault(KeyDate, "unknown")

	// Search config in home directory with name ".apigear" (without extension).

	cfgFile := helper.Join(cfgDir, "config.json")
	nv.AddConfigPath(cfgDir)
	nv.SetConfigType("json")
	nv.SetConfigName("config")

	if !helper.IsFile(cfgFile) {
		err := helper.MakeDir(cfgDir)
		if err != nil {
			return nil, fmt.Errorf("failed to create config dir: %w", err)
		}
		err = helper.WriteFile(cfgFile, []byte("{}"))
		if err != nil {
			return nil, fmt.Errorf("failed to create config file: %w", err)
		}
	}

	// If a config file is found, read it in.
	err = nv.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	return nv, nil
}

func SetConfig(c *viper.Viper) {
	v = c
}
