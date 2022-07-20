package git

func Pull(target string) (string, error) {
	return ExecGit([]string{"pull"}, target)
}
