package git

import (
	"net/url"

	"github.com/gitsight/go-vcsurl"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func ParseAsUrl(urlStr string) (*url.URL, error) {
	// Use go-git's transport.NewEndpoint which is secure and well-maintained
	endpoint, err := transport.NewEndpoint(urlStr)
	if err != nil {
		return nil, err
	}
	// Convert endpoint to standard URL
	return url.Parse(endpoint.String())
}

func IsValidGitUrl(urlStr string) bool {
	// Use go-git's transport.NewEndpoint for validation
	_, err := transport.NewEndpoint(urlStr)
	return err == nil
}

func ParseAsVcsUrl(url string) (*vcsurl.VCS, error) {
	return vcsurl.Parse(url)
}
