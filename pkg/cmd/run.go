package cmd

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type VersionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

var root = NewRootCommand()

func Run() error {
	return root.Execute()
}
