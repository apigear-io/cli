package git

import (
	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
)

// GetLatestTagFromRepo returns the latest tag and a list of all tags from a repo
func GetTagsFromRemote(remote *git.Remote) (VersionInfo, VersionCollection, error) {
	refs, err := remote.List(&git.ListOptions{Auth: auth()})
	if err != nil {
		return VersionInfo{}, nil, err
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
		}
	}
	return latestTag, tags, nil
}

func GetTagsFromRepo(source string) (VersionInfo, VersionCollection, error) {
	repo, err := git.PlainOpen(source)
	if err != nil {
		return VersionInfo{}, nil, err
	}
	var latestTag VersionInfo
	tags := make(VersionCollection, 0)
	// extract latest head hash
	head, err := repo.Head()
	if err != nil {
		return VersionInfo{}, nil, err
	}
	// extract latest tag
	tagRefs, err := repo.Tags()
	if err != nil {
		return VersionInfo{}, nil, err
	}
	for {
		tag, err := tagRefs.Next()
		if err != nil {
			break
		}
		name := tag.Name().Short()
		sha := tag.Hash().String()
		vers, err := semver.NewVersion(name)
		if err != nil {
			continue
		}
		tagInfo := VersionInfo{
			Name:    name,
			SHA:     sha,
			Version: vers,
		}
		tags = append(tags, tagInfo)
		// compare commit hash with tag hash
		if tag.Hash() == head.Hash() {
			latestTag = VersionInfo{
				Name: tag.Name().Short(),
				SHA:  tag.Hash().String(),
			}
		}
	}
	return latestTag, tags, nil
}
