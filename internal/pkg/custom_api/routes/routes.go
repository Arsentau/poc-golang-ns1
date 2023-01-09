package routes

import (
	"github.com/gorilla/mux"
	"github.com/poc-golang-ns1/internal/pkg/custom_api/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// TODO: Add routes
	router.HandleFunc("/networks", controllers.GetAllNetworks).Methods("GET")
	return router
}
