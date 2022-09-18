package prj

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/vfs"
)

var currentProject *ProjectInfo

func OpenProject(source string) (*ProjectInfo, error) {
	log.Infof("Open Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return nil, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(filepath.Join(source, "apigear")); err != nil {
		return nil, err
	}

	return readProject(source)
}

func CurrentProject() *ProjectInfo {
	return currentProject
}

// InitProject initializes a new project inside destination
func InitProject(d string) (*ProjectInfo, error) {
	log.Debugf("Init Project %s", d)
	// create destination if not exists
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return nil, err
		}
	}
	// create apigear directory
	if err := os.Mkdir(filepath.Join(d, "apigear"), 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
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

func GetProjectInfo(d string) (*ProjectInfo, error) {
	return readProject(d)
}

func RecentProjectInfos() []*ProjectInfo {
	var infos []*ProjectInfo
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
func ImportProject(repo string, dir string) (*ProjectInfo, error) {
	log.Infof("Import Project %s", repo)
	// check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("target directory %s does not exist", dir)
	}
	// check if repo is valid url
	if !git.IsValidGitUrl(repo) {
		return nil, fmt.Errorf("invalid repo url: '%s'", repo)
	}
	err := git.Clone(repo, dir)
	if err != nil {
		return nil, fmt.Errorf("failed to clone project repository: %s", err)
	}
	return readProject(dir)
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
func CreateProjectDocument(prjDir string, docType string, name string) (string, error) {
	target := filepath.Join(prjDir, "apigear", MakeDocumentName(docType, name))
	var err error
	switch docType {
	case "module":
		err = writeDemo(target, vfs.DemoModule)
	case "solution":
		err = writeDemo(target, vfs.DemoSolution)
	case "scenario":
		err = writeDemo(target, vfs.DemoScenario)
	default:
		err = fmt.Errorf("invalid document type %s", docType)
	}
	if err != nil {
		return "", err
	}
	return target, nil
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
