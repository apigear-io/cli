package prj

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// createArchive creates a tar.gz archive of the project
func createArchive(source string, target string) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// walk through the source directory
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(source, path)
		fmt.Printf("rel path: %s\n", relPath)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Size = info.Size()
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
