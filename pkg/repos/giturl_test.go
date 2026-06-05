package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGitURL(t *testing.T) {
	tests := []struct {
		label    string
		input    string
		expected bool
	}{
		{"https", "https://github.com/me/tpl.git", true},
		{"http", "http://github.com/me/tpl.git", true},
		{"ssh scheme", "ssh://git@github.com/me/tpl.git", true},
		{"git scheme", "git://github.com/me/tpl.git", true},
		{"file scheme", "file:///tmp/tpl.git", true},
		{"scp style", "git@github.com:me/tpl.git", true},
		{"https with version", "https://github.com/me/tpl.git@v1.2.0", true},
		{"registry id", "apigear-io/template-go", false},
		{"registry id with version", "apigear-io/template-go@1.2.3", false},
		{"relative path", "./tpl", false},
		{"parent path", "../templates/foo", false},
		{"bare name", "foo", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsGitURL(tt.input))
		})
	}
}

func TestSplitGitURLVersion(t *testing.T) {
	tests := []struct {
		label           string
		input           string
		expectedURL     string
		expectedVersion string
	}{
		{"https with tag", "https://github.com/me/tpl.git@v1.2.0", "https://github.com/me/tpl.git", "v1.2.0"},
		{"https no version", "https://github.com/me/tpl.git", "https://github.com/me/tpl.git", ""},
		{"scp with branch", "git@github.com:me/tpl.git@main", "git@github.com:me/tpl.git", "main"},
		{"scp no version", "git@github.com:me/tpl.git", "git@github.com:me/tpl.git", ""},
		{"https with credentials and version", "https://user@host.com/me/tpl.git@v1", "https://user@host.com/me/tpl.git", "v1"},
		{"https with credentials no version", "https://user@host.com/me/tpl.git", "https://user@host.com/me/tpl.git", ""},
		{"branch with slash", "https://github.com/me/tpl.git@release/1.0", "https://github.com/me/tpl.git", "release/1.0"},
		{"ssh scheme with commit", "ssh://git@github.com/me/tpl.git@abc1234", "ssh://git@github.com/me/tpl.git", "abc1234"},
		{"git scheme with version", "git://github.com/me/tpl.git@v1", "git://github.com/me/tpl.git", "v1"},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			url, version := SplitGitURLVersion(tt.input)
			assert.Equal(t, tt.expectedURL, url)
			assert.Equal(t, tt.expectedVersion, version)
		})
	}
}
