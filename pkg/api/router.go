// Package api provides generic methods to set up the API of the application.
package api

import (
	"github.com/gorilla/mux"
	m "github.com/poc-golang-ns1/pkg/api/models"
)

// Router is a wrapper for the gorilla/mux router.
type Router struct {
	*mux.Router
}

// NewRouter returns a new router with the routes provided.
func NewRouter(routes []m.Route) *Router {
	router := &Router{mux.NewRouter().StrictSlash(true)}
	for _, route := range routes {
		router.
			Methods(route.Method()).
			Path(route.Path()).
			HandlerFunc(route.Handler())
	}
	return router
}
