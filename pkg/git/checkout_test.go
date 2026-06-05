package git

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/require"
)

// makeTestRepo creates a local git repo with two commits, a lightweight tag
// "v1.0.0" on the first commit, and a branch "feature" on the first commit.
// It returns the repo dir and the two commit hashes (in order).
func makeTestRepo(t *testing.T) (string, plumbing.Hash, plumbing.Hash) {
	t.Helper()
	dir := t.TempDir()
	repo, err := gogit.PlainInit(dir, false)
	require.NoError(t, err)
	w, err := repo.Worktree()
	require.NoError(t, err)

	commit := func(file, content string) plumbing.Hash {
		require.NoError(t, os.WriteFile(filepath.Join(dir, file), []byte(content), 0644))
		_, err := w.Add(file)
		require.NoError(t, err)
		h, err := w.Commit("commit "+file, &gogit.CommitOptions{
			Author: &object.Signature{Name: "test", Email: "test@test", When: time.Unix(0, 0)},
		})
		require.NoError(t, err)
		return h
	}

	first := commit("a.txt", "first")

	// lightweight tag and branch on the first commit
	_, err = repo.CreateTag("v1.0.0", first, nil)
	require.NoError(t, err)
	require.NoError(t, repo.Storer.SetReference(
		plumbing.NewHashReference(plumbing.NewBranchReferenceName("feature"), first)))

	second := commit("b.txt", "second")
	return dir, first, second
}

func headHash(t *testing.T, dir string) plumbing.Hash {
	t.Helper()
	repo, err := gogit.PlainOpen(dir)
	require.NoError(t, err)
	head, err := repo.Head()
	require.NoError(t, err)
	return head.Hash()
}

func TestCheckoutRef_Tag(t *testing.T) {
	dir, first, _ := makeTestRepo(t)
	require.NoError(t, CheckoutRef(dir, "v1.0.0"))
	require.Equal(t, first, headHash(t, dir))
}

func TestCheckoutRef_Branch(t *testing.T) {
	dir, first, _ := makeTestRepo(t)
	require.NoError(t, CheckoutRef(dir, "feature"))
	require.Equal(t, first, headHash(t, dir))
}

func TestCheckoutRef_Commit(t *testing.T) {
	dir, first, second := makeTestRepo(t)
	require.NoError(t, CheckoutRef(dir, second.String()))
	require.Equal(t, second, headHash(t, dir))
	// also resolvable by abbreviated hash
	require.NoError(t, CheckoutRef(dir, first.String()[:8]))
	require.Equal(t, first, headHash(t, dir))
}

func TestCheckoutRef_Unknown(t *testing.T) {
	dir, _, _ := makeTestRepo(t)
	require.Error(t, CheckoutRef(dir, "does-not-exist"))
}
