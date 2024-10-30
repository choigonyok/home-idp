package http

import (
	"net/http"
	"strconv"
)

type TraceManagerServer struct {
	Http   *http.Server
	Router *TraceManagerRouter
}

func New(port int) *TraceManagerServer {
	r := NewRouter()

	return &TraceManagerServer{
		Http: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: r.Router,
		},
		Router: r,
	}
}

func (svr *TraceManagerServer) Run() {
	svr.Http.ListenAndServe()
}

func (svr *TraceManagerServer) Stop() {
	svr.Http.Close()
}
