package prj

import (
	"archive/tar"
	"compress/gzip"
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
	defer func() {
		if err := out.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close output file %s", target)
		}
	}()
	gw := gzip.NewWriter(out)
	defer func() {
		if err := gw.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close gzip writer")
		}
	}()
	tw := tar.NewWriter(gw)
	defer func() {
		if err := tw.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close tar writer")
		}
	}()
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
			defer func() {
				if err := file.Close(); err != nil {
					log.Error().Err(err).Msgf("failed to close file %s", path)
				}
			}()
			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
