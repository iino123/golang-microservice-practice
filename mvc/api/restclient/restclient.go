package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	return client.Do(request)

}
