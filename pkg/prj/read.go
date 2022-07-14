package prj

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/log"
)

func readProject(d string) (ProjectInfo, error) {
	log.Debugf("Read Project %s", d)
	// check if source is directory
	if _, err := os.Stat(d); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(filepath.Join(d, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// read apigear directory
	entries, err := os.ReadDir(filepath.Join(d, "apigear"))
	if err != nil {
		return ProjectInfo{}, err
	}
	// convert entries to documents
	var docs []DocumentInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		docs = append(docs, DocumentInfo{
			Name: entry.Name(),
			Path: filepath.Join(d, "apigear", entry.Name()),
			Type: "module",
		})
	}
	project := ProjectInfo{
		Name:      filepath.Base(d),
		Path:      d,
		Documents: docs,
	}
	// save current project
	currentProject = project
	config.AppendRecentEntry(d)
	return project, nil
}
