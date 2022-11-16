package tpl

import (
	"encoding/json"
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

type TemplateRegistry struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Entries     []*git.RepoInfo `json:"entries"`
}

// ReadRegistry reads the registry file from path
func ReadRegistry() (*TemplateRegistry, error) {
	src := cfg.RegistryCachePath()
	if !helper.IsFile(src) {
		log.Info().Msgf("registry file not found: %s", src)
		err := UpdateRegistry()
		if err != nil {
			return nil, err
		}
	}
	// read registry.json
	bytes, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}
	// unmarshal
	var registry TemplateRegistry
	err = json.Unmarshal(bytes, &registry)
	if err != nil {
		return nil, err
	}
	// sort entries
	git.SortRepoInfo(registry.Entries)
	return &registry, nil
}

// WriteRegistry writes the registry to path
func WriteRegistry(r *TemplateRegistry) error {
	dst := cfg.RegistryCachePath()
	bytes, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dst, bytes, 0644)
}
