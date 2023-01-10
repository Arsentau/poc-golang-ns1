package sdk

import (
	"log"
	"net/http"
	"time"

	c "github.com/poc-golang-ns1/internal/pkg/config"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

var config = c.GetConfig()

func getHTTPClient() *api.Client {
	httpClient := &http.Client{Timeout: time.Duration(config.Timeout) * time.Second}
	client := api.NewClient(httpClient, api.SetAPIKey(config.Ns1ApiKey))
	return client
}

func GetZones() ([]*dns.Zone, error) {
	zones, _, err := getHTTPClient().Zones.List()
	if err != nil {
		log.Fatal(err)
	}

	return zones, err
}
