package project

type DocumentInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type ProjectInfo struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Documents []DocumentInfo `json:"documents"`
}
