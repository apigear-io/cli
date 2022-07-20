package git

import (
	"net/url"
	"strings"

	urls "github.com/whilp/git-urls"
)

func ParseGitUrl(url string) (*url.URL, error) {
	return urls.Parse(url)
}

func IsValidGitUrl(url string) bool {
	_, err := urls.ParseTransport(url)
	return err == nil
}

func RepositoryNameFromGitUrl(url string) (string, error) {
	u, err := ParseGitUrl(url)
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(u.Path, ".git") {
		return strings.TrimSuffix(u.Path, ".git"), nil
	}
	return u.Path, nil
}
