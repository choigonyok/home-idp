package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type gatewayRouter struct {
	router *mux.Router
}

func newRouter() *gatewayRouter {
	return &gatewayRouter{
		router: mux.NewRouter(),
	}
}

func (r *gatewayRouter) registerRoute(method, path string, f http.HandlerFunc) {
	r.router.NewRoute().Methods(method).Path(path).HandlerFunc(f)
}
