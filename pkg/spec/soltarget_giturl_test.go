package spec

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/repos"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// makeTemplateRepo creates a local git repo shaped like a template package:
// a templates/ directory and a rules.yaml file. The first commit (tagged
// "v1.0.0") writes marker.txt="v1"; a second commit on the default branch
// overwrites it with "v2", so checking out the tag must yield "v1".
func makeTemplateRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	repo, err := gogit.PlainInit(dir, false)
	require.NoError(t, err)
	w, err := repo.Worktree()
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "templates"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "templates", "keep.txt"), []byte("x"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "rules.yaml"), []byte("features: []\n"), 0644))

	commit := func(content string) {
		require.NoError(t, os.WriteFile(filepath.Join(dir, "marker.txt"), []byte(content), 0644))
		_, err := w.Add(".")
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

// withHermeticRepos points the repos package globals at throwaway directories
// so resolution never touches the real cache, registry, or the network. The
// registry is seeded empty so the registry lookup fails cleanly (offline),
// letting the direct git-url path take over. It returns the temp cache dir.
func withHermeticRepos(t *testing.T) string {
	t.Helper()
	prevCache := repos.Cache
	prevReg := repos.Registry
	cacheDir := t.TempDir()
	regDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(regDir, "registry.json"), []byte(`{"entries":[]}`), 0644))
	repos.Cache = repos.New(cacheDir)
	repos.Registry = repos.NewRegistry(regDir, "")
	t.Cleanup(func() {
		repos.Cache = prevCache
		repos.Registry = prevReg
	})
	return cacheDir
}

func TestSolutionTargetValidate_GitURLTemplate(t *testing.T) {
	src := makeTemplateRepo(t)
	withHermeticRepos(t)

	doc := &SolutionDoc{
		Version: "1.0.0",
		Name:    "solution",
		RootDir: t.TempDir(),
	}
	target := &SolutionTarget{
		Name:     "go-sdk",
		Template: "file://" + src + "@v1.0.0",
		Output:   "./out",
	}

	require.NoError(t, target.Validate(doc))

	// the template field is rewritten to the resolved cache repo id
	assert.Contains(t, target.Template, "@v1.0.0")
	// resolved paths point into the cache and exist
	assert.True(t, isDir(target.TemplateDir))
	assert.True(t, isDir(target.TemplatesDir))
	assert.True(t, isFile(target.RulesFile))
	// the tagged ref was checked out (v1, not the later v2)
	marker, err := os.ReadFile(filepath.Join(target.TemplateDir, "marker.txt"))
	require.NoError(t, err)
	assert.Equal(t, "v1", string(marker))
}

// An scp-style git url (git@host:repo.git@ref) has two '@'. It must be routed
// straight to the git-url resolver and must never reach the registry repo-id
// parser, which os.Exit()s on such input. We pre-seed the cache so the git-url
// resolver short-circuits without any network access.
func TestSolutionTargetValidate_ScpGitURLDoesNotCrash(t *testing.T) {
	cacheDir := withHermeticRepos(t)

	// pre-seed the cache entry the git-url resolver will look for
	repoID := "github.com/me/tpl@main"
	tplDir := filepath.Join(cacheDir, "github.com/me/tpl@main")
	require.NoError(t, os.MkdirAll(filepath.Join(tplDir, "templates"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(tplDir, "rules.yaml"), []byte("features: []\n"), 0644))

	doc := &SolutionDoc{Version: "1.0.0", Name: "solution", RootDir: t.TempDir()}
	target := &SolutionTarget{
		Name:     "go-sdk",
		Template: "git@github.com:me/tpl.git@main",
		Output:   "./out",
	}

	require.NoError(t, target.Validate(doc))
	assert.Equal(t, repoID, target.Template)
}

func isDir(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

func isFile(p string) bool {
	info, err := os.Stat(p)
	return err == nil && !info.IsDir()
}
