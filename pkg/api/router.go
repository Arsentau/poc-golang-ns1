package api

import (
	"github.com/gorilla/mux"
	m "github.com/poc-golang-ns1/pkg/api/models"
)

type Router struct {
	*mux.Router
}

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
