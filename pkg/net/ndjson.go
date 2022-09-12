package net

import (
	"bufio"
	"io"
	"os"
	"time"
)

func ScanJsonDelimitedFile(fn string, sleep time.Duration, repeat int, emitter chan []byte) {
	log.Debugf("scan file %s", fn)
	// read file line by line using scanner
	file, err := os.Open(fn)
	defer func() {
		file.Close()
		close(emitter)
	}()
	if err != nil {
		log.Error("failed to open file ", fn, ": ", err)
		return
	}
	if repeat == 0 {
		repeat = 1
	}
	for i := 0; i < repeat; i++ {
		log.Debugf("read json messages from file %s", fn)
		err = readJsonLines(file, sleep, emitter)
		if err != nil {
			log.Errorf("failed to read messages from file %s: %v", fn, err)
			return
		}
	}
}

func readJsonLines(r io.Reader, sleep time.Duration, emitter chan []byte) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debugf("line: %s", line)
		time.Sleep(sleep)
		log.Debugf("emit message: %s", line)
		emitter <- []byte(line)
	}
	return scanner.Err()
}
