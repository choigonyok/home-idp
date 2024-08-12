package server

import (
	"net/http"
	"strconv"

	"github.com/choigonyok/home-idp/gateway/pkg/client"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
)

type Gateway struct {
	ClientSet *client.GatewayClientSet
	Server    *http.Server
	Router    *gatewayRouter
}

func New(port int, opts ...pkgclient.ClientOption) *Gateway {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	r := newRouter()

	return &Gateway{
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: r.router,
		},
		Router:    r,
		ClientSet: cs,
	}
}

func (svr *Gateway) Stop() {
	for _, cli := range svr.ClientSet.GrpcClient {
		cli.Close()
	}
	return
}

func (svr *Gateway) Start() {
	svr.Server.ListenAndServe()
	return
}
