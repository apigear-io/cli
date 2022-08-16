package rpc

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

func ReadJsonMessagesFromFile(fn string, sleep time.Duration, repeat int, emitter chan RpcMessage) {
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error("failed to open file ", fn, ": ", err)
		return
	}
	defer func() {
		file.Close()
		close(emitter)
	}()

	for i := 0; i < repeat; i++ {
		log.Debug("read file ", fn, ": ", i)
		err = ReadJsonMessages(file, sleep, emitter)
		if err != nil {
			log.Error("failed to read file: ", fn, ": ", err)
			return
		}
	}
}

func ReadJsonMessages(r io.Reader, sleep time.Duration, emitter chan RpcMessage) error {
	log.Debugf("read json messages from %v", r)
	scanner := bufio.NewScanner(r)
	id := uint64(0)
	for scanner.Scan() {
		line := scanner.Text()
		log.Infof("send: %s\n", line)
		// decode each line into an event using json
		var m RpcMessage
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(&m)
		if err != nil {
			log.Warnf("failed to decode line: %s: %v", line, err)
		}
		m.Version = "2.0"
		if m.Id == 0 {
			m.Id = id
			id++
		}
		time.Sleep(sleep)
		emitter <- m
	}
	return scanner.Err()
}
