package mon

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func HttpPost(url string, data any) error {
	client := &http.Client{}
	client.Timeout = time.Second * 15
	// create a new http request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create http request")
		return err
	}
	// set the content type to json
	req.Header.Set("Content-Type", "application/json")
	// loop through the events

	// encode the event into json
	json, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msgf("failed to encode data: %v", data)
		return err
	}
	// set the body of the request to the json encoded event
	req.Body = io.NopCloser(bytes.NewReader(json))
	// send the request to the monitor server
	log.Debug().Msgf("http post %s %s", url, json)
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to send request")
		return err
	}
	// close the response body
	defer resp.Body.Close()
	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
		return err
	}
	// log the response
	log.Debug().Msgf("response: %s", string(body))
	return nil
}
