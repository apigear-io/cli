package git

import (
	"path/filepath"

	"github.com/apigear-io/helper"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func IsLocalGitRepo(source string) bool {
	_, err := git.PlainOpen(source)
	return err == nil
}

// LocalRepoInfo returns information about a local git-repository
func LocalRepoInfo(base, source string) (*RepoInfo, error) {
	log.Debug().Msgf("local repo info for %s", source)
	repo, err := git.PlainOpen(source)
	if err != nil {
		return nil, err
	}
	path, err := templateRelativePath(base, source)
	if err != nil {
		return nil, err
	}
	cfg, err := repo.Config()
	if err != nil {
		return nil, err
	}
	headRef, err := repo.Head()
	if err != nil {
		return nil, err
	}
	commit := headRef.Hash().String()
	tag, err := getTagFromHeadCommit(repo)
	if err != nil {
		return nil, err
	}
	author := cfg.Author.Name
	info := &RepoInfo{
		Name:        path,
		Path:        source,
		Author:      author,
		Description: "local",
		InCache:     true,
		InRegistry:  false,
		Commit:      commit,
		Tag:         tag,
	}
	return info, nil
}

// templateRelativePath returns the relative path of source from base
func templateRelativePath(base, source string) (string, error) {
	rel, err := helper.RelativePath(base, source)
	if err != nil {
		return "", err
	}
	rel = filepath.ToSlash(rel)
	return rel, nil
}

// resolveTagRefToCommitHash resolves a tag reference to a commit hash
func resolveTagRefToCommitHash(repo *git.Repository, tagRef *plumbing.Reference) (*plumbing.Hash, error) {
	rev := plumbing.Revision(tagRef.Name().String())
	tagCommitHash, err := repo.ResolveRevision(rev)
	if err != nil {
		log.Warn().Msgf("error resolving revision %s: %s", rev, err)
		return nil, err
	}
	return tagCommitHash, nil
}

// getTagFromHeadCommit returns the tag of the current commit
func getTagFromHeadCommit(repo *git.Repository) (string, error) {
	headRef, err := repo.Head()
	if err != nil {
		return "", err
	}
	tagRefs, err := repo.Tags()
	if err != nil {
		return "", err
	}
	var tag string
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		// check if tag points to current commit
		tagCommitHash, err := resolveTagRefToCommitHash(repo, tagRef)
		if err != nil {
			return nil
		}
		// if tag points to current commit, use it
		if *tagCommitHash == headRef.Hash() {
			tag = tagRef.Name().Short()
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return tag, nil
}
