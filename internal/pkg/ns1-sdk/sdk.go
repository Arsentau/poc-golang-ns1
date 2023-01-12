// package sdk is responsible of the implementation of NS1 SDK
package sdk

import (
	"net/http"
	"time"

	c "github.com/poc-golang-ns1/internal/pkg/config"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

var config = c.GetConfig()

// getHTTPClient initializes the NS1 client
func getHTTPClient() *api.Client {
	httpClient := &http.Client{Timeout: time.Duration(config.Timeout) * time.Second}
	client := api.NewClient(httpClient, api.SetAPIKey(config.Ns1ApiKey))
	return client
}

// GetZones retrieves all zones or error
func GetZones() ([]*dns.Zone, int, error) {
	zones, response, err := getHTTPClient().Zones.List()
	return zones, response.StatusCode, err
}
