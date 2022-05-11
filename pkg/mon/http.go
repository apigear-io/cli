package mon

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"objectapi/pkg/log"
)

func HttpPost(addr string, data any) error {
	client := &http.Client{}
	// create a new http request
	req, err := http.NewRequest("POST", addr, nil)
	if err != nil {
		log.Error("failed to create http request: ", err)
		return err
	}
	// set the content type to json
	req.Header.Set("Content-Type", "application/json")
	// loop through the events

	// encode the event into json
	json, err := json.Marshal(data)
	if err != nil {
		log.Errorf("failed to encode data: %v: %v", data, err)
		return err
	}
	// set the body of the request to the json encoded event
	req.Body = io.NopCloser(bytes.NewReader(json))
	// send the request to the monitor server
	resp, err := client.Do(req)
	if err != nil {
		log.Error("failed to send request: ", err)
		return err
	}
	// close the response body
	defer resp.Body.Close()
	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("failed to read response body: ", err)
		return err
	}
	// log the response
	log.Debug("response: ", string(body))
	return nil
}