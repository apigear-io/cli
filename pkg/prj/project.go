package prj

import (
	"apigear/pkg/config"
	"apigear/pkg/log"
	"os"
	"os/exec"
	"path"
)

var currentProject ProjectInfo

func OpenProject(source string) (ProjectInfo, error) {
	log.Infof("Open Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(path.Join(source, "apigear")); err != nil {
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
	if err := os.Mkdir(path.Join(d, "apigear"), 0755); err != nil {
		if !os.IsExist(err) {
			return ProjectInfo{}, err
		}
	}
	// write demo module and demo solution
	if err := writeDemoModule(path.Join(d, "apigear", "demo.module.yaml")); err != nil {
		log.Debugf("Failed to write demo module: %s", err)
	}
	if err := writeDemoSolution(path.Join(d, "apigear", "demo.solution.yaml")); err != nil {
		log.Debugf("Failed to write demo solution: %s", err)
	}
	return readProject(d)
}

func GetProjectInfo(d string) (ProjectInfo, error) {
	return readProject(d)
}

func RecentProjectInfos() []ProjectInfo {
	var infos []ProjectInfo
	for _, d := range config.ReadRecentProjects() {
		info, err := readProject(d)
		if err != nil {
			log.Warnf("Failed to read project %s: %s", d, err)
			config.RemoveRecentFile(d)
			continue
		}
		infos = append(infos, info)
	}
	return infos
}

func OpenEditor(d string) error {
	path, err := exec.LookPath("code")
	if err != nil {
		return err
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

func ImportProject(source string, target string) (ProjectInfo, error) {
	log.Infof("Import Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); err != nil {
		return ProjectInfo{}, err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(path.Join(source, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// check if destination is directory
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return ProjectInfo{}, err
	}
	// check if destination contains apigear directory
	if _, err := os.Stat(path.Join(target, "apigear")); err != nil {
		return ProjectInfo{}, err
	}
	// copy apigear directory
	if err := copyFiles(source, target); err != nil {
		return ProjectInfo{}, err
	}
	return readProject(target)
}

func PackProject(source string, target string) (string, error) {
	log.Infof("Pack Project %s", source)
	// check if source is directory
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return "", err
	}
	// check if source contains apigear directory
	if _, err := os.Stat(path.Join(source, "apigear")); err != nil {
		return "", err
	}
	// create archive file
	if err := createArchive(source, target); err != nil {
		return "", err
	}
	return target, nil
}
