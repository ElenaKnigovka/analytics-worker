package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/analytics-worker/models"
)

func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func sendRequest(url string, method string, data interface{}) (*http.Response, error) {
	client := getHTTPClient()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req.Body = jsonData
		req.Header.Set("Content-Type", "application/json")
	}
	return client.Do(req)
}

func parseResponse(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}

func handleRequestError(err error) {
	if err != nil {
		log.Printf("request error: %v\n", err)
	}
}

func validateModel(m models.Model) error {
	if m.ID == 0 {
		return errors.New("model id is required")
	}
	if m.CreatedAt.IsZero() {
		return errors.New("model created at is required")
	}
	return nil
}