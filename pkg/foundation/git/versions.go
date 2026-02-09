package git

import "github.com/Masterminds/semver/v3"

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
	if len(c) == 0 {
		return VersionInfo{}
	}
	v := c[0]
	// iterate over all entries and return the highest
	for _, i := range c {
		if v.Version.LessThan(i.Version) {
			v = i
		}
	}
	return v
}

func (c VersionCollection) AsList() []string {
	result := make([]string, 0)
	for _, v := range c {
		result = append(result, v.Name)
	}
	return result
}

func (c VersionCollection) String() string {
	result := ""
	for _, v := range c {
		result += v.Name + ", "
	}
	return result
}

// VersionInfo contains information about a tag
type VersionInfo struct {
	Name    string          `json:"name"`
	SHA     string          `json:"sha"`
	Version *semver.Version `json:"version"`
}
