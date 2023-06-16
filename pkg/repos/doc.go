// Package tpl contains cache and registry for template repositories.
// Cache is a local directory that contains cloned template repositories.
// Registry is a git repository that contains a list of templates and their repository URL.
//
// Cache can install, upgrade, remove, and list local cached repositories.
// A repository is a git repository that contains a template.
// A repository can exists several times in different versions in the cache.
// The repository directory is the name of the repository and the version (e.g. $organization/$name/$version)

// Registry can search and list repositories.
//

package repos
