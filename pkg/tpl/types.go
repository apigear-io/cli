package tpl

// TemplateInfo contains information about a template
type TemplateInfo struct {
	Name       string `json:"name"`
	Git        string `json:"git"`
	Commit     string `json:"commit"`
	Path       string `json:"path"`
	InCache    bool   `json:"inCache"`
	InRegistry bool   `json:"inRegistry"`
}
