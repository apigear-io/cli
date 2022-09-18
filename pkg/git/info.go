package git

func LastCommit(target string) (string, error) {
	return ExecGit([]string{"log", "-1", "--pretty=%H"}, target)
}

func RemoteUrl(target string) (string, error) {
	return ExecGit([]string{"config", "--get", "remote.origin.url"}, target)
}
