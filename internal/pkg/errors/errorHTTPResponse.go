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

func ErrorHTTPResponse(w http.ResponseWriter, err string, code int) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	message := ErrorHttpResponse{
		Code:    code,
		Message: err,
	}
	e := json.NewEncoder(w).Encode(message)
	if e != nil {
		log.Panic("Error while encoding response")
	}
	return w
}
