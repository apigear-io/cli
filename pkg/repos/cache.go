package repos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

var Cache *cache = NewDefaultCache()

// cache is a template cache
// It contains templates in the local cache and registry
// It also contains methods to manage the cache
// Templates are managed in a local git repository,
// cloned from a remote git repository
type cache struct {
	cacheDir string
}

// NewTemplateCache creates a new template cache
func New(cacheDir string) *cache {
	return &cache{
		cacheDir: cacheDir,
	}
}

// NewDefault creates a new template cache with default configuration
func NewDefaultCache() *cache {
	cacheDir := cfg.CacheDir()
	return New(cacheDir)
}

// List lists all templates in the cache
func (c *cache) List() ([]*git.RepoInfo, error) {
	cached, err := c.ListCachedRepos()
	if err != nil {
		return nil, err
	}
	git.SortRepoInfo(cached)
	return cached, nil
}

func (c *cache) ListVersions(repoID string) ([]*git.RepoInfo, error) {
	repoID = EnsureRepoID(repoID)
	infos, err := c.List()
	if err != nil {
		return nil, err
	}
	var versions []*git.RepoInfo
	for _, info := range infos {
		if info.Name == repoID {
			versions = append(versions, info)
		}
	}
	return versions, nil
}

func (c *cache) Search(pattern string) ([]*git.RepoInfo, error) {
	result, err := c.List()
	if err != nil {
		return nil, err
	}
	var filtered []*git.RepoInfo
	for _, info := range result {
		if helper.Contains(info.Name, pattern) {
			filtered = append(filtered, info)
		}
	}
	return filtered, nil
}

// Remove removes template by name from the cache
func (c *cache) Remove(name string) error {
	name = EnsureRepoID(name)
	log.Info().Msgf("remove template %s from %s", name, c.cacheDir)
	// remove dir from packageDir
	// check if dir exists
	target := helper.Join(c.cacheDir, name)
	if !helper.IsDir(target) {
		return fmt.Errorf("template %s does not exist", name)
	}
	return os.RemoveAll(target)
}

func (c *cache) Clean() error {
	log.Info().Msgf("remove all templates from %s", c.cacheDir)
	// remove dir from packageDir
	// check if dir exists
	err := os.RemoveAll(c.cacheDir)
	if err != nil {
		return err
	}
	return os.MkdirAll(c.cacheDir, os.ModePerm)
}

// Info returns information about a template
// either from an installed of from a template registry
func (c *cache) Info(repoID string) (*git.RepoInfo, error) {
	repoID = EnsureRepoID(repoID)
	// get git info for template
	target := helper.Join(c.cacheDir, repoID)
	if !helper.IsDir(target) {
		return nil, fmt.Errorf("template %s not found", repoID)
	}
	info, err := git.LocalRepoInfo(target)
	if err != nil {
		return nil, err
	}
	info.Name = repoID
	return info, nil
}

// Exists returns true if template exists in the cache
func (c *cache) Exists(repoID string) bool {
	repoID = EnsureRepoID(repoID)
	target := helper.Join(c.cacheDir, repoID)
	return helper.IsDir(target)
}

// Install installs template template registry into the cache
func (c *cache) Install(url string, version string) (string, error) {
	if version == "" {
		return "", fmt.Errorf("version is required")
	}
	vcs, err := git.ParseAsVcsUrl(url)
	if err != nil {
		return "", err
	}
	name := vcs.FullName
	name = MakeRepoID(name, version)
	dst := helper.Join(c.cacheDir, name)
	err = git.CloneOrPull(url, dst)
	if err != nil {
		return "", err
	}
	err = git.CheckoutTag(dst, version)
	if err != nil {
		return "", err
	}
	return name, nil
}

// ListTemplates lists all templates in the cache
func (c *cache) ListCachedRepos() ([]*git.RepoInfo, error) {
	// walk package dir to find a dir that contains a .git dir
	var infos []*git.RepoInfo
	err := filepath.Walk(c.cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk template dir: %s", err)
		}
		if info.IsDir() && info.Name() != "." && info.Name() != ".." {
			if helper.IsDir(helper.Join(path, ".git")) {
				name, err := filepath.Rel(c.cacheDir, path)
				if err != nil {
					return fmt.Errorf("get relative path for %s", path)
				}
				info, err := git.LocalRepoInfo(path)
				if err != nil {
					return fmt.Errorf("get git info for %s", path)
				}
				info.Name = strings.ReplaceAll(name, "\\", "/")
				info.InCache = true
				info.InRegistry = false
				infos = append(infos, info)
				// no need to traverse into this dir
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list templates: %s", err)
	}
	git.SortRepoInfo(infos)
	return infos, nil
}

// Upgrade upgrade templates from remote git repo
func (c *cache) Upgrade(names []string) error {
	log.Info().Msgf("update templates %s", names)
	for _, name := range names {
		// update template
		name = EnsureRepoID(name)
		dst := helper.Join(cfg.CacheDir(), name)
		err := git.Pull(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpgradeAll upgrade all templates from remote git repo
func (c *cache) UpgradeAll() error {
	// update all templates
	templates, err := c.List()
	if err != nil {
		return err
	}
	names := []string{}
	for _, t := range templates {
		names = append(names, t.Name)
	}
	return c.Upgrade(names)
}
