package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/poc-golang-ns1/internal/pkg/custom_api/router"
)

func main() {
	port := ":8080"
	host := "localhost"
	r := router.GetRouter()
	fmt.Println("Initializing server on " + host + port)
	log.Fatal(http.ListenAndServe(port, r))
}
