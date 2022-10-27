package git

import (
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/go-git/go-git/v5"
)

func Clone(src string, dst string) error {
	log.Debug().Msgf("clone %s %s", src, dst)
	_, err := git.PlainClone(dst, false, &git.CloneOptions{
		URL:      src,
		Auth:     auth(),
		Progress: os.Stdout,
	})
	return err
}

func CloneOrPull(src string, dst string) error {
	log.Debug().Msgf("clone or pull %s %s", src, dst)
	if helper.IsDir(dst) {
		return Pull(dst)
	}
	return Clone(src, dst)
}
