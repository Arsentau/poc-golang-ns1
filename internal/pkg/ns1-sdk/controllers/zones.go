package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	e "github.com/poc-golang-ns1/internal/pkg/errors"
	"github.com/poc-golang-ns1/internal/pkg/ns1-sdk/sdk"
)

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
	fmt.Println(err)
	json.NewEncoder(w).Encode(&zones)
}
