package gen

import (
	"bytes"
	"text/template"
)

type Renderer struct {
	t *template.Template
}

// NewRenderer creates a new renderer
func NewRenderer(tplDir string) *Renderer {
	tpl, err := template.ParseGlob(tplDir + "/*.tmpl")
	if err != nil {
		panic(err)
	}
	return &Renderer{t: tpl}
}

// RenderString renders a string from a string template
func (r *Renderer) RenderString(s string, data any) (string, error) {
	buf := bytes.Buffer{}
	t := template.Must(template.New("").Parse(s))
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderFile renders a string from a file template
// the file must be in the renderer's template directory
func (r *Renderer) RenderFile(fn string, data any) (string, error) {
	buf := bytes.Buffer{}
	err := r.t.ExecuteTemplate(&buf, fn, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
