package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseAsUrl(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid HTTPS URL",
			url:     "https://github.com/apigear-io/cli.git",
			wantErr: false,
		},
		{
			name:    "valid SSH URL",
			url:     "git@github.com:apigear-io/cli.git",
			wantErr: false,
		},
		{
			name:    "valid git:// URL",
			url:     "git://github.com/apigear-io/cli.git",
			wantErr: false,
		},
		{
			name:    "valid file:// URL",
			url:     "file:///path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "simple HTTPS without .git",
			url:     "https://github.com/apigear-io/cli",
			wantErr: false,
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: false, // Empty string parses as file:// URL
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAsUrl(tt.url)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestIsValidGitUrl(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		valid bool
	}{
		{
			name:  "valid HTTPS URL",
			url:   "https://github.com/apigear-io/cli.git",
			valid: true,
		},
		{
			name:  "valid SSH URL",
			url:   "ssh://git@github.com/apigear-io/cli.git",
			valid: true,
		},
		{
			name:  "valid git:// URL",
			url:   "git://github.com/apigear-io/cli.git",
			valid: true,
		},
		{
			name:  "valid file:// URL",
			url:   "file:///path/to/repo.git",
			valid: true,
		},
		{
			name:  "SSH URL with colon notation",
			url:   "github.com:apigear-io/cli.git",
			valid: false, // This format is not recognized by ParseTransport
		},
		{
			name:  "simple HTTPS without .git",
			url:   "https://github.com/apigear-io/cli",
			valid: true,
		},
		{
			name:  "empty URL",
			url:   "",
			valid: false,
		},
		{
			name:  "invalid URL",
			url:   "not a valid url",
			valid: false,
		},
		{
			name:  "just a path",
			url:   "/path/to/repo",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidGitUrl(tt.url)
			assert.Equal(t, tt.valid, result)
		})
	}
}

func TestParseAsVcsUrl(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		wantErr  bool
		wantHost string
		wantRepo string
	}{
		{
			name:     "GitHub HTTPS URL",
			url:      "https://github.com/apigear-io/cli.git",
			wantErr:  false,
			wantHost: "github.com",
			wantRepo: "cli",
		},
		{
			name:     "GitHub SSH URL",
			url:      "git@github.com:apigear-io/cli.git",
			wantErr:  false,
			wantHost: "github.com",
			wantRepo: "cli",
		},
		{
			name:     "GitLab HTTPS URL",
			url:      "https://gitlab.com/user/project.git",
			wantErr:  false,
			wantHost: "gitlab.com",
			wantRepo: "project",
		},
		{
			name:     "Bitbucket HTTPS URL",
			url:      "https://bitbucket.org/user/project.git",
			wantErr:  false,
			wantHost: "bitbucket.org",
			wantRepo: "project",
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: true,
		},
		{
			name:    "invalid URL",
			url:     "not a valid url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAsVcsUrl(tt.url)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, tt.wantHost, string(result.Host))
				assert.Equal(t, tt.wantRepo, result.Name)
			}
		})
	}
}
