package repos

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

type TemplateRegistry struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Entries     []*git.RepoInfo `json:"entries"`
}

var Registry = NewDefaultRegistry()

type registry struct {
	RegistryDir string
	RegistryURL string
	Registry    *TemplateRegistry
}

func NewDefaultRegistry() *registry {
	registryDir := cfg.RegistryDir()
	registryURL := cfg.RegistryUrl()
	return NewRegistry(registryDir, registryURL)
}

func NewRegistry(registryDir, registryURL string) *registry {
	return &registry{
		RegistryDir: registryDir,
		RegistryURL: registryURL,
	}
}

// Load reads the registry file from path
func (r *registry) Load() error {
	src := helper.Join(r.RegistryDir, "registry.json")
	if !helper.IsFile(src) {
		log.Debug().Msgf("registry file not found: %s", src)
		return fmt.Errorf("registry file not found: %s", src)
	}
	// read registry.json
	bytes, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	log.Info().Msgf("read registry file: %s", src)
	// unmarshal
	var registry TemplateRegistry
	err = json.Unmarshal(bytes, &registry)
	if err != nil {
		return err
	}
	for _, entry := range registry.Entries {
		entry.Name = strings.ReplaceAll(entry.Name, "\\", "/")
	}
	// sort entries
	git.SortRepoInfo(registry.Entries)
	r.Registry = &registry
	return nil
}

// Save writes the registry to path
func (c *registry) Save() error {
	dst := helper.Join(c.RegistryDir, "registry.json")
	bytes, err := json.MarshalIndent(c.Registry, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dst, bytes, 0644)
}

// List lists all templates in the registry
func (r *registry) List() ([]*git.RepoInfo, error) {
	err := r.ensureRegistry()
	if err != nil {
		return nil, err
	}
	return r.Registry.Entries, nil
}

// Info returns the template info
func (r *registry) Info(repoID string) (*git.RepoInfo, error) {
	repoID = NameFromRepoID(repoID)
	err := r.ensureRegistry()
	if err != nil {
		return nil, err
	}
	for _, info := range r.Registry.Entries {
		if info.Name == repoID {
			return info, nil
		}
	}
	return nil, fmt.Errorf("template not found: %s", repoID)
}

// Search searches for templates in the registry
func (c *registry) Search(pattern string) ([]*git.RepoInfo, error) {
	err := c.ensureRegistry()
	if err != nil {
		return nil, err
	}
	if pattern == "" {
		return c.Registry.Entries, nil
	}
	for _, info := range c.Registry.Entries {
		if helper.Contains(info.Name, pattern) {
			return []*git.RepoInfo{info}, nil
		}
	}
	return nil, nil
}

func (c *registry) Get(repoID string) (*git.RepoInfo, error) {
	name := NameFromRepoID(repoID)
	err := c.ensureRegistry()
	if err != nil {
		return nil, err
	}
	for _, info := range c.Registry.Entries {
		if info.Name == name {
			return info, nil
		}
	}
	return nil, fmt.Errorf("template not found: %s", name)
}

func (r *registry) ensureRegistry() error {
	if !helper.IsDir(r.RegistryDir) {
		r.Reset()
	}
	if r.Registry == nil {
		err := r.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

// Update updates the local template registry
// The registry is a git repository that contains a list of templates
// and their versions.
func (r *registry) Update() error {
	log.Info().Msgf("updating registry %s", r.RegistryDir)
	err := r.Reset()
	if err != nil {
		return err
	}
	err = r.ensureRegistry()
	if err != nil {
		return err
	}
	for _, entry := range r.Registry.Entries {
		log.Info().Msgf("updating template %s", entry.Name)
		info, err := git.RemoteRepoInfo(entry.Git)
		entry.Versions = info.Versions
		entry.Latest = info.Latest
		if err != nil {
			return err
		}
	}
	return r.Save()
}

func (r *registry) Reset() error {
	log.Info().Msgf("resetting registry %s", r.RegistryDir)
	err := helper.RemoveDir(r.RegistryDir)
	if err != nil {
		return err
	}
	return git.CloneOrPull(r.RegistryURL, r.RegistryDir)
}
