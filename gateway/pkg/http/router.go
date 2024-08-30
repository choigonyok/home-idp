package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type GatewayRouter struct {
	Router *mux.Router
}

func NewRouter() *GatewayRouter {
	r := &GatewayRouter{
		Router: mux.NewRouter(),
	}

	return r
}

func (r *GatewayRouter) RegisterRoute(method, path string, f http.HandlerFunc) {
	r.Router.NewRoute().Methods(method).PathPrefix(path).HandlerFunc(f)
}
