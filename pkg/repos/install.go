package repos

import "fmt"

// GetOrInstallTemplateFromGitURL resolves a template from a direct git url and
// version (a tag, branch or commit). The repo is cloned into the cache and
// checked out at the given version. If a matching clone already exists in the
// cache it is reused without touching the network.
func GetOrInstallTemplateFromGitURL(url string, version string) (string, error) {
	if version == "" {
		return "", fmt.Errorf("template git url %q requires a version (url@version)", url)
	}
	name, err := RepoNameFromGitURL(url)
	if err != nil {
		return "", err
	}
	repoID := MakeRepoID(name, version)
	if Cache.Exists(repoID) {
		log.Info().Msgf("template %s already installed", repoID)
		return repoID, nil
	}
	log.Info().Msgf("installing template %s from %s", repoID, url)
	return Cache.InstallRef(url, version)
}

// InstallTemplateFromFQN tries to install a template
// from a fully qualified name (e.g. name@version)
func GetOrInstallTemplateFromRepoID(repoID string) (string, error) {
	log.Info().Msgf("installing template %s", repoID)
	fixedRepoId, err := Registry.FixRepoId(repoID)
	if err != nil {
		return "", err
	}
	if Cache.Exists(fixedRepoId) {
		log.Info().Msgf("template %s already installed", fixedRepoId)
		return fixedRepoId, nil
	}
	info, err := Registry.Get(fixedRepoId)
	if err != nil {
		return "", err
	}
	url := info.Git
	log.Info().Msgf("installing template %s from %s", fixedRepoId, url)
	version := VersionFromRepoID(fixedRepoId)
	_, err = Cache.Install(url, version)
	if err != nil {
		return "", err
	}
	return fixedRepoId, nil
}
