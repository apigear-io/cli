package sim

import (
	"fmt"
	"path/filepath"
)

type Script struct {
	Content string `json:"content"`
	Path    string `json:"path"`
	Dir     string `json:"dir"`
	Name    string `json:"name"`
}

func NewScript(path string, content string) Script {
	return Script{
		Content: content,
		Path:    path,
		Dir:     filepath.Dir(path),
		Name:    filepath.Base(path),
	}
}

func (s Script) String() string {
	return fmt.Sprintf("Script{Name: %s, Path: %s, Dir: %s}", s.Name, s.Path, s.Dir)
}
