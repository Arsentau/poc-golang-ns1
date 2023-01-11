// Package services contains the services for the NS1 Client.
package services

import (
	m "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/models"
	u "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/utils"
)

type ZonesService struct{}

// List returns the list of all zones from NS1.
func (z ZonesService) List() m.Zones {
	endpoint := "zones"
	var zones m.Zones
	u.SendGetRequest(endpoint, &zones)
	return zones
}
