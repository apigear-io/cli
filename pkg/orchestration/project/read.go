package project

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/foundation/config"
	"github.com/apigear-io/cli/pkg/foundation"
)

func ReadProject(d string) (*ProjectInfo, error) {
	log.Debug().Msgf("Read Project %s", d)
	// check if source is directory
	if _, err := os.Stat(d); err != nil {
		return nil, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(foundation.Join(d, "apigear")); err != nil {
		return nil, err
	}
	// read apigear directory
	entries, err := os.ReadDir(foundation.Join(d, "apigear"))
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
			Path: foundation.Join(d, "apigear", entry.Name()),
			Type: foundation.GetDocumentType(entry.Name()),
		})
	}
	project := &ProjectInfo{
		Name:      filepath.Base(d),
		Path:      d,
		Documents: docs,
	}
	// save current project
	currentProject = project
	err = config.AppendRecentEntry(d)
	if err != nil {
		return nil, err
	}
	return project, nil
}
