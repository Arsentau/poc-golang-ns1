package models

import "net/http"

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
