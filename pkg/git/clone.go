package git

func Clone(repo string, target string) error {
	log.Info().Msgf("clone %s to %s", repo, target)
	out, err := ExecGit([]string{"clone", repo, target}, "")
	if err != nil {
		log.Warn().Msgf("failed to clone %s to %s", repo, target)
		log.Warn().Msgf("clone out: %s", out)
	}
	return err
}
