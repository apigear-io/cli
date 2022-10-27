package config

import (
	"path/filepath"
)

func ConfigDir() string {
	file := v.ConfigFileUsed()
	return filepath.Dir(file)
}

// AppendRecentEntry appends a new entry to the list of recent entries
// entries are limited to 5
// the most recent entry is at the beginning of the list
// stores the list in the config file
func AppendRecentEntry(file string) error {
	recent := RecentEntries()
	for i, f := range recent {
		if f == file {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	if len(recent) >= 5 {
		recent = recent[1:]
	}
	// prepend the new entry
	recent = append([]string{file}, recent...)
	v.Set(KeyRecent, recent)
	return v.WriteConfig()
}

// RemoveRecentEntry removes a recent entry from the list
func RemoveRecentEntry(d string) error {
	recent := RecentEntries()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	v.Set(KeyRecent, recent)
	return v.WriteConfig()
}

// RecentEntries returns the list of recent entries
func RecentEntries() []string {
	items := v.GetStringSlice(KeyRecent)
	if len(items) > 5 {
		return items[len(items)-5:]
	}
	return items
}

func SetBuildInfo(version, commit, date string) {
	v.Set(KeyVersion, version)
	v.Set(KeyCommit, commit)
	v.Set(KeyDate, date)
}

func IsSet(key string) bool {
	return v.IsSet(key)
}

func Set(key string, value any) {
	v.Set(key, value)
}

func Get(key string) any {
	return v.Get(key)
}

func WriteConfig() error {
	return v.WriteConfig()
}

func EditorCommand() string {
	return v.GetString(KeyEditorCommand)
}

func ServerPort() string {
	return v.GetString(KeyServerPort)
}

func UpdateChannel() string {
	return v.GetString(KeyUpdateChannel)
}

func RegistryDir() string {
	return v.GetString(KeyRegistryDir)
}

func RegistryCachePath() string {
	return filepath.Join(RegistryDir(), "registry.json")
}

func AllSettings() map[string]interface{} {
	return v.AllSettings()
}

func ConfigFileUsed() string {
	return v.ConfigFileUsed()
}

func GitAuthToken() string {
	return v.GetString(KeyGitAuthToken)
}

func TemplateCacheDir() string {
	return v.GetString(KeyTemplatesDir)
}

func RegistryUrl() string {
	return v.GetString(KeyRegistryUrl)
}

func BuildVersion() string {
	return v.GetString(KeyVersion)
}

func BuildDate() string {
	return v.GetString(KeyDate)
}

func BuildCommit() string {
	return v.GetString(KeyCommit)
}
