package repos

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
