package prj

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/vfs"
)

var currentProject *ProjectInfo

func OpenProject(source string) (*ProjectInfo, error) {
	log.Info().Msgf("Open Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return nil, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(helper.Join(source, "apigear")); err != nil {
		return nil, err
	}

	return ReadProject(source)
}

func CurrentProject() *ProjectInfo {
	return currentProject
}

// InitProject initializes a new project inside destination
func InitProject(d string) (*ProjectInfo, error) {
	log.Debug().Msgf("Init Project %s", d)
	// create destination if not exists
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return nil, err
		}
	}
	// create apigear directory
	if err := os.Mkdir(helper.Join(d, "apigear"), 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}
	// write demo module
	target := helper.Join(d, "apigear", "demo.module.yaml")
	if err := writeDemo(target, vfs.DemoModuleYaml); err != nil {
		log.Debug().Msgf("write demo module: %s", err)
	}
	target = helper.Join(d, "apigear", "demo.module.idl")
	if err := writeDemo(target, vfs.DemoModuleIdl); err != nil {
		log.Debug().Msgf("write demo module: %s", err)
	}
	// write demo solution
	target = helper.Join(d, "apigear", "demo.solution.yaml")
	if err := writeDemo(target, vfs.DemoSolutionYaml); err != nil {
		log.Debug().Msgf("write demo solution: %s", err)
	}
	// write demo simulation (client/service)
	target = helper.Join(d, "apigear", "demo.sim.js")
	if err := writeDemo(target, vfs.DemoSimulationJs); err != nil {
		log.Debug().Msgf("write demo service: %s", err)
	}
	return ReadProject(d)
}

func GetProjectInfo(d string) (*ProjectInfo, error) {
	return ReadProject(d)
}

func RecentProjectInfos() []*ProjectInfo {
	var infos []*ProjectInfo
	for _, d := range cfg.RecentEntries() {
		info, err := ReadProject(d)
		if err != nil {
			log.Warn().Msgf("read project %s: %s", d, err)
			err = cfg.RemoveRecentEntry(d)
			if err != nil {
				log.Warn().Msgf("remove recent entry %s: %s", d, err)
			}
			continue
		}
		infos = append(infos, info)
	}
	return infos
}

// OpenEditor opens the project directory in a editor
func OpenEditor(d string) error {
	editor := cfg.EditorCommand()
	path, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("find editor %s: %s", editor, err)
	}
	cmd := exec.Command(path, d)
	err = cmd.Run()
	if err != nil {
		log.Error().Err(err).Msgf("run editor %s from %s", editor, path)
		return fmt.Errorf("run editor %s: %s", editor, err)
	}
	return nil
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
	log.Info().Msgf("Import Project %s", repo)
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
		return nil, fmt.Errorf("clone project repository: %s", err)
	}
	return ReadProject(dir)
}

// PackProject packs the project into a zip file
func PackProject(source string, target string) (string, error) {
	log.Info().Msgf("Pack Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return "", err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(helper.Join(source, "apigear")); err != nil {
		return "", err
	}
	// create archive file
	if err := createArchive(source, target); err != nil {
		return "", err
	}
	return target, nil
}

// AddDocument creates a new document inside the project
func AddDocument(prjDir string, docType string, name string) (string, error) {
	target := helper.Join(prjDir, "apigear", MakeDocumentName(docType, name))
	var err error
	switch docType {
	case "module":
		err = writeDemo(target, vfs.DemoModuleYaml)
	case "solution":
		err = writeDemo(target, vfs.DemoSolutionYaml)
	case "simulation":
		err = writeDemo(target, vfs.DemoSimulationJs)
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
