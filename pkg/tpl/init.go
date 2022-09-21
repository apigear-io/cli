package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/spf13/viper"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	packageDir := helper.Join(home, ".apigear", "templates")
	viper.SetDefault("packageDir", packageDir)
	err = os.MkdirAll(packageDir, 0755)
	if err != nil {
		panic(err)
	}
}
