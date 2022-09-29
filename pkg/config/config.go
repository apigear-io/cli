package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	KeyServerPort    = "server_port"
	KeyEditorCommand = "editor_command"
	KeyUpdateChannel = "update_channel"
	KeyRecent        = "recent"
	KeyVersion       = "version"
	KeyTemplatesDir  = "templates_dir"
	KeyRegistryDir   = "registry_dir"
	KeyRegistryUrl   = "registry_url"
	KeyGitAuthToken  = "git_auth_token"
	KeyCommit        = "commit"
	KeyDate          = "date"
)

func GetRecentEntries() []string {
	entries := viper.GetStringSlice(KeyRecent)
	// limit to 5 entries
	if len(entries) > 5 {
		entries = entries[:5]
	}
	return entries
}

func AppendRecentEntry(file string) error {
	// check if file is already in recent
	recent := GetRecentEntries()
	for _, f := range recent {
		if f == file {
			return nil
		}
	}
	// limit to 5 entries
	if len(recent) >= 5 {
		recent = recent[1:]
	}
	recent = append(recent, file)
	viper.Set(KeyRecent, recent)
	return viper.WriteConfig()
}

func RemoveRecentEntry(d string) error {
	recent := GetRecentEntries()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	viper.Set(KeyRecent, recent)
	return viper.WriteConfig()
}

func Set(key string, value any) {
	viper.Set(key, value)
}

func WriteConfig() error {
	return viper.WriteConfig()
}

func GetEditorCommand() string {
	cmd := viper.GetString(KeyEditorCommand)
	if cmd == "" {
		return "code"
	}
	return cmd
}

func GetServerPort() string {
	port := viper.GetString(KeyServerPort)
	if port == "" {
		return "8082"
	}
	return port
}

func GetUpdateChannel() string {
	ch := viper.GetString(KeyUpdateChannel)
	if ch == "" {
		return "stable"
	}
	return ch
}

func RegistryDir() string {
	dir := viper.GetString(KeyRegistryDir)
	if dir == "" {
		return "registry"
	}
	return dir
}

func CachedRegistryPath() string {
	return filepath.Join(RegistryDir(), "registry.json")
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func ConfigFileUsed() string {
	return viper.ConfigFileUsed()
}

func SetVersion(version string) {
	viper.Set(KeyVersion, version)
}

func IsSet(key string) bool {
	return viper.IsSet(key)
}

func Get(key string) string {
	return viper.GetString(key)
}

func GitAuthToken() string {
	return viper.GetString(KeyGitAuthToken)
}

func TemplatesDir() string {
	return viper.GetString(KeyTemplatesDir)
}

func RegistryUrl() string {
	return viper.GetString(KeyRegistryUrl)
}
