// Package controllers contains the controllers for the NS1 API.
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	s "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/services"
)

// GetAllZones returns all networks from NS1.
func GetAllNetworks(w http.ResponseWriter, r *http.Request) {
	client := s.Ns1Client{}
	ns1Networks := client.Networks.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ns1Networks); err != nil {
		log.Println("Failed to get Networks - ", err)
	}
}
