package helper

import (
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		s := Join(src, entry.Name())
		t := Join(dst, entry.Name())
		if entry.IsDir() {
			if err := CopyDir(s, t); err != nil {
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
