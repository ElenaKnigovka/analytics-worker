package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/analytics-worker/models"
)

func getHTTPRequest(url string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non-200 status code: " + resp.Status)
	}
	return resp, nil
}

func unmarshalJSON(data []byte, target interface{}) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return err
	}
	return nil
}

func getModelsFromResponse(resp *http.Response) ([]models.DataPoint, error) {
	defer resp.Body.Close()
	var dataPoints []models.DataPoint
	err := unmarshalJSON(json.NewDecoder(resp.Body).Bytes(), &dataPoints)
	if err != nil {
		return nil, err
	}
	return dataPoints, nil
}

func handleHTTPError(err error) {
	if err != nil {
		log.Printf("HTTP error: %v", err)
	}
}