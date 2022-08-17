package gen

import (
	"crypto/md5"
	"os"
)

func CompareContentWithFile(sourceBytes []byte, target string) (bool, error) {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return false, nil
	}
	var targetBytes, err = os.ReadFile(target)
	if err != nil {
		return false, err
	}
	return md5.Sum(sourceBytes) == md5.Sum(targetBytes), nil
}
