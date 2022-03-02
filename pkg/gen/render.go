package gen

import (
	"fmt"

	"github.com/flosch/pongo2"
)

func init() {
	pongo2.SetAutoescape(false)
}

type Renderer struct {
	s *pongo2.TemplateSet
}

// NewRenderer creates a new renderer
func NewRenderer(tplDir string) *Renderer {
	l, err := pongo2.NewLocalFileSystemLoader(tplDir)
	if err != nil {
		panic(fmt.Errorf("error creating fs loader for %s: %s", tplDir, err))
	}
	return &Renderer{
		s: pongo2.NewSet("", l),
	}
}

// AddTemplate adds a local file system loader to the renderer
func (r *Renderer) AddTemplateDir(dir string) error {
	l, err := pongo2.NewLocalFileSystemLoader(dir)
	if err != nil {
		return err
	}
	r.s.AddLoader(l)
	return nil
}

// RenderString renders a string from a string template
func (r *Renderer) RenderString(s string, ctx Context) (string, error) {
	return r.s.RenderTemplateString(s, ctx)
}

// RenderFile renders a string from a file template
// the file must be in the renderer's template directory
func (r *Renderer) RenderFile(fn string, ctx Context) (string, error) {
	return r.s.RenderTemplateFile(fn, ctx)
}
