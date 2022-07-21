package config

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/spf13/viper"
)

const (
	KeyServerPort    = "server_port"
	KeyEditorCommand = "editor_command"
	KeyUpdateChannel = "update_channel"
	KeyRecent        = "recent"
	KeyVersion       = "version"
	KeyPackageDir    = "package_dir"
)

func GetRecentEntries() []string {
	entries := viper.GetStringSlice(KeyRecent)
	// limit to 5 entries
	if len(entries) > 5 {
		entries = entries[:5]
	}
	return entries
}

func AppendRecentEntry(file string) {
	log.Debugf("Append recent project: %s", file)
	// check if file is already in recent
	recent := GetRecentEntries()
	for _, f := range recent {
		if f == file {
			log.Debugf("File %s is already in recent", file)
			return
		}
	}
	// limit to 5 entries
	if len(recent) >= 5 {
		recent = recent[1:]
	}
	recent = append(recent, file)
	viper.Set(KeyRecent, recent)
	err := viper.WriteConfig()
	if err != nil {
		log.Warnf("Failed to write config: %s", err)
	}
}

func RemoveRecentEntry(d string) {
	recent := GetRecentEntries()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	viper.Set(KeyRecent, recent)
	err := viper.WriteConfig()
	if err != nil {
		log.Warnf("Failed to write config: %s", err)
	}
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

func WriteConfig() error {
	err := viper.WriteConfig()
	if err != nil {
		log.Warnf("Failed to write config: %s", err)
	}
	return err
}

func GetEditorCommand() string {
	return viper.GetString(KeyEditorCommand)
}

func GetServerPort() int {
	return viper.GetInt(KeyServerPort)
}

func GetUpdateChannel() string {
	return viper.GetString(KeyUpdateChannel)
}
