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

var c *config

type config struct {
	*viper.Viper
}

func New(v *viper.Viper) *config {
	if v == nil {
		v = viper.New()
	}
	return &config{
		Viper: v,
	}
}

// RecentEntries returns the list of recent entries
func (c *config) RecentEntries() []string {
	items := c.GetStringSlice(KeyRecent)
	if len(items) > 5 {
		return items[len(items)-5:]
	}
	return items
}

// AppendRecentEntry appends a new entry to the list of recent entries
// entries are limited to 5
// the most recent entry is at the beginning of the list
// stores the list in the config file
func (c *config) AppendRecentEntry(file string) error {
	recent := c.RecentEntries()
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
	c.Set(KeyRecent, recent)
	return c.WriteConfig()
}

// RemoveRecentEntry removes a recent entry from the list
func (c *config) RemoveRecentEntry(d string) error {
	recent := c.RecentEntries()
	for i, f := range recent {
		if f == d {
			recent = append(recent[:i], recent[i+1:]...)
			break
		}
	}
	c.Set(KeyRecent, recent)
	return c.WriteConfig()
}

func (c *config) EditorCommand() string {
	return c.GetString(KeyEditorCommand)
}

func (c *config) ServerPort() string {
	return c.GetString(KeyServerPort)
}

func (c *config) UpdateChannel() string {
	return c.GetString(KeyUpdateChannel)
}

func (c *config) RegistryCacheDir() string {
	return c.GetString(KeyRegistryDir)
}

func (c *config) RegistryCacheFile() string {
	return filepath.Join(c.RegistryCacheDir(), "registry.json")
}

func (c *config) Get(key string) string {
	return c.GetString(key)
}

func (c *config) GitAuthToken() string {
	return c.GetString(KeyGitAuthToken)
}

func (c *config) TemplateCacheDir() string {
	return c.GetString(KeyTemplatesDir)
}

func (c *config) RegistryUrl() string {
	return c.GetString(KeyRegistryUrl)
}

// recent entries
func AppendRecentEntry(file string) error {
	return c.AppendRecentEntry(file)
}

func RemoveRecentEntry(d string) error {
	return c.RemoveRecentEntry(d)
}

func RecentEntries() []string {
	return c.RecentEntries()
}

func SetBuildInfo(version, commit, date string) {
	c.Set(KeyVersion, version)
	c.Set(KeyCommit, commit)
	c.Set(KeyDate, date)
}

func IsSet(key string) bool {
	return c.IsSet(key)
}

func Set(key string, value any) {
	c.Set(key, value)
}

func Get(key string) any {
	return c.Get(key)
}

func WriteConfig() error {
	return c.WriteConfig()
}

func EditorCommand() string {
	return c.EditorCommand()
}

func ServerPort() string {
	return c.ServerPort()
}

func UpdateChannel() string {
	return c.UpdateChannel()
}

func RegistryDir() string {
	return c.RegistryCacheDir()
}

func RegistryCachePath() string {
	return c.RegistryCacheFile()
}

func AllSettings() map[string]interface{} {
	return c.AllSettings()
}

func ConfigFileUsed() string {
	return c.ConfigFileUsed()
}

func GitAuthToken() string {
	return c.GitAuthToken()
}

func TemplateCacheDir() string {
	return c.TemplateCacheDir()
}

func RegistryUrl() string {
	return c.RegistryUrl()
}

func BuildVersion() string {
	return c.GetString(KeyVersion)
}

func BuildDate() string {
	return c.GetString(KeyDate)
}

func BuildCommit() string {
	return c.GetString(KeyCommit)
}
