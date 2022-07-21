package git

func Clone(repo string, target string) error {
	log.Infof("clone %s to %s", repo, target)
	out, err := ExecGit([]string{"clone", repo, target}, "")
	if err != nil {
		log.Warnf("failed to clone %s to %s", repo, target)
		log.Warnf("clone out: %s", out)
	}
	return err
}
