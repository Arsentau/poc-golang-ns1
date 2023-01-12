// Package controllers contains the controllers for the NS1 API.
package controllers

import (
	"encoding/json"
	"net/http"

	s "github.com/poc-golang-ns1/internal/pkg/ns1_custom_client/services"
	e "github.com/poc-golang-ns1/pkg/errors"
)

// GetAllZones returns all networks from NS1.
func GetAllNetworks(w http.ResponseWriter, r *http.Request) {
	client := s.Ns1Client{}
	ns1Networks, err := client.Networks.List()

	if err != nil {
		e.ErrorHTTPResponse(w, "Failed to get Networks", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ns1Networks); err != nil {
		e.ErrorHTTPResponse(w, err.Error(), 500)
		return
	}
}
