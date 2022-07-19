package tpl

import (
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/apigear-io/cli/pkg/log"
	urls "github.com/whilp/git-urls"
)

func clone(repo string, target string) error {
	_, err := execGit([]string{"clone", repo, target}, "")
	return err
}

// execGit executes a git command.
func execGit(args []string, cwd string) (string, error) {
	if cwd == "" {
		cwd = GetPackageDir()
	}
	log.Debugf("exec git %s", args)
	_, err := exec.LookPath("git")
	if err != nil {
		log.Warnf("git not found")
		return "", err
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = cwd
	out, err := cmd.Output()
	if err != nil {
		log.Warnf("git %s failed", args)
		return "", err
	}
	log.Debugf("git output: %s", out)
	return string(out), nil

}

func IsUrl(path string) bool {
	u, err := url.Parse(path)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	return s.IsDir() && os.IsExist(err)
}

func ParseGitUrl(url string) (*url.URL, error) {
	return urls.Parse(url)
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
