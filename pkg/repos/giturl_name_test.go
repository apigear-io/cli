package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepoNameFromGitURL(t *testing.T) {
	tests := []struct {
		label    string
		input    string
		expected string
	}{
		{"https github", "https://github.com/me/tpl.git", "github.com/me/tpl"},
		{"scp github", "git@github.com:me/tpl.git", "github.com/me/tpl"},
		{"ssh scheme github", "ssh://git@github.com/me/tpl.git", "github.com/me/tpl"},
		{"self hosted", "https://git.example.com/me/tpl.git", "git.example.com/me/tpl"},
		{"no dot git suffix", "https://github.com/me/tpl", "github.com/me/tpl"},
		{"nested path", "https://github.com/me/sub/tpl.git", "github.com/me/sub/tpl"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			name, err := RepoNameFromGitURL(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, name)
		})
	}
}

// https and scp forms of the same repo must produce the same cache key so they
// share a cached clone.
func TestRepoNameFromGitURL_ProtocolEquivalence(t *testing.T) {
	https, err := RepoNameFromGitURL("https://github.com/me/tpl.git")
	require.NoError(t, err)
	scp, err := RepoNameFromGitURL("git@github.com:me/tpl.git")
	require.NoError(t, err)
	assert.Equal(t, https, scp)
}
