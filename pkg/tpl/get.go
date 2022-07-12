package tpl

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
)

// GetTemplate clones a template using git from an url into a local directory.
func GetTemplate(name string, repo string) error {
	// check if repo is a local dir or an url
	target := filepath.Join(GetPackageDir(), name)
	_, err := os.Stat(target)
	if err == nil {
		return fmt.Errorf("%s already exists", name)
	}
	log.Infof("clone template from %s into %s", repo, target)
	_, err = ExecGit([]string{"clone", repo, target}, "")
	if err != nil {
		log.Warnf("failed to clone template from %s into %s", repo, target)
	}
	return err
}

// GetNameFromURL returns the name of the template from an git url.
func GetNameFromURL(source string) (string, error) {
	u, err := url.Parse(source)
	if err != nil {
		return "", fmt.Errorf("invalid url: %s", source)
	}
	return u.Path, nil
}

// ExecGit executes a git command.
func ExecGit(args []string, cwd string) (string, error) {
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
