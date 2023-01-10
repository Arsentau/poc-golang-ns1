package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllNetworks(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Client
	ns1_networks := []string{"network1", "network2"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ns1_networks); err != nil {
		log.Println("Failed to get Networks - ", err)
	}
}
