package prj

import (
	"apigear/pkg/config"
	"apigear/pkg/log"
	"os"
	"path"
)

func readProject(d string) (ProjectInfo, error) {
	log.Debugf("Read Project %s", d)
	// check if source is directory
	if _, err := os.Stat(d); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(path.Join(d, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// read apigear directory
	entries, err := os.ReadDir(path.Join(d, "apigear"))
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
			Path: path.Join(d, "apigear", entry.Name()),
			Type: "module",
		})
	}
	project := ProjectInfo{
		Name:      path.Base(d),
		Path:      d,
		Documents: docs,
	}
	// save current project
	currentProject = project
	config.AppendRecentProject(d)
	return project, nil
}