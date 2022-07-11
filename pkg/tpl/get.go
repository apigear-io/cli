package tpl

import (
	"apigear/pkg/log"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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
	err = ExecGit([]string{"clone", repo, target}, "")
	if err != nil {
		log.Warnf("failed to clone template from %s into %s", repo, target)
	}
	return err
}

// ExecGit executes a git command.
func ExecGit(args []string, cwd string) error {
	if cwd == "" {
		cwd = GetPackageDir()
	}
	log.Debugf("exec git %s", args)
	_, err := exec.LookPath("git")
	if err != nil {
		log.Warnf("git not found")
		return err
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = cwd
	err = cmd.Run()
	if err != nil {
		log.Warnf("git %s failed", args)
		return err
	}
	return nil
}

func IsUrl(path string) bool {
	u, err := url.Parse(path)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	return s.IsDir() && os.IsExist(err)
}
