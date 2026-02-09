package foundation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPSender(t *testing.T) {
	t.Run("SendValue sends JSON", func(t *testing.T) {
		// Create test server
		var receivedData map[string]interface{}
		var receivedContentType string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedContentType = r.Header.Get("Content-Type")
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &receivedData)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		// Create sender and send data
		sender := NewHTTPSender(server.URL)
		data := map[string]interface{}{
			"name":  "test",
			"value": 42,
		}

		err := sender.SendValue(data)
		require.NoError(t, err)

		assert.Equal(t, "application/json", receivedContentType)
		assert.Equal(t, "test", receivedData["name"])
		assert.Equal(t, float64(42), receivedData["value"])
	})

	t.Run("Send sends raw data", func(t *testing.T) {
		var receivedData []byte
		var receivedContentType string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedContentType = r.Header.Get("Content-Type")
			receivedData, _ = io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		sender := NewHTTPSender(server.URL)
		data := []byte("test data")

		err := sender.Send(data, "text/plain")
		require.NoError(t, err)

		assert.Equal(t, "text/plain", receivedContentType)
		assert.Equal(t, data, receivedData)
	})

	t.Run("Write implements io.Writer", func(t *testing.T) {
		var receivedData []byte

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedData, _ = io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		sender := NewHTTPSender(server.URL)
		data := []byte("write test")

		n, err := sender.Write(data)
		require.NoError(t, err)
		assert.Equal(t, len(data), n)
		assert.Equal(t, data, receivedData)
	})

	t.Run("handles server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		sender := NewHTTPSender(server.URL)
		data := map[string]interface{}{"test": "data"}

		// Note: The current implementation doesn't check status codes,
		// so this won't error. This test documents current behavior.
		err := sender.SendValue(data)
		assert.NoError(t, err)
	})

	t.Run("handles invalid URL", func(t *testing.T) {
		sender := NewHTTPSender("http://invalid-url-that-does-not-exist:99999")
		data := map[string]interface{}{"test": "data"}

		err := sender.SendValue(data)
		assert.Error(t, err)
	})

	t.Run("SendValue with invalid data", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		sender := NewHTTPSender(server.URL)

		// Channels cannot be marshaled to JSON
		invalidData := make(chan int)

		err := sender.SendValue(invalidData)
		assert.Error(t, err)
	})

	t.Run("multiple sends", func(t *testing.T) {
		callCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			callCount++
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		sender := NewHTTPSender(server.URL)

		for i := 0; i < 5; i++ {
			err := sender.SendValue(map[string]int{"count": i})
			require.NoError(t, err)
		}

		assert.Equal(t, 5, callCount)
	})
}

func TestHttpPost(t *testing.T) {
	t.Run("successful post", func(t *testing.T) {
		var receivedData []byte
		var receivedContentType string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			receivedContentType = r.Header.Get("Content-Type")
			receivedData, _ = io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		data := []byte("test data")
		err := HttpPost(server.URL, "text/plain", data)
		require.NoError(t, err)

		assert.Equal(t, "text/plain", receivedContentType)
		assert.Equal(t, data, receivedData)
	})

	t.Run("post with different content types", func(t *testing.T) {
		contentTypes := []string{
			"application/json",
			"application/xml",
			"text/plain",
			"application/octet-stream",
		}

		for _, ct := range contentTypes {
			t.Run(ct, func(t *testing.T) {
				var receivedContentType string

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					receivedContentType = r.Header.Get("Content-Type")
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()

				err := HttpPost(server.URL, ct, []byte("data"))
				require.NoError(t, err)
				assert.Equal(t, ct, receivedContentType)
			})
		}
	})

	t.Run("handles invalid URL", func(t *testing.T) {
		err := HttpPost("http://invalid-url-that-does-not-exist:99999", "text/plain", []byte("data"))
		assert.Error(t, err)
	})

	t.Run("empty data", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			assert.Empty(t, body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		err := HttpPost(server.URL, "text/plain", []byte{})
		require.NoError(t, err)
	})
}

func TestHttpPostJson(t *testing.T) {
	t.Run("successful JSON post", func(t *testing.T) {
		var receivedData map[string]interface{}
		var receivedContentType string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			receivedContentType = r.Header.Get("Content-Type")
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &receivedData)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		data := map[string]interface{}{
			"name":  "test",
			"value": 42,
			"items": []string{"a", "b", "c"},
		}

		err := HttpPostJson(server.URL, data)
		require.NoError(t, err)

		assert.Equal(t, "application/json", receivedContentType)
		assert.Equal(t, "test", receivedData["name"])
		assert.Equal(t, float64(42), receivedData["value"])
	})

	t.Run("post struct", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		var receivedData Person

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &receivedData)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		person := Person{Name: "Alice", Age: 30}
		err := HttpPostJson(server.URL, person)
		require.NoError(t, err)

		assert.Equal(t, person.Name, receivedData.Name)
		assert.Equal(t, person.Age, receivedData.Age)
	})

	t.Run("post nested data", func(t *testing.T) {
		type Config struct {
			Settings map[string]interface{} `json:"settings"`
			Items    []int                  `json:"items"`
		}

		var receivedData Config

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &receivedData)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		config := Config{
			Settings: map[string]interface{}{
				"enabled": true,
				"count":   10,
			},
			Items: []int{1, 2, 3},
		}

		err := HttpPostJson(server.URL, config)
		require.NoError(t, err)

		assert.Equal(t, config.Items, receivedData.Items)
	})

	t.Run("handles marshal error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		// Channels cannot be marshaled to JSON
		invalidData := make(chan int)

		err := HttpPostJson(server.URL, invalidData)
		assert.Error(t, err)
	})

	t.Run("handles invalid URL", func(t *testing.T) {
		data := map[string]string{"test": "data"}
		err := HttpPostJson("http://invalid-url-that-does-not-exist:99999", data)
		assert.Error(t, err)
	})

	t.Run("post nil data", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			assert.Equal(t, []byte("null"), body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		err := HttpPostJson(server.URL, nil)
		require.NoError(t, err)
	})

	t.Run("post empty map", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			assert.Equal(t, []byte("{}"), body)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		err := HttpPostJson(server.URL, map[string]interface{}{})
		require.NoError(t, err)
	})
}
