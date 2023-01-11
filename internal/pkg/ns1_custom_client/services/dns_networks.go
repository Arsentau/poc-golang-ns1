// Package services contains the services for the NS1 Client.
package services

import (
	m "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/models"
	u "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/utils"
)

type NetworksService struct{}

// List returns the list of all networks from NS1.
func (n NetworksService) List() m.Networks {
	endpoint := "networks"
	var networks m.Networks
	u.SendGetRequest(endpoint, &networks)
	return networks
}
