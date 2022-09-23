// registry is the registry of templates
package git

import (
	"encoding/json"
	"errors"
	"os"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/go-git/go-git/v5"
	gconf "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

var (
	registryUrl = "https://github.com/apigear-io/template-registry.git"
	registryDir = config.GetTemplateRegistryDir()
	auth        = &http.BasicAuth{
		Username: "x-oauth-basic", // yes, this can be anything except an empty string
		Password: config.GetAuthToken(),
	}
)

// TagCollection is a collection of tags
// it implements sort.Interface
type TagCollection []TagInfo

// Len is the number of elements in the collection.
func (c TagCollection) Len() int {
	return len(c)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c TagCollection) Less(i, j int) bool {
	return c[i].Version.LessThan(c[j].Version)
}

// Swap swaps the elements with indexes i and j.
func (c TagCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Latest returns the latest tag info
func (c TagCollection) Latest() TagInfo {
	return c[0]
}

// TagInfo contains information about a tag
type TagInfo struct {
	Name    string          `json:"name"`
	SHA     string          `json:"sha"`
	Version *semver.Version `json:"version"`
}

type TemplateInfo struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Author      string        `json:"author"`
	Repo        string        `json:"repo"`
	Tags        TagCollection `json:"tags"`
}

type TemplateRegistry struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Templates   []*TemplateInfo `json:"templates"`
}

// CloneUrPullRegistry clones the registry repository
// or pulls it if it already exists
func CloneOrPullRegistry() (*git.Repository, error) {
	if helper.IsDir(registryDir) {
		log.Info().Msgf("pulling registry from %s", registryUrl)
		r, err := git.PlainOpen(registryDir)
		if err != nil {
			return nil, err
		}
		w, err := r.Worktree()
		if err != nil {
			return nil, err
		}
		err = w.Pull(&git.PullOptions{Auth: auth})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return nil, err
		}
	} else {
		log.Info().Msgf("cloning registry from %s", registryUrl)
		log.Info().Msgf("cloning registry from %s", registryUrl)
		r, err := git.PlainClone(registryDir, false, &git.CloneOptions{
			URL:  registryUrl,
			Auth: auth,
		})
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, nil
}

// UpdateRegistry fills in information from the remote repository
// it retrieves all remote tags and converts them to versions
// at the end it adds the latest commit SHA for each tag
// and sorts the tags by version
func UpdateRegistry(t *TemplateInfo) error {
	// reference to the remote repository
	remote := git.NewRemote(memory.NewStorage(), &gconf.RemoteConfig{
		Name: "origin",
		URLs: []string{t.Repo},
	})
	// list all remote tags
	refs, err := remote.List(&git.ListOptions{Auth: auth})
	if err != nil {
		return err
	}
	// convert refs to tag infos with version and hash
	for _, ref := range refs {
		v, err := semver.NewVersion(ref.Name().Short())
		if err != nil {
			continue
		}
		tag := TagInfo{
			Name:    ref.Name().Short(),
			SHA:     ref.Hash().String(),
			Version: v,
		}
		if ref.Name().IsTag() {
			t.Tags = append(t.Tags, tag)
		}
	}
	sort.Sort(sort.Reverse(t.Tags))
	return nil
}

// ReadRegistry reads the registry file from path
func ReadRegistry(path string) (*TemplateRegistry, error) {
	// read registry.json
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// unmarshal
	var registry TemplateRegistry
	err = json.Unmarshal(bytes, &registry)
	if err != nil {
		return nil, err
	}
	return &registry, nil
}
