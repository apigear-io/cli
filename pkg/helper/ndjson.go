package helper

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

// Scan scans a reader line by line and writes to the writer.
func ScanNDJSON[T any](r io.Reader) ([]T, error) {
	var items []T
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var item T
		line := scanner.Bytes()
		err := json.Unmarshal(line, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, scanner.Err()
}

// ScanFile scans a file line by line and writes to the writer.
func ScanNDJSONFile[T any](path string) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error closing file %s: %v", path, err)
			_ = err
		}
	}()
	return ScanNDJSON[T](f)
}
