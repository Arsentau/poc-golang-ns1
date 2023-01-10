package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorHttpResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}

func HandlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

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
