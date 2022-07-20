package git

import (
	"os/exec"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/log"
)

// ExecGit executes a git command.
func ExecGit(args []string, cwd string) (string, error) {
	if cwd == "" {
		cwd = config.GetPackageDir()
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
