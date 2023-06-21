package repos

import (
	"fmt"
	"strings"
)

func EnsureRepoID(name string) string {
	parts := strings.Split(name, "@")
	if len(parts) != 2 {
		return fmt.Sprintf("%s@latest", name)
	}
	if parts[1] == "" {
		parts[1] = "latest"
	}
	return fmt.Sprintf("%s@%s", parts[0], parts[1])
}

func IsRepoID(name string) bool {
	return strings.Contains(name, "@")
}

// SplitRepoID splits a qualified name into name and version
func SplitRepoID(name string) (string, string) {
	name = EnsureRepoID(name)
	parts := strings.Split(name, "@")
	if len(parts) != 2 {
		log.Fatal().Msgf("invalid repo id: %s", name)
	}
	return parts[0], parts[1]
}

func MakeRepoID(name, version string) string {
	if version == "" {
		version = "latest"
	}
	name = NameFromRepoID(name)
	return fmt.Sprintf("%s@%s", name, version)
}

func NameFromRepoID(name string) string {
	if !IsRepoID(name) {
		return name
	}
	name, _ = SplitRepoID(name)
	return name
}

func VersionFromRepoID(name string) string {
	name = EnsureRepoID(name)
	_, version := SplitRepoID(name)
	return version
}
