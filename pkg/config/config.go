package config

import (
	"github.com/apigear-io/cli/pkg/log"

	"github.com/spf13/viper"
)

func ReadRecentProjects() []string {
	return viper.GetStringSlice("recent")
}

func AppendRecentProject(file string) {
	// check if file is already in recent
	recent := ReadRecentProjects()
	for _, f := range recent {
		if f == file {
			log.Debugf("File %s is already in recent", file)
			return
		}
	}
	viper.Set("recent", append(recent, file))
	err := viper.WriteConfig()
	if err != nil {
		log.Warnf("Failed to write config: %s", err)
	}
}

func RemoveRecentFile(d string) {
	recent := ReadRecentProjects()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	viper.Set("recent", recent)
	err := viper.WriteConfig()
	if err != nil {
		log.Warnf("Failed to write config: %s", err)
	}
}
