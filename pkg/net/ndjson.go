package net

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/apigear-io/cli/pkg/log"
)

func ScanJsonDelimitedFile(fn string, sleep time.Duration, repeat int, emitter chan []byte) {
	log.Debug().Msgf("scan file %s", fn)
	// read file line by line using scanner
	file, err := os.Open(fn)
	defer func() {
		file.Close()
		close(emitter)
	}()
	if err != nil {
		log.Error().Err(err).Msgf("open file %s", fn)
		return
	}
	if repeat == 0 {
		repeat = 1
	}
	for i := 0; i < repeat; i++ {
		log.Debug().Msgf("read json messages from file %s", fn)
		err = readJsonLines(file, sleep, emitter)
		if err != nil {
			log.Error().Msgf("read messages from file %s: %v", fn, err)
			return
		}
	}
}

func readJsonLines(r io.Reader, sleep time.Duration, emitter chan []byte) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug().Msgf("line: %s", line)
		time.Sleep(sleep)
		log.Debug().Msgf("emit message: %s", line)
		emitter <- []byte(line)
	}
	return scanner.Err()
}
