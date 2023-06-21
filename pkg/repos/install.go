package repos

import (
	"fmt"
)

// InstallTemplateFromFQN tries to install a template
// from a fully qualified name (e.g. name@version)
func InstallTemplateFromRepoID(repoID string) error {
	version := VersionFromRepoID(repoID)
	info, err := Registry.Get(repoID)
	if err != nil {
		return err
	}
	if version == "latest" {
		version = info.Latest.Name
		if version == "" {
			return fmt.Errorf("no version found for template: %s", repoID)
		}
		repoID = MakeRepoID(info.Name, version)
		log.Info().Msgf("use latest version %s ", repoID)
	}
	if Cache.Exists(repoID) {
		log.Info().Msgf("template %s already installed", repoID)
		return nil
	}
	url := info.Git
	log.Info().Msgf("installing template %s@%s from %s", repoID, version, url)
	_, err = Cache.Install(url, version)
	if err != nil {
		return err
	}
	return nil
}
