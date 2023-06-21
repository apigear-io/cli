package registry

import (
	"fmt"
	"strconv"

	"github.com/apigear-io/cli/pkg/git"
	"github.com/pterm/pterm"
)

func displayRepoInfos(infos []*git.RepoInfo) {
	cells := make([][]string, len(infos)+1)
	cells[0] = []string{"name", "installed", "registry", "url"}
	for i, info := range infos {
		cells[i+1] = []string{
			info.Name,
			strconv.FormatBool(info.InCache),
			strconv.FormatBool(info.InRegistry),
			info.Git,
		}
	}
	err := pterm.DefaultTable.WithHasHeader().WithData(cells).Render()
	if err != nil {
		pterm.Error.Println(err)
	}
}

func DisplayTemplateInfos(infos []*git.RepoInfo) {
	cells := make([][]string, len(infos)+1)
	cells[0] = []string{"source", "url", "installed", "latest"}
	for i, info := range infos {
		vers := info.Commit
		if info.Version.Name != "" {
			vers = info.Version.Name
		}
		cells[i+1] = []string{
			info.Name,
			info.Git,
			vers,
			info.Latest.Name,
		}
	}

	err := pterm.DefaultTable.WithHasHeader().WithData(cells).Render()
	if err != nil {
		pterm.Error.Println(err)
	}
}

func DisplayTemplateInfo(info *git.RepoInfo) {
	vers := info.Commit
	if info.Version.Name != "" {
		vers = info.Version.Name
	}
	fmt.Printf("Name:    	%s\n", info.Name)
	fmt.Printf("URL:      	%s\n", info.Git)
	fmt.Printf("Version:  	%s\n", vers)
	fmt.Printf("Latest:   	%s\n", info.Latest.Name)
	fmt.Printf("Versions: 	%v\n", info.Versions)
	fmt.Println()
}
