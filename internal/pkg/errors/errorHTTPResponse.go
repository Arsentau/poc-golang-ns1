package errors

import (
	"encoding/json"
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
	json.NewEncoder(w).Encode(message)
	return w
}
