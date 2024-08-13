package service

import (
	"net/http"

	"github.com/choigonyok/home-idp/gateway/pkg/client"
	gatewayhttp "github.com/choigonyok/home-idp/gateway/pkg/http"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
)

type Gateway struct {
	ClientSet *client.GatewayClientSet
	Server    *gatewayhttp.GatewayServer
}

func New(port int, opts ...pkgclient.ClientOption) *Gateway {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	svr := gatewayhttp.New(port)
	svc := &Gateway{
		Server:    svr,
		ClientSet: cs,
	}

	svr.Router.RegisterRoute(http.MethodPost, "/test0", svc.InstallArgoCDHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/test1", svc.TestHandler1())
	svr.Router.RegisterRoute(http.MethodPost, "/test2", svc.TestHandler2())
	svr.Router.RegisterRoute(http.MethodPost, "/test3", svc.TestHandler3())

	return svc
}

func (svc *Gateway) Stop() {
	for _, cli := range svc.ClientSet.GrpcClient {
		cli.Close()
	}
	svc.Server.Stop()
	return
}

func (svc *Gateway) Start() {
	svc.Server.Run()
	return
}
