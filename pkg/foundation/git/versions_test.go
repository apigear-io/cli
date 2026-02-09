package git

import (
	"sort"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionCollectionLen(t *testing.T) {
	t.Run("returns zero for empty collection", func(t *testing.T) {
		vc := VersionCollection{}
		assert.Equal(t, 0, vc.Len())
	})

	t.Run("returns correct length for collection", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")
		vc := VersionCollection{
			{Name: "v1.0.0", Version: v1},
			{Name: "v2.0.0", Version: v2},
		}
		assert.Equal(t, 2, vc.Len())
	})
}

func TestVersionCollectionLess(t *testing.T) {
	v1, _ := semver.NewVersion("1.0.0")
	v2, _ := semver.NewVersion("2.0.0")
	v3, _ := semver.NewVersion("1.5.0")

	vc := VersionCollection{
		{Name: "v1.0.0", Version: v1},
		{Name: "v2.0.0", Version: v2},
		{Name: "v1.5.0", Version: v3},
	}

	t.Run("v1.0.0 is less than v2.0.0", func(t *testing.T) {
		assert.True(t, vc.Less(0, 1))
	})

	t.Run("v2.0.0 is not less than v1.0.0", func(t *testing.T) {
		assert.False(t, vc.Less(1, 0))
	})

	t.Run("v1.0.0 is less than v1.5.0", func(t *testing.T) {
		assert.True(t, vc.Less(0, 2))
	})

	t.Run("v1.5.0 is less than v2.0.0", func(t *testing.T) {
		assert.True(t, vc.Less(2, 1))
	})
}

func TestVersionCollectionSwap(t *testing.T) {
	v1, _ := semver.NewVersion("1.0.0")
	v2, _ := semver.NewVersion("2.0.0")

	vc := VersionCollection{
		{Name: "v1.0.0", Version: v1},
		{Name: "v2.0.0", Version: v2},
	}

	t.Run("swaps elements at indices", func(t *testing.T) {
		assert.Equal(t, "v1.0.0", vc[0].Name)
		assert.Equal(t, "v2.0.0", vc[1].Name)

		vc.Swap(0, 1)

		assert.Equal(t, "v2.0.0", vc[0].Name)
		assert.Equal(t, "v1.0.0", vc[1].Name)
	})
}

func TestVersionCollectionSorting(t *testing.T) {
	v1, _ := semver.NewVersion("1.0.0")
	v2, _ := semver.NewVersion("2.0.0")
	v3, _ := semver.NewVersion("1.5.0")
	v4, _ := semver.NewVersion("0.9.0")

	t.Run("sorts versions in ascending order", func(t *testing.T) {
		vc := VersionCollection{
			{Name: "v2.0.0", Version: v2},
			{Name: "v1.0.0", Version: v1},
			{Name: "v1.5.0", Version: v3},
			{Name: "v0.9.0", Version: v4},
		}

		sort.Sort(vc)

		assert.Equal(t, "v0.9.0", vc[0].Name)
		assert.Equal(t, "v1.0.0", vc[1].Name)
		assert.Equal(t, "v1.5.0", vc[2].Name)
		assert.Equal(t, "v2.0.0", vc[3].Name)
	})
}

func TestVersionCollectionLatest(t *testing.T) {
	t.Run("returns empty VersionInfo for empty collection", func(t *testing.T) {
		vc := VersionCollection{}
		latest := vc.Latest()
		assert.Equal(t, VersionInfo{}, latest)
		assert.Empty(t, latest.Name)
	})

	t.Run("returns single version for collection with one element", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		vc := VersionCollection{
			{Name: "v1.0.0", SHA: "abc123", Version: v1},
		}
		latest := vc.Latest()
		assert.Equal(t, "v1.0.0", latest.Name)
		assert.Equal(t, "abc123", latest.SHA)
	})

	t.Run("returns latest version from unsorted collection", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")
		v3, _ := semver.NewVersion("1.5.0")

		vc := VersionCollection{
			{Name: "v1.0.0", SHA: "abc123", Version: v1},
			{Name: "v2.0.0", SHA: "def456", Version: v2},
			{Name: "v1.5.0", SHA: "ghi789", Version: v3},
		}

		latest := vc.Latest()
		assert.Equal(t, "v2.0.0", latest.Name)
		assert.Equal(t, "def456", latest.SHA)
	})

	t.Run("handles pre-release versions", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0-beta.1")
		v3, _ := semver.NewVersion("1.5.0")

		vc := VersionCollection{
			{Name: "v1.0.0", Version: v1},
			{Name: "v2.0.0-beta.1", Version: v2},
			{Name: "v1.5.0", Version: v3},
		}

		latest := vc.Latest()
		// Pre-release versions are considered less than release versions
		// So v2.0.0-beta.1 > v1.5.0 > v1.0.0
		assert.Equal(t, "v2.0.0-beta.1", latest.Name)
	})
}

func TestVersionCollectionAsList(t *testing.T) {
	t.Run("returns empty list for empty collection", func(t *testing.T) {
		vc := VersionCollection{}
		list := vc.AsList()
		assert.Empty(t, list)
	})

	t.Run("returns list of version names", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")
		v3, _ := semver.NewVersion("1.5.0")

		vc := VersionCollection{
			{Name: "v1.0.0", Version: v1},
			{Name: "v2.0.0", Version: v2},
			{Name: "v1.5.0", Version: v3},
		}

		list := vc.AsList()
		assert.Len(t, list, 3)
		assert.Contains(t, list, "v1.0.0")
		assert.Contains(t, list, "v2.0.0")
		assert.Contains(t, list, "v1.5.0")
	})

	t.Run("maintains order of collection", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")

		vc := VersionCollection{
			{Name: "v2.0.0", Version: v2},
			{Name: "v1.0.0", Version: v1},
		}

		list := vc.AsList()
		assert.Equal(t, []string{"v2.0.0", "v1.0.0"}, list)
	})
}

func TestVersionCollectionString(t *testing.T) {
	t.Run("returns empty string for empty collection", func(t *testing.T) {
		vc := VersionCollection{}
		result := vc.String()
		assert.Equal(t, "", result)
	})

	t.Run("returns comma-separated list of version names", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("2.0.0")

		vc := VersionCollection{
			{Name: "v1.0.0", Version: v1},
			{Name: "v2.0.0", Version: v2},
		}

		result := vc.String()
		assert.Contains(t, result, "v1.0.0")
		assert.Contains(t, result, "v2.0.0")
		assert.Contains(t, result, ", ")
	})

	t.Run("includes trailing comma and space", func(t *testing.T) {
		v1, _ := semver.NewVersion("1.0.0")

		vc := VersionCollection{
			{Name: "v1.0.0", Version: v1},
		}

		result := vc.String()
		assert.Equal(t, "v1.0.0, ", result)
	})
}

func TestVersionInfo(t *testing.T) {
	t.Run("creates VersionInfo with all fields", func(t *testing.T) {
		v, err := semver.NewVersion("1.2.3")
		require.NoError(t, err)

		info := VersionInfo{
			Name:    "v1.2.3",
			SHA:     "abc123def456",
			Version: v,
		}

		assert.Equal(t, "v1.2.3", info.Name)
		assert.Equal(t, "abc123def456", info.SHA)
		assert.NotNil(t, info.Version)
		assert.Equal(t, uint64(1), info.Version.Major())
		assert.Equal(t, uint64(2), info.Version.Minor())
		assert.Equal(t, uint64(3), info.Version.Patch())
	})
}
