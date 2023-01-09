package controllers

import (
	"encoding/json"
	"net/http"
)

func GetAllNetworks(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]string{"network1", "network2"})
}
