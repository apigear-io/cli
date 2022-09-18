package tpl

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	packageDir := filepath.Join(home, ".apigear", "templates")
	viper.SetDefault("packageDir", packageDir)
	err = os.MkdirAll(packageDir, 0755)
	if err != nil {
		panic(err)
	}
}
