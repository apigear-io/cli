package repos

import (
	"fmt"
	"strings"
)

func InstallTemplate(name, version string) error {
	fqn := MakeFQN(name, version)
	return InstallTemplateFromFQN(fqn)
}

// InstallTemplateFromFQN tyies to install a template
// from a fully qualified name (e.g. name@version)
func InstallTemplateFromFQN(fqn string) error {
	if !IsFQN(fqn) {
		return fmt.Errorf("invalid fqn: %s", fqn)
	}
	if Cache.Exists(fqn) {
		return nil
	}
	info, err := Registry.Get(fqn)
	if err != nil {
		return err
	}
	if info.Latest.Name == "" {
		return fmt.Errorf("no latest version found for template: %s", fqn)
	}
	url := info.Git
	version := info.Latest.Name
	log.Info().Msgf("installing template %s@%s from %s", fqn, version, url)
	_, err = Cache.Install(url, version)
	if err != nil {
		return err
	}
	return nil
}

func IsFQN(fqn string) bool {
	return strings.Contains(fqn, "@")
}

func MakeFQN(name, version string) string {
	return fmt.Sprintf("%s@%s", name, version)
}

func ParseFQN(fqn string) (string, string, error) {
	parts := strings.Split(fqn, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid fqn: %s", fqn)
	}
	return parts[0], parts[1], nil
}

func NameFromFQN(fqn string) string {
	if !IsFQN(fqn) {
		return fqn
	}
	name, _, _ := ParseFQN(fqn)
	return name
}
