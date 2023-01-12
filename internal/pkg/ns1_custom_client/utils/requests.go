// Package utils contains utility functions for the NS1 API client
package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	c "github.com/poc-golang-ns1/internal/pkg/config"
)

var config = c.GetConfig()

// SendGetRequest is a generic function to send a GET request to the NS1 API.
func SendGetRequest(endpoint string, model interface{}) error {
	// Create request
	method := "GET"
	req, err := newRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	//Set client timeout
	http.DefaultClient.Timeout = time.Duration(config.Timeout) * time.Second

	// Send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Handle 401, 403, 404, 500, etc.
	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	// Parse response
	if err := json.Unmarshal(body, &model); err != nil {
		return err
	}

	return nil
}

// newRequest returns a new http.Request with the NS1 API key set in the header.
func newRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {

	apiUrl := config.Ns1Host + endpoint

	req, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, err
	}

	// Set API key in header
	req.Header.Set("X-NSONE-Key", config.Ns1ApiKey)

	return req, nil
}
