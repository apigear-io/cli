package rpc

import (
	"apigear/pkg/log"
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

func ReadJsonMessagesFromFile(fn string, sleepDuration time.Duration, emitter chan RpcMessage) {
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error("failed to open file ", fn, ": ", err)
		return
	}
	defer file.Close()

	err = ReadJsonMessages(file, sleepDuration, emitter)
	if err != nil {
		log.Error("failed to read file: ", fn, ": ", err)
		return
	}
}

func ReadJsonMessages(r io.Reader, sleepDuration time.Duration, emitter chan RpcMessage) error {
	scanner := bufio.NewScanner(r)
	id := uint64(0)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("read line: ", line)
		// decode each line into an event using json
		var m RpcMessage
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(&m)
		if err != nil {
			log.Fatalf("failed to decode line: %s: %v", line, err)
		}
		m.Version = "2.0"
		if m.Id == 0 {
			m.Id = id
			id++
		}
		time.Sleep(sleepDuration)
		emitter <- m
	}
	close(emitter)
	err := scanner.Err()
	if err != nil {
		log.Error("failed to read messages: ", err)
		return err
	}
	return nil
}
