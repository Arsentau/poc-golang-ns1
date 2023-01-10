package router

import (
	c "github.com/poc-golang-ns1/internal/pkg/ns1-sdk/controllers"
	"github.com/poc-golang-ns1/pkg/api"
	m "github.com/poc-golang-ns1/pkg/api/models"
)

func getRoutes() []m.Route {
	routes := []m.Route{
		*m.NewRoute("/zones", "GET", c.GetAllZonesHandler),
	}
	return routes
}

func GetRouter() *api.Router {
	router := api.NewRouter(getRoutes())
	return router
}
