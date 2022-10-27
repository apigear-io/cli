package up

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/creativeprojects/go-selfupdate"
)

// Updater checks a github repository for new releases
// and updates the current executable
// It is a wrapper around github.com/creativeprojects/go-selfupdate
type Updater struct {
	repo    string
	version string
	updater *selfupdate.Updater
}

// NewUpdater creates a new updater for a github repository
func NewUpdater(repo string, version string) (*Updater, error) {
	source, err := selfupdate.NewGitHubSource(selfupdate.GitHubConfig{
		APIToken: cfg.GitPublicToken(),
	})
	if err != nil {
		return nil, err
	}
	up, err := selfupdate.NewUpdater(selfupdate.Config{
		Validator: &selfupdate.ChecksumValidator{
			UniqueFilename: "checksums.txt",
		},
		Source: source,
	})
	if err != nil {
		return nil, err
	}
	return &Updater{
		repo:    repo,
		version: version,
		updater: up,
	}, nil
}

// Check checks for a new release
// returns a release if there is one, or nil if there is no new release
func (u *Updater) Check() (*selfupdate.Release, error) {
	log.Info().Msgf("check for updates: %s", u.repo)
	latest, found, err := u.updater.DetectLatest(u.repo)
	if err != nil {
		return nil, err
	}
	if found && latest.GreaterThan(u.version) {
		return latest, nil
	}
	return nil, nil
}

// Update updates the current executable to the latest release
func (u *Updater) Update(release *selfupdate.Release) error {
	// get the current executable path
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	// exe might be a symlink, so we need to resolve it
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return err
	}
	if !helper.IsFile(exe) {
		return fmt.Errorf("executable not found: %s", exe)
	}
	return u.updater.UpdateTo(release, exe)
}
