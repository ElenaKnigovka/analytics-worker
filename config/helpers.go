package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/analytics-worker/models"
)

func handleError(err error) {
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func sendRequest(url string, method string, data interface{}) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	var req *http.Request
	var err error

	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, url, nil)
	case http.MethodPost:
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, jsonData)
	default:
		return nil, errors.New("unsupported method")
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func parseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func trackEvent(event models.Event) error {
	url := "https://example.com/track"
	data := event
	resp, err := sendRequest(url, http.MethodPost, data)
	if err != nil {
		return err
	}
	var result models.Result
	err = parseResponse(resp, &result)
	if err != nil {
		return err
	}
	return nil
}