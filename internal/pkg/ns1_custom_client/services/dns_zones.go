// Package services contains the services for the NS1 Client.
package services

import (
	m "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/models"
	u "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/utils"
)

type ZonesService struct{}

// List returns the list of all zones from NS1.
func (z ZonesService) List() (m.Zones, error) {
	endpoint := "zones"
	var zones m.Zones
	if err := u.SendGetRequest(endpoint, &zones); err != nil {
		return nil, err
	}
	return zones, nil
}
