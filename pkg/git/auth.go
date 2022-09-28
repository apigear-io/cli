package git

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

var auth = &http.BasicAuth{
	Username: "x-oauth-basic", // yes, this can be anything except an empty string
	Password: config.GitAuthToken(),
}
