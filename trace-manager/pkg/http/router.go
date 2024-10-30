package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TraceManagerRouter struct {
	Router *mux.Router
}

func NewRouter() *TraceManagerRouter {
	r := &TraceManagerRouter{
		Router: mux.NewRouter(),
	}

	return r
}

func (r *TraceManagerRouter) RegisterRoute(method, path string, f http.HandlerFunc) {
	r.Router.NewRoute().Methods(method).Path(path).HandlerFunc(f)
}

func (r *TraceManagerRouter) RegisterRoutePrefix(method, path string, f http.HandlerFunc) {
	r.Router.NewRoute().Methods(method).PathPrefix(path).HandlerFunc(f)
}
