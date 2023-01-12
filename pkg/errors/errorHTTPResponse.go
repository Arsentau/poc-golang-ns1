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
	message := ErrorHttpResponse{
		Code:    code,
		Message: errorMessage,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		w.WriteHeader(500)
		defer handlePanic()
		log.Panic(err.Error())
		return w
	}
	return w
}

// If a panic occurs, this func logs and recover the execution
func handlePanic() {
	if r := recover(); r != nil {
		log.Println("Recovering from panic:", r)
	}
}
