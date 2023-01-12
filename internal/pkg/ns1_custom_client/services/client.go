// Package services contains the services for the NS1 Client.
package services

// Ns1Client is the client for the NS1 API.
type Ns1Client struct {
	Networks NetworksService
	Zones    ZonesService
}
