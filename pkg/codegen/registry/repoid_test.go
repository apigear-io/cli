package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRepoID(t *testing.T) {
	// table driven test for testing repo ids (e.g. name@version)
	tests := []struct {
		label    string
		name     string
		version  string
		expected string
	}{
		{"name only", "foo", "", "foo@latest"},
		{"name and version", "foo", "1.0.0", "foo@1.0.0"},
		{"name and latest", "foo", "latest", "foo@latest"},
		{"name and empty", "foo", "", "foo@latest"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := MakeRepoID(tt.name, tt.version)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSplitRepoID(t *testing.T) {
	// table driven test for testing repo ids (e.g. name@version)
	tests := []struct {
		label           string
		name            string
		expectedName    string
		expectedVersion string
	}{
		{"name only", "foo", "foo", "latest"},
		{"name and version", "foo@1.0.0", "foo", "1.0.0"},
		{"name and latest", "foo@latest", "foo", "latest"},
		{"name and empty", "foo@", "foo", "latest"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actualName, actualVersion := SplitRepoID(tt.name)
			assert.Equal(t, tt.expectedName, actualName)
			assert.Equal(t, tt.expectedVersion, actualVersion)
		})
	}
}

func TestNameFromRepoID(t *testing.T) {
	// table driven test for testing repo ids (e.g. name@version)
	tests := []struct {
		label    string
		name     string
		expected string
	}{
		{"name only", "foo", "foo"},
		{"name and version", "foo@1.0.0", "foo"},
		{"name and latest", "foo@latest", "foo"},
		{"name and empty", "foo@", "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := NameFromRepoID(tt.name)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestVersionFromRepoID(t *testing.T) {
	// table driven test for testing repo ids (e.g. name@version)
	tests := []struct {
		label    string
		name     string
		expected string
	}{
		{"name only", "foo", "latest"},
		{"name and version", "foo@1.0.0", "1.0.0"},
		{"name and latest", "foo@latest", "latest"},
		{"name and empty", "foo@", "latest"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := VersionFromRepoID(tt.name)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestEnsureRepoID(t *testing.T) {
	// table driven test for testing repo ids (e.g. name@version)
	tests := []struct {
		label    string
		name     string
		expected string
	}{
		{"name only", "foo", "foo@latest"},
		{"name and version", "foo@1.0.0", "foo@1.0.0"},
		{"name and latest", "foo@latest", "foo@latest"},
		{"name and empty", "foo@", "foo@latest"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := EnsureRepoID(tt.name)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIsRepoID(t *testing.T) {
	tests := []struct {
		label    string
		name     string
		expected bool
	}{
		{"name only", "foo", false},
		{"name and version", "foo@1.0.0", true},
		{"name and latest", "foo@latest", true},
		{"name with empty version", "foo@", true},
		{"complex name", "github/user/repo", false},
		{"complex name with version", "github/user/repo@1.0.0", true},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := IsRepoID(tt.name)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMakeRepoIDEdgeCases(t *testing.T) {
	tests := []struct {
		label    string
		name     string
		version  string
		expected string
	}{
		{"name with @ and empty version", "foo@1.0.0", "", "foo@latest"},
		{"name with @ and new version", "foo@1.0.0", "2.0.0", "foo@2.0.0"},
		{"complex name", "github.com/user/repo", "1.0.0", "github.com/user/repo@1.0.0"},
		{"name with special chars", "my-template_v1", "1.0.0", "my-template_v1@1.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := MakeRepoID(tt.name, tt.version)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestVersionFromRepoIDEdgeCases(t *testing.T) {
	tests := []struct {
		label    string
		input    string
		expected string
	}{
		{"v prefix", "foo@v1.0.0", "v1.0.0"},
		{"semver with patch", "foo@1.2.3", "1.2.3"},
		{"semver with prerelease", "foo@1.0.0-alpha.1", "1.0.0-alpha.1"},
		{"semver with build", "foo@1.0.0+build.123", "1.0.0+build.123"},
		{"tag name", "foo@release-1.0", "release-1.0"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := VersionFromRepoID(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNameFromRepoIDEdgeCases(t *testing.T) {
	tests := []struct {
		label    string
		input    string
		expected string
	}{
		{"org/repo format", "github/user/repo@1.0.0", "github/user/repo"},
		{"with dots", "my.template@1.0.0", "my.template"},
		{"with hyphens", "my-template@1.0.0", "my-template"},
		{"with underscores", "my_template@1.0.0", "my_template"},
		{"complex path", "a/b/c/d@1.0.0", "a/b/c/d"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			actual := NameFromRepoID(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
