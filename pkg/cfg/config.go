package cfg

import (
	"fmt"
	"os"
	"sync"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/viper"
)

const (
	KeyRecent        = "recent"
	KeyServerPort    = "server_port"
	KeyEditorCommand = "editor_command"
	KeyUpdateChannel = "update_channel"
	KeyCacheDir      = "templates_dir"
	KeyRegistryDir   = "registry_dir"
	KeyRegistryUrl   = "registry_url"
	KeyVersion       = "version"
	KeyCommit        = "commit"
	KeyDate          = "date"
	KeyWindowHeight  = "window_height"
	KeyWindowWidth   = "window_width"
)

const (
	registryUrl  = "https://github.com/apigear-io/template-registry.git"
	cacheName    = "cache"
	registryName = "registry"
)

var (
	v  *viper.Viper
	rw = sync.RWMutex{}
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfgDir := helper.Join(home, ".apigear")
	vip, err := NewConfig(cfgDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rw.Lock()
	v = vip
	rw.Unlock()
}

func NewConfig(cfgDir string) (*viper.Viper, error) {
	nv := viper.New()

	nv.SetEnvPrefix("apigear")
	nv.AutomaticEnv() // read in environment variables that match

	cacheDir := helper.Join(cfgDir, "cache")

	err := helper.MakeDir(cacheDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create cache dir: %w", err)
	}

	registryDir := helper.Join(cfgDir, "registry")

	err = helper.MakeDir(registryDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry dir: %w", err)
	}

	nv.SetDefault(KeyCacheDir, cacheDir)
	nv.SetDefault(KeyRegistryUrl, registryUrl)
	nv.SetDefault(KeyRegistryDir, registryDir)
	nv.SetDefault(KeyServerPort, 4333)
	nv.SetDefault(KeyEditorCommand, "code")
	nv.SetDefault(KeyUpdateChannel, "stable")
	nv.SetDefault(KeyVersion, "0.0.0")
	nv.SetDefault(KeyWindowWidth, 960)
	nv.SetDefault(KeyWindowHeight, 720)
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
		// try to write a new config file
		err = nv.WriteConfigAs(cfgFile)
		if err != nil {
			return nil, fmt.Errorf("failed to write config file %s: %w", cfgFile, err)
		}
		// try to read the new config file
		err = nv.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to read config file %s: %w", cfgFile, err)
		}
	}
	return nv, nil
}

func SetConfig(c *viper.Viper) {
	rw.Lock()
	v = c
	rw.Unlock()
}
