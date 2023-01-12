// Package models contains the models used by the API handlers.
package models

import "net/http"

// Route is a struct that contains all the information to set up a route.
type Route struct {
	path    string
	method  string
	handler func(w http.ResponseWriter, r *http.Request)
}

func NewRoute(path string, method string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return &Route{path, method, handler}
}

func (r *Route) Path() string {
	return r.path
}

func (r *Route) Method() string {
	return r.method
}

func (r *Route) Handler() func(w http.ResponseWriter, r *http.Request) {
	return r.handler
}
