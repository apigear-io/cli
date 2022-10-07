package git

import (
	"errors"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/apigear-io/cli/pkg/log"
)

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
	err = w.Pull(&git.PullOptions{
		Auth:     auth,
		Progress: os.Stdout,
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return err
	}
	return nil
}
