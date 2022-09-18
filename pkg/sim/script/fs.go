package script

import "os"

func ReadScript(fn string) (string, error) {
	bytes, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
