package tpl

import "os"

type TemplateInfo struct {
	Name string
}

func ListTemplates() ([]TemplateInfo, error) {
	// list all dirs in packageDir
	dir := GetPackageDir()
	items, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var infos []TemplateInfo
	for _, d := range items {
		if d.IsDir() {
			infos = append(infos, TemplateInfo{Name: d.Name()})
		}
	}
	return infos, nil
}
