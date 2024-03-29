package prj

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
)

func ReadProject(d string) (*ProjectInfo, error) {
	log.Debug().Msgf("Read Project %s", d)
	// check if source is directory
	if _, err := os.Stat(d); err != nil {
		return nil, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(helper.Join(d, "apigear")); err != nil {
		return nil, err
	}
	// read apigear directory
	entries, err := os.ReadDir(helper.Join(d, "apigear"))
	if err != nil {
		return nil, err
	}
	// convert entries to documents
	var docs []DocumentInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		docs = append(docs, DocumentInfo{
			Name: entry.Name(),
			Path: helper.Join(d, "apigear", entry.Name()),
			Type: helper.GetDocumentType(entry.Name()),
		})
	}
	project := &ProjectInfo{
		Name:      filepath.Base(d),
		Path:      d,
		Documents: docs,
	}
	// save current project
	currentProject = project
	err = cfg.AppendRecentEntry(d)
	if err != nil {
		return nil, err
	}
	return project, nil
}
