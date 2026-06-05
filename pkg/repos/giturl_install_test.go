package repos

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// makeSourceRepo creates a local git repo to clone from. The first commit
// writes marker.txt="v1" and is tagged "v1.0.0"; the second commit (on the
// default branch) overwrites marker.txt="v2". Checking out the tag must
// therefore yield "v1".
func makeSourceRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	repo, err := gogit.PlainInit(dir, false)
	require.NoError(t, err)
	w, err := repo.Worktree()
	require.NoError(t, err)

	commit := func(content string) {
		require.NoError(t, os.WriteFile(filepath.Join(dir, "marker.txt"), []byte(content), 0644))
		_, err := w.Add("marker.txt")
		require.NoError(t, err)
		_, err = w.Commit("commit "+content, &gogit.CommitOptions{
			Author: &object.Signature{Name: "test", Email: "test@test", When: time.Unix(0, 0)},
		})
		require.NoError(t, err)
	}

	commit("v1")
	head, err := repo.Head()
	require.NoError(t, err)
	_, err = repo.CreateTag("v1.0.0", head.Hash(), nil)
	require.NoError(t, err)
	commit("v2")
	return dir
}

func TestGetOrInstallTemplateFromGitURL_InstallsAndChecksOutRef(t *testing.T) {
	src := makeSourceRepo(t)
	withTempCache(t)

	repoID, err := GetOrInstallTemplateFromGitURL(src, "v1.0.0")
	require.NoError(t, err)
	require.True(t, Cache.Exists(repoID))

	dir, err := Cache.GetTemplateDir(repoID)
	require.NoError(t, err)
	marker, err := os.ReadFile(filepath.Join(dir, "marker.txt"))
	require.NoError(t, err)
	assert.Equal(t, "v1", string(marker))
}

func TestGetOrInstallTemplateFromGitURL_UsesCacheWithoutSource(t *testing.T) {
	src := makeSourceRepo(t)
	withTempCache(t)

	repoID, err := GetOrInstallTemplateFromGitURL(src, "v1.0.0")
	require.NoError(t, err)

	// Removing the source proves the second call is served from the cache
	// and does not touch the network/source again.
	require.NoError(t, os.RemoveAll(src))
	repoID2, err := GetOrInstallTemplateFromGitURL(src, "v1.0.0")
	require.NoError(t, err)
	assert.Equal(t, repoID, repoID2)
}

func TestGetOrInstallTemplateFromGitURL_RequiresVersion(t *testing.T) {
	_, err := GetOrInstallTemplateFromGitURL("https://github.com/me/tpl.git", "")
	require.Error(t, err)
}

// withTempCache points the package-global Cache at a throwaway directory for
// the duration of the test, restoring it afterwards.
func withTempCache(t *testing.T) {
	t.Helper()
	prev := Cache
	Cache = New(t.TempDir())
	t.Cleanup(func() { Cache = prev })
}
