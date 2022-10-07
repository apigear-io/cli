package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/apigear-io/cli/pkg/log"
)

func CheckoutCommit(target string, commit string) error {
	log.Info().Msgf("checkout %s", target)
	repo, err := git.PlainOpen(target)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})
	if err != nil {
		return err
	}
	return nil
}
