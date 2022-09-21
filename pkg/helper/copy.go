package helper

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFiles(source, target string) error {
	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		s := Join(source, entry.Name())
		t := Join(target, entry.Name())
		if entry.IsDir() {
			if err := CopyFiles(s, t); err != nil {
				return err
			}
		} else {
			if err := CopyFile(s, t); err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(source, target string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()
	// ensure target directory exists
	err = os.MkdirAll(filepath.Dir(target), 0755)
	if err != nil {
		return err
	}
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
