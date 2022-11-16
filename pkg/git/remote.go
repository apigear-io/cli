package git

import (
	"sort"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	gconf "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

// VersionCollection is a collection of tags
// it implements sort.Interface
type VersionCollection []VersionInfo

// Len is the number of elements in the collection.
func (c VersionCollection) Len() int {
	return len(c)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c VersionCollection) Less(i, j int) bool {
	return c[i].Version.LessThan(c[j].Version)
}

// Swap swaps the elements with indexes i and j.
func (c VersionCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Latest returns the latest tag info
func (c VersionCollection) Latest() VersionInfo {
	return c[0]
}

func (c VersionCollection) AsList() []string {
	result := make([]string, 0)
	for _, v := range c {
		result = append(result, v.Name)
	}
	return result
}

// VersionInfo contains information about a tag
type VersionInfo struct {
	Name    string          `json:"name"`
	SHA     string          `json:"sha"`
	Version *semver.Version `json:"version"`
}

type RepoInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Author      string            `json:"author"`
	Path        string            `json:"path"`
	Git         string            `json:"git"`
	Commit      string            `json:"commit"`
	Latest      string            `json:"latest"`
	Versions    VersionCollection `json:"tags"`
	InCache     bool              `json:"inCache"`
	InRegistry  bool              `json:"inRegistry"`
}

func SortRepoInfo(infos []*RepoInfo) {
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name
	})
}

func RemoteRepoInfo(url string) (RepoInfo, error) {
	log.Debug().Msgf("remote repo info for %s", url)
	result := RepoInfo{
		Git: url,
	}
	remote := git.NewRemote(memory.NewStorage(), &gconf.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})
	refs, err := remote.List(&git.ListOptions{Auth: auth()})
	if err != nil {
		return RepoInfo{}, err
	}
	var latestTag VersionInfo
	tags := make(VersionCollection, 0)
	for _, ref := range refs {
		if ref.Name().IsTag() {
			v, err := semver.NewVersion(ref.Name().Short())
			if err != nil {
				continue
			}
			tag := VersionInfo{
				Name:    ref.Name().Short(),
				SHA:     ref.Hash().String(),
				Version: v,
			}
			if latestTag.Version == nil {
				// first tag
				latestTag = tag
			} else if tag.Version.GreaterThan(latestTag.Version) {
				// newer tag
				latestTag = tag
			}
			tags = append(tags, tag)
			result.Versions = tags
		}
	}
	result.Latest = latestTag.SHA
	return result, nil
}
