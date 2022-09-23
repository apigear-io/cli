package git

import (
	"os/exec"

	"github.com/apigear-io/cli/pkg/config"
)

// ExecGit executes a git command.
func ExecGit(args []string, cwd string) (string, error) {
	if cwd == "" {
		cwd = config.GetPackageDir()
	}
	log.Debug().Msgf("exec git %s", args)
	_, err := exec.LookPath("git")
	if err != nil {
		log.Warn().Msgf("git not found")
		return "", err
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = cwd
	out, err := cmd.Output()
	if err != nil {
		log.Warn().Msgf("git %s failed", args)
		return "", err
	}
	log.Debug().Msgf("git output: %s", out)
	return string(out), nil

}
