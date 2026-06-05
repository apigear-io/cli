package repos

import (
	"regexp"
	"strings"

	"github.com/apigear-io/cli/pkg/git"
)

// scpLike matches the scp-style git syntax (e.g. git@github.com:me/tpl.git).
// The user and host parts may not contain a slash, which keeps registry repo
// IDs (e.g. apigear-io/template-go@1.2.3) from being detected as git URLs.
var scpLike = regexp.MustCompile(`^[A-Za-z0-9._-]+@[A-Za-z0-9._-]+:`)

// IsGitURL reports whether s looks like a git URL we can clone from directly,
// as opposed to a local path or a registry repo ID.
func IsGitURL(s string) bool {
	for _, prefix := range []string{"https://", "http://", "ssh://", "git://", "file://"} {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return scpLike.MatchString(s)
}

// SplitGitURLVersion splits a git URL of the form <url>@<version> into its url
// and version parts. The version separator is only recognised at or after the
// start of the URL path, so neither the scp user@host nor http credentials
// (user@host) are mistaken for a version. The version may itself contain a
// slash (e.g. a branch name like release/1.0). If no version is present the
// returned version is empty.
func SplitGitURLVersion(s string) (string, string) {
	pathStart := 0
	if i := strings.Index(s, "://"); i >= 0 {
		// scheme://[user@]host/path -> path starts at the first '/' after "://"
		rest := i + len("://")
		if slash := strings.IndexByte(s[rest:], '/'); slash >= 0 {
			pathStart = rest + slash
		} else {
			pathStart = len(s)
		}
	} else if at := strings.IndexByte(s, '@'); at >= 0 {
		// scp-style user@host:path -> path starts after the ':' following host
		if colon := strings.IndexByte(s[at:], ':'); colon >= 0 {
			pathStart = at + colon + 1
		}
	}
	if at := strings.IndexByte(s[pathStart:], '@'); at >= 0 {
		idx := pathStart + at
		return s[:idx], s[idx+1:]
	}
	return s, ""
}

// RepoNameFromGitURL derives a host-qualified cache name (e.g.
// github.com/me/tpl) from a git URL. The name is derived generically so that
// any host works (not only the ones a VCS-aware parser knows), and so that the
// https and scp forms of the same repo map to the same name. The url must not
// carry a @version suffix; strip it with SplitGitURLVersion first.
func RepoNameFromGitURL(rawurl string) (string, error) {
	u, err := git.ParseAsUrl(rawurl)
	if err != nil {
		return "", err
	}
	name := strings.TrimSuffix(u.Path, ".git")
	name = strings.Trim(name, "/")
	if host := u.Hostname(); host != "" {
		name = host + "/" + name
	}
	return name, nil
}
