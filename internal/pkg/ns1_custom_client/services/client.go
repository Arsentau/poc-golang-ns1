// Package services contains the services for the NS1 Client.
package services

type Ns1Client struct {
	Networks NetworksService
	Zones    ZonesService
}
