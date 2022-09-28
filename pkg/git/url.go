package git

import (
	"net/url"

	"github.com/gitsight/go-vcsurl"
	urls "github.com/whilp/git-urls"
)

func ParseGitUrl(url string) (*url.URL, error) {
	return urls.Parse(url)
}

func IsValidGitUrl(url string) bool {
	_, err := urls.ParseTransport(url)
	return err == nil
}

func GitUrlInfo(url string) (*vcsurl.VCS, error) {
	return vcsurl.Parse(url)
}
