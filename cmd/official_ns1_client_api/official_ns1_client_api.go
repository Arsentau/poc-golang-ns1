package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/poc-golang-ns1/internal/pkg/config"
	"github.com/poc-golang-ns1/internal/pkg/official_api/router"
)

func main() {
	r := router.GetRouter()
	port := c.GetConfig().Port
	fmt.Println("Initializing server on " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
