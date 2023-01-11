// Package utils contains utility functions for the NS1 API client
package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	c "github.com/poc-golang-ns1/internal/pkg/config"
)

var config = c.GetConfig()

// newRequest returns a new http.Request with the NS1 API key set in the header.
func newRequest(method string, endpoint string, body io.Reader) *http.Request {

	apiUrl := config.Ns1Host + endpoint

	req, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		log.Fatal(err)
	}

	// Set API key in header
	req.Header.Set("X-NSONE-Key", config.Ns1ApiKey)

	return req
}

// SendGetRequest is a generic function to send a GET request to the NS1 API.
func SendGetRequest(endpoint string, model interface{}) {
	// Create request
	method := "GET"
	req := newRequest(method, endpoint, nil)

	//Set client timeout
	http.DefaultClient.Timeout = time.Duration(config.Timeout) * time.Second

	// Send request
	res, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// Read response body
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	defer res.Body.Close()

	// Parse response
	parseErr := json.Unmarshal(body, &model)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}
