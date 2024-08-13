package http

import (
	"net/http"
	"strconv"
)

type GatewayServer struct {
	Http   *http.Server
	Router *GatewayRouter
}

func New(port int) *GatewayServer {
	r := NewRouter()

	return &GatewayServer{
		Http: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: r.Router,
		},
		Router: r,
	}
}

func (svr *GatewayServer) Run() {
	svr.Run()
}

func (svr *GatewayServer) Stop() {
	svr.Stop()
}
