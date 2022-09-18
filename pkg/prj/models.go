package prj

type DocumentInfo struct {
	Name string
	Path string
	Type string
}

type ProjectInfo struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Documents []DocumentInfo `json:"documents"`
}
