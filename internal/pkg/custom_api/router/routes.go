package router

import (
	c "github.com/poc-golang-ns1/internal/pkg/custom_api/controllers"
	"github.com/poc-golang-ns1/pkg/api"
	m "github.com/poc-golang-ns1/pkg/api/models"
)

func getRoutes() []m.Route {
	routes := []m.Route{
		*m.NewRoute("/networks", "GET", c.GetAllNetworks),
	}
	return routes
}

func GetRouter() *api.Router {
	router := api.NewRouter(getRoutes())
	return router
}
