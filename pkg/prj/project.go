package prj

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/vfs"
)

var currentProject ProjectInfo

func OpenProject(source string) (ProjectInfo, error) {
	log.Infof("Open Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(filepath.Join(source, "apigear")); err != nil {
		return ProjectInfo{}, err
	}

	return readProject(source)
}

func CurrentProject() ProjectInfo {
	return currentProject
}

// InitProject initializes a new project inside destination
func InitProject(d string) (ProjectInfo, error) {
	log.Debugf("Init Project %s", d)
	// create destination if not exists
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return ProjectInfo{}, err
		}
	}
	// create apigear directory
	if err := os.Mkdir(filepath.Join(d, "apigear"), 0755); err != nil {
		if !os.IsExist(err) {
			return ProjectInfo{}, err
		}
	}
	// write demo module
	target := filepath.Join(d, "apigear", "demo.module.yaml")
	if err := writeDemo(target, vfs.DemoModule); err != nil {
		log.Debugf("Failed to write demo module: %s", err)
	}
	// write demo solution
	target = filepath.Join(d, "apigear", "demo.solution.yaml")
	if err := writeDemo(target, vfs.DemoSolution); err != nil {
		log.Debugf("Failed to write demo solution: %s", err)
	}
	// write demo scenario
	target = filepath.Join(d, "apigear", "demo.scenario.yaml")
	if err := writeDemo(target, vfs.DemoScenario); err != nil {
		log.Debugf("Failed to write demo scenario: %s", err)
	}
	return readProject(d)
}

func GetProjectInfo(d string) (ProjectInfo, error) {
	return readProject(d)
}

func RecentProjectInfos() []ProjectInfo {
	var infos []ProjectInfo
	for _, d := range config.GetRecentEntries() {
		info, err := readProject(d)
		if err != nil {
			log.Warnf("Failed to read project %s: %s", d, err)
			config.RemoveRecentEntry(d)
			continue
		}
		infos = append(infos, info)
	}
	return infos
}

// OpenEditor opens the project directory in a editor
func OpenEditor(d string) error {
	editor := config.GetEditorCommand()
	path, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("Failed to find editor %s: %s", editor, err)
	}
	cmd := exec.Command(path, d)
	return cmd.Run()
}

func OpenStudio(d string) error {
	path, err := exec.LookPath("studio")
	if err != nil {
		return err
	}
	cmd := exec.Command(path, d)
	return cmd.Run()
}

// ImportProject imports a project from a zip file
func ImportProject(source string, target string) (ProjectInfo, error) {
	log.Infof("Import Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(filepath.Join(source, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// check if destination is directory
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return ProjectInfo{}, err
	}
	// check if destination contains apigear directory
	if _, err := os.Stat(filepath.Join(target, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// copy apigear directory
	// TODO: check is source is a zip file and unpack it
	if err := copyFiles(source, target); err != nil {
		return ProjectInfo{}, err
	}
	return readProject(target)
}

// PackProject packs the project into a zip file
func PackProject(source string, target string) (string, error) {
	log.Infof("Pack Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return "", err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(filepath.Join(source, "apigear")); err != nil {
		return "", err
	}
	// create archive file
	if err := createArchive(source, target); err != nil {
		return "", err
	}
	return target, nil
}

// CreateDocument creates a new document inside the project
func CreateProjectDocument(docType string, target string) error {
	switch docType {
	case "module":
		return writeDemo(target, vfs.DemoModule)
	case "solution":
		return writeDemo(target, vfs.DemoSolution)
	case "scenario":
		return writeDemo(target, vfs.DemoScenario)
	default:
		return fmt.Errorf("invalid document type %s", docType)
	}
}

// MakeDocumentName creates a new document name
func MakeDocumentName(docType string, name string) string {
	switch docType {
	case "module":
		return fmt.Sprintf("%s.module.yaml", name)
	case "solution":
		return fmt.Sprintf("%s.solution.yaml", name)
	case "scenario":
		return fmt.Sprintf("%s.scenario.yaml", name)
	default:
		return ""
	}
}
