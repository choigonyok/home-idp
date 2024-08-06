package server

import "net/http"

type GatewayServer struct {
	Http *http.Server
}

func (svr *GatewayServer) Close() error {
	return svr.Http.Close()
}

func (svr *GatewayServer) Run() {
	svr.Http.ListenAndServe()
}
