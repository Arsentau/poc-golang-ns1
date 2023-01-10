// This package handles reusable customized errors
package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorHttpResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}

// Custom http Response with the structure of ErrorHttpResponse and status code.
func ErrorHTTPResponse(w http.ResponseWriter, errorMessage string, code int) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	message := ErrorHttpResponse{
		Code:    code,
		Message: errorMessage,
	}
	e := json.NewEncoder(w).Encode(message)
	if e != nil {
		log.Panic("Error while encoding response")
	}
	return w
}
