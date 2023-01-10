// package router is responsible of the creation of the endpoints
package router

import (
	c "github.com/poc-golang-ns1/internal/pkg/ns1-sdk/controllers"
	"github.com/poc-golang-ns1/pkg/api"
	m "github.com/poc-golang-ns1/pkg/api/models"
)

// getRoutes returns a list of routes with path, Method and HandlerFunc
func getRoutes() []m.Route {
	routes := []m.Route{
		*m.NewRoute("/zones", "GET", c.GetAllZonesHandler),
	}
	return routes
}

// GetRouter returns the initialized router for this API
func GetRouter() *api.Router {
	router := api.NewRouter(getRoutes())
	return router
}
