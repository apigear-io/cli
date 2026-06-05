package git

import (
	"fmt"

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

// CheckoutRef checks out the given ref, which may be a tag, a branch
// (local or remote-tracking under origin/), or a commit hash. The ref is
// resolved to a commit hash and checked out as a detached HEAD.
func CheckoutRef(target, ref string) error {
	repo, err := git.PlainOpen(target)
	if err != nil {
		return err
	}
	hash, err := resolveRef(repo, ref)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	return w.Checkout(&git.CheckoutOptions{Hash: *hash, Force: true})
}

// resolveRef resolves ref to a commit hash, trying in order: tag, local
// branch, remote-tracking branch (origin/<ref>), then a generic revision
// (which covers full and abbreviated commit hashes).
func resolveRef(repo *git.Repository, ref string) (*plumbing.Hash, error) {
	if t, err := repo.Reference(plumbing.NewTagReferenceName(ref), true); err == nil {
		h := t.Hash()
		// annotated tags point at a tag object; dereference to its commit
		if to, err := repo.TagObject(h); err == nil {
			if c, err := to.Commit(); err == nil {
				return &c.Hash, nil
			}
		}
		return &h, nil
	}
	if b, err := repo.Reference(plumbing.NewBranchReferenceName(ref), true); err == nil {
		h := b.Hash()
		return &h, nil
	}
	if rb, err := repo.Reference(plumbing.NewRemoteReferenceName("origin", ref), true); err == nil {
		h := rb.Hash()
		return &h, nil
	}
	if h, err := repo.ResolveRevision(plumbing.Revision(ref)); err == nil {
		return h, nil
	}
	return nil, fmt.Errorf("could not resolve git ref %q", ref)
}
