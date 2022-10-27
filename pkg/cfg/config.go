package cfg

import (
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	KeyRecent         = "recent"
	KeyServerPort     = "server_port"
	KeyEditorCommand  = "editor_command"
	KeyUpdateChannel  = "update_channel"
	KeyTemplatesDir   = "templates_dir"
	KeyRegistryDir    = "registry_dir"
	KeyRegistryUrl    = "registry_url"
	KeyGitAuthToken   = "git_auth_token"
	KeyGitPublicToken = "git_public_token"
	KeyGitAuthUser    = "git_auth_user"
	KeyVersion        = "version"
	KeyCommit         = "commit"
	KeyDate           = "date"
)

const (
	registryUrl = "https://github.com/apigear-io/template-registry.git"
)

var (
	v           *viper.Viper
	repoToken   string // populated during build-process
	publicToken string // populated during build-process
)

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	cfgDir := helper.Join(home, ".apigear")
	v = NewConfig(cfgDir)
}

func NewConfig(cfgDir string) *viper.Viper {
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
	nv.SetDefault(KeyServerPort, 8085)
	nv.SetDefault(KeyEditorCommand, "code")
	nv.SetDefault(KeyUpdateChannel, "stable")
	nv.SetDefault(KeyVersion, "0.0.0")
	// public repo token for github to avoid rate limit
	nv.SetDefault(KeyGitAuthToken, repoToken)
	nv.SetDefault(KeyGitPublicToken, publicToken)
	nv.SetDefault(KeyGitAuthUser, "jryannel")
	nv.SetDefault(KeyCommit, "none")
	nv.SetDefault(KeyDate, "unknown")
	nv.SetDefault(KeyRegistryUrl, registryUrl)

	// Search config in home directory with name ".apigear" (without extension).

	cfgFile := helper.Join(cfgDir, "config.json")
	nv.AddConfigPath(cfgDir)
	nv.SetConfigType("json")
	nv.SetConfigName("config")

	if !helper.IsFile(cfgFile) {
		err := helper.MakeDir(cfgDir)
		cobra.CheckErr(err)
		err = helper.WriteFile(cfgFile, []byte("{}"))
		cobra.CheckErr(err)
	}

	// If a config file is found, read it in.
	if err := nv.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
	return nv
}

func SetConfig(c *viper.Viper) {
	v = c
}
