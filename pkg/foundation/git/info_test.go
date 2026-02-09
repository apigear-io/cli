package git

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestRepoInfoFQN(t *testing.T) {
	t.Run("returns name with version when version is set", func(t *testing.T) {
		v, _ := semver.NewVersion("1.2.3")
		info := &RepoInfo{
			Name: "test-repo",
			Version: VersionInfo{
				Name:    "v1.2.3",
				Version: v,
			},
		}

		fqn := info.FQN()
		assert.Equal(t, "test-repo@v1.2.3", fqn)
	})

	t.Run("returns only name when version is not set", func(t *testing.T) {
		info := &RepoInfo{
			Name:    "test-repo",
			Version: VersionInfo{},
		}

		fqn := info.FQN()
		assert.Equal(t, "test-repo", fqn)
	})

	t.Run("returns only name when version name is empty", func(t *testing.T) {
		info := &RepoInfo{
			Name: "test-repo",
			Version: VersionInfo{
				Name: "",
			},
		}

		fqn := info.FQN()
		assert.Equal(t, "test-repo", fqn)
	})
}

func TestRepoInfoVersionName(t *testing.T) {
	t.Run("returns version name when set", func(t *testing.T) {
		v, _ := semver.NewVersion("1.2.3")
		info := &RepoInfo{
			Version: VersionInfo{
				Name:    "v1.2.3",
				Version: v,
			},
			Commit: "abc123def456",
		}

		versionName := info.VersionName()
		assert.Equal(t, "v1.2.3", versionName)
	})

	t.Run("returns commit hash when version name is empty", func(t *testing.T) {
		info := &RepoInfo{
			Version: VersionInfo{
				Name: "",
			},
			Commit: "abc123def456",
		}

		versionName := info.VersionName()
		assert.Equal(t, "abc123def456", versionName)
	})

	t.Run("returns commit hash when version is not set", func(t *testing.T) {
		info := &RepoInfo{
			Commit: "abc123def456",
		}

		versionName := info.VersionName()
		assert.Equal(t, "abc123def456", versionName)
	})

	t.Run("returns empty string when both version and commit are empty", func(t *testing.T) {
		info := &RepoInfo{
			Version: VersionInfo{},
			Commit:  "",
		}

		versionName := info.VersionName()
		assert.Equal(t, "", versionName)
	})
}

func TestSortRepoInfo(t *testing.T) {
	t.Run("sorts repos by name alphabetically", func(t *testing.T) {
		infos := []*RepoInfo{
			{Name: "zebra-repo"},
			{Name: "alpha-repo"},
			{Name: "beta-repo"},
		}

		SortRepoInfo(infos)

		assert.Equal(t, "alpha-repo", infos[0].Name)
		assert.Equal(t, "beta-repo", infos[1].Name)
		assert.Equal(t, "zebra-repo", infos[2].Name)
	})

	t.Run("sorts versions within each repo in descending order", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")
		v3, _ := semver.NewVersion("1.5.0")

		infos := []*RepoInfo{
			{
				Name: "test-repo",
				Versions: VersionCollection{
					{Name: "v1.0.0", Version: v1},
					{Name: "v2.0.0", Version: v2},
					{Name: "v1.5.0", Version: v3},
				},
			},
		}

		SortRepoInfo(infos)

		// Versions should be sorted in descending order (latest first)
		assert.Equal(t, "v2.0.0", infos[0].Versions[0].Name)
		assert.Equal(t, "v1.5.0", infos[0].Versions[1].Name)
		assert.Equal(t, "v1.0.0", infos[0].Versions[2].Name)
	})

	t.Run("sorts both repos and versions", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")
		v3, _ := semver.NewVersion("3.0.0")
		v4, _ := semver.NewVersion("4.0.0")

		infos := []*RepoInfo{
			{
				Name: "zebra-repo",
				Versions: VersionCollection{
					{Name: "v1.0.0", Version: v1},
					{Name: "v2.0.0", Version: v2},
				},
			},
			{
				Name: "alpha-repo",
				Versions: VersionCollection{
					{Name: "v3.0.0", Version: v3},
					{Name: "v4.0.0", Version: v4},
				},
			},
		}

		SortRepoInfo(infos)

		// Repos sorted alphabetically
		assert.Equal(t, "alpha-repo", infos[0].Name)
		assert.Equal(t, "zebra-repo", infos[1].Name)

		// Versions sorted descending
		assert.Equal(t, "v4.0.0", infos[0].Versions[0].Name)
		assert.Equal(t, "v3.0.0", infos[0].Versions[1].Name)
		assert.Equal(t, "v2.0.0", infos[1].Versions[0].Name)
		assert.Equal(t, "v1.0.0", infos[1].Versions[1].Name)
	})

	t.Run("handles empty repo list", func(t *testing.T) {
		infos := []*RepoInfo{}
		SortRepoInfo(infos)
		assert.Empty(t, infos)
	})

	t.Run("handles repo with no versions", func(t *testing.T) {
		infos := []*RepoInfo{
			{Name: "test-repo", Versions: VersionCollection{}},
		}

		SortRepoInfo(infos)

		assert.Equal(t, "test-repo", infos[0].Name)
		assert.Empty(t, infos[0].Versions)
	})

	t.Run("handles single repo", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")

		infos := []*RepoInfo{
			{
				Name: "single-repo",
				Versions: VersionCollection{
					{Name: "v1.0.0", Version: v1},
				},
			},
		}

		SortRepoInfo(infos)

		assert.Equal(t, "single-repo", infos[0].Name)
		assert.Len(t, infos[0].Versions, 1)
	})
}
