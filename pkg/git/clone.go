package git

import (
	"errors"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/go-git/go-git/v5"
)

func Clone(src string, dst string) error {
	log.Debug().Msgf("clone %s %s", src, dst)
	_, err := git.PlainClone(dst, false, &git.CloneOptions{
		URL:  src,
		Auth: auth(),
	})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}
	return err
}

func CloneOrPull(src string, dst string) error {
	log.Debug().Msgf("clone or pull %s %s", src, dst)
	if helper.IsDir(dst) {
		return Pull(dst)
	}
	return Clone(src, dst)
}

func Pull(dst string) error {
	log.Debug().Msgf("pull %s", dst)
	repo, err := git.PlainOpen(dst)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{Auth: auth()})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return err
	}
	return nil
}
