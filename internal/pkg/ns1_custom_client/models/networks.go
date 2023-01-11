// Package models contains the models for the NS1 Client.
package models

// Network represents a network in the NS1 Client.
type Network struct {
	Network_id int    `json:"network_id"`
	Name       string `json:"name"`
	Label      string `json:"label"`
}

// Networks represents a list of networks in the NS1 Client.
type Networks []Network
