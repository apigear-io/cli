package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type HTTPSender struct {
	url    string
	client *http.Client
}

func NewHTTPSender(url string) *HTTPSender {
	return &HTTPSender{
		url: url,
		client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// SendValue sends the value as json to the sender.
func (s *HTTPSender) SendValue(value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.Send(data, "application/json")
}

// Send sends the data to the sender.
func (s *HTTPSender) Send(data []byte, contentType string) error {
	_, err := s.client.Post(s.url, contentType, bytes.NewReader(data))
	return err
}

// Write writes the event to the sender.
// This is used to implement the io.Writer interface.
// Data is send as json.
func (s *HTTPSender) Write(data []byte) (int, error) {
	return len(data), s.Send(data, "application/json")
}

func HttpPost(url string, contentType string, data []byte) error {
	c := &http.Client{
		Timeout: time.Second * 15,
	}
	_, err := c.Post(url, contentType, bytes.NewReader(data))
	return err
}

func HttpPostJson(url string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return HttpPost(url, "application/json", b)
}
