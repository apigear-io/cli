package git

import (
	"github.com/go-git/go-git/v5"
)

func RepoLastCommit(src string) (string, error) {
	repo, err := git.PlainOpen(src)
	if err != nil {
		return "", err
	}
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	return ref.Hash().String(), nil
}

func RepoRemoteUrl(src string) (string, error) {
	repo, err := git.PlainOpen(src)
	if err != nil {
		return "", err
	}
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", err
	}
	return remote.Config().URLs[0], nil
}
