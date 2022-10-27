package tpl

import (
	"encoding/json"
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
)

type TemplateRegistry struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Entries     []*git.RepoInfo `json:"entries"`
}

// ReadRegistry reads the registry file from path
func ReadRegistry() (*TemplateRegistry, error) {
	src := cfg.RegistryCachePath()
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
