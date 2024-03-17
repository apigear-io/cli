package repos

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
