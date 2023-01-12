// package controllers is responsible of handle requests, response and error response
package controllers

import (
	"encoding/json"
	"net/http"

	sdk "github.com/poc-golang-ns1/internal/pkg/ns1-sdk"
	e "github.com/poc-golang-ns1/pkg/errors"
)

// GetAllZonesHandler gets the zones using SDK implementation, if fails sends a custom HTTP error
func GetAllZonesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	zones, code, err := sdk.GetZones()

	if err != nil {
		message := err.Error()
		e.ErrorHTTPResponse(w, message, code)
		return
	}

	if err = json.NewEncoder(w).Encode(zones); err != nil {
		e.ErrorHTTPResponse(w, err.Error(), 500)
		return
	}
	w.WriteHeader(code)
}
