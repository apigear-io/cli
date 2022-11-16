package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func CheckoutCommit(target string, commit string) error {
	repo, err := git.PlainOpen(target)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	return w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})
}

func CheckoutTag(target, name string) error {
	repo, err := git.PlainOpen(target)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/tags/" + name),
	})
}
