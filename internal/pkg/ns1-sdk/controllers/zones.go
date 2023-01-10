// package controllers is responsible of handle requests, response and error response
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/poc-golang-ns1/internal/pkg/ns1-sdk/sdk"
	e "github.com/poc-golang-ns1/pkg/errors"
)

// GetAllZonesHandler gets the zones using SDK implementation, if fails sends a custom HTTP error
func GetAllZonesHandler(w http.ResponseWriter, r *http.Request) {
	zones, err := sdk.GetZones()

	if err != nil {
		message := strings.Fields(err.Error())
		code, err := strconv.Atoi(message[2])
		if err != nil {
			code = 500
		}
		e.ErrorHTTPResponse(w, message[1], code)
		return
	}
	w.Header().Set("Content-type", "application/json")
	e := json.NewEncoder(w).Encode(&zones)
	if e != nil {
		log.Panic("Error while encoding response")
	}
}
