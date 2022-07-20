package config

import "github.com/spf13/viper"

func GetPackageDir() string {
	return viper.GetString("packageDir")
}
