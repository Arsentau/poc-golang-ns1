// package controllers is responsible of handle requests, response and error response
package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	sdk "github.com/poc-golang-ns1/internal/pkg/ns1-sdk"
	e "github.com/poc-golang-ns1/pkg/errors"
)

// GetAllZonesHandler gets the zones using SDK implementation, if fails sends a custom HTTP error
func GetAllZonesHandler(w http.ResponseWriter, r *http.Request) {
	zones, err := sdk.GetZones()

	err = errors.New("aaaa aaaa 400 aaa")
	if err != nil {
		message := strings.Fields(err.Error())
		code, err := strconv.Atoi(message[2])
		if err != nil {
			code = 500
		}

		e.ErrorHTTPResponse(w, message[3], code)
		return
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(zones)
	if err != nil {
		e.ErrorHTTPResponse(w, err.Error(), 500)
	}
}
