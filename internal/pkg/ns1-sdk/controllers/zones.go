package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/poc-golang-ns1/internal/pkg/ns1-sdk/sdk"
)

func GetAllZonesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	zones, _ := sdk.GetZones()
	json.NewEncoder(w).Encode(&zones)
}
