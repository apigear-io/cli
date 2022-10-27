package git

import (
	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func auth() *http.BasicAuth {
	if cfg.GitAuthToken() == "" {
		return nil
	}
	return &http.BasicAuth{
		Username: cfg.GitAuthUser(),
		Password: cfg.GitAuthToken(),
	}
}
