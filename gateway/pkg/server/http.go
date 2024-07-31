package server

import "net/http"

type GatewayServer struct {
	Server *http.Server
}

func (svr *GatewayServer) Close() error {
	return svr.Server.Close()
}

func (svr *GatewayServer) Run() {
	svr.Server.ListenAndServe()
}
