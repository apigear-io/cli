package git

import (
	"sort"

	"github.com/go-git/go-git/v5"
	gconf "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

// func FetchRepo(source string) error {
// 	// get git info for template
// 	repo, err := git.PlainOpen(source)
// 	if err != nil {
// 		return err
// 	}
// 	// Fetch using default options
// 	err = repo.Fetch(&git.FetchOptions{})
// 	if err != nil && err != git.NoErrAlreadyUpToDate {
// 		return err
// 	}
// 	return nil
// }

type RepoInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Author      string            `json:"author"`
	Path        string            `json:"path"`
	Git         string            `json:"git"`
	Commit      string            `json:"commit"`
	Version     VersionInfo       `json:"version"`
	Latest      VersionInfo       `json:"latest"`
	Versions    VersionCollection `json:"versions"`
	InCache     bool              `json:"inCache"`
	InRegistry  bool              `json:"inRegistry"`
}

func (r *RepoInfo) FQN() string {
	if r.Version.Name != "" {
		return r.Name + "@" + r.Version.Name
	}
	return r.Name
}

func (r *RepoInfo) VersionName() string {
	if r.Version.Name != "" {
		return r.Version.Name
	}
	return r.Commit
}

func SortRepoInfo(infos []*RepoInfo) {
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name
	})
}

func LocalRepoInfo(source string) (*RepoInfo, error) {
	// get git info for template
	info := &RepoInfo{}
	repo, err := git.PlainOpen(source)
	if err != nil {
		return nil, err
	}
	cfg, err := repo.Config()
	if err != nil {
		return nil, err
	}
	info.Git = cfg.Remotes["origin"].URLs[0]
	info.Path = source
	info.Author = cfg.Author.Name
	latest, versions, err := GetTagsFromRepo(source)
	if err != nil {
		return nil, err
	}
	info.Latest = latest
	info.Versions = versions

	// extract latest head hash
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}
	info.Commit = head.Hash().String()
	return info, nil
}

func RemoteRepoInfo(url string) (*RepoInfo, error) {
	log.Debug().Msgf("remote repo info for %s", url)
	result := &RepoInfo{
		Git: url,
	}
	remote := git.NewRemote(memory.NewStorage(), &gconf.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})
	result.Git = remote.Config().URLs[0]
	result.Name = remote.Config().Name
	latestTag, tags, err := GetTagsFromRemote(remote)
	if err != nil {
		return nil, err
	}
	result.Latest = latestTag
	result.Versions = tags
	return result, nil
}
