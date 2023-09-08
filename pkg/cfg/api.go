package cfg

import (
	"log"
	"path/filepath"
)

func ConfigDir() string {
	rw.RLock()
	file := v.ConfigFileUsed()
	rw.RUnlock()
	return filepath.Dir(file)
}

// AppendRecentEntry appends a new entry to the list of recent entries
// entries are limited to 5
// the most recent entry is at the beginning of the list
// stores the list in the config file
func AppendRecentEntry(value string) error {
	recent := RecentEntries()
	for i, item := range recent {
		if item == value {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	if len(recent) >= 5 {
		recent = recent[1:]
	}
	// prepend the new value
	recent = append([]string{value}, recent...)

	rw.Lock()
	v.Set(KeyRecent, recent)
	err := v.WriteConfig()
	rw.Unlock()
	if err != nil {
		return err
	}
	return nil
}

// RemoveRecentEntry removes a recent entry from the list
func RemoveRecentEntry(value string) error {
	recent := RecentEntries()
	for i, item := range recent {
		if item == value {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	rw.Lock()
	v.Set(KeyRecent, recent)
	err := v.WriteConfig()
	rw.Unlock()
	if err != nil {
		return err
	}
	return nil
}

// RecentEntries returns the list of recent entries
func RecentEntries() []string {
	rw.RLock()
	items := v.GetStringSlice(KeyRecent)
	rw.RUnlock()
	if len(items) == 0 {
		return []string{}
	}
	if len(items) > 5 {
		return items[len(items)-5:]
	}
	return items
}

func SetBuildInfo(version, commit, date string) {
	rw.Lock()
	v.Set(KeyVersion, version)
	v.Set(KeyCommit, commit)
	v.Set(KeyDate, date)
	err := v.WriteConfig()
	rw.Unlock()
	if err != nil {
		log.Printf("error writing config: %v", err)
	}
}

func IsSet(key string) bool {
	rw.RLock()
	result := v.IsSet(key)
	rw.RUnlock()
	return result
}

func Set(key string, value any) {
	rw.Lock()
	v.Set(key, value)
	rw.Unlock()
}

func Get(key string) any {
	rw.RLock()
	result := v.Get(key)
	rw.RUnlock()
	return result
}

func GetInt(key string) int {
	rw.RLock()
	result := v.GetInt(key)
	rw.RUnlock()
	return result
}

func GetBool(key string) bool {
	rw.RLock()
	result := v.GetBool(key)
	rw.RUnlock()
	return result
}

func SetBool(key string, value bool) {
	rw.Lock()
	v.Set(key, value)
	rw.Unlock()
}

func GetString(key string) string {
	rw.RLock()
	result := v.GetString(key)
	rw.RUnlock()
	return result
}

func WriteConfig() error {
	rw.Lock()
	err := v.WriteConfig()
	rw.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func EditorCommand() string {
	return GetString(KeyEditorCommand)
}

func ServerPort() string {
	return GetString(KeyServerPort)
}

func UpdateChannel() string {
	return GetString(KeyUpdateChannel)
}

func RegistryDir() string {
	return GetString(KeyRegistryDir)
}

func RegistryCachePath() string {
	return filepath.Join(RegistryDir(), "registry.json")
}

func AllSettings() map[string]interface{} {
	rw.RLock()
	result := v.AllSettings()
	rw.RUnlock()
	return result
}

func ConfigFileUsed() string {
	rw.RLock()
	result := v.ConfigFileUsed()
	rw.RUnlock()
	return result
}

func CacheDir() string {
	return GetString(KeyCacheDir)
}

func RegistryUrl() string {
	return GetString(KeyRegistryUrl)
}

func BuildVersion() string {
	return GetString(KeyVersion)
}

func BuildDate() string {
	return GetString(KeyDate)
}

func BuildCommit() string {
	return GetString(KeyCommit)
}
