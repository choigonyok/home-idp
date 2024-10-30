package service

import (
	util "net/http"

	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/trace-manager/pkg/client"
	"github.com/choigonyok/home-idp/trace-manager/pkg/http"
)

type TraceManager struct {
	ClientSet *client.TraceManagerClientSet
	Server    *http.TraceManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *TraceManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}
	svr := http.New(port)
	svc := &TraceManager{
		Server:    svr,
		ClientSet: cs,
	}

	svr.Router.RegisterRoute(util.MethodPost, "/api/span", svc.apiPostSpanHandler())
	svr.Router.RegisterRoute(util.MethodPut, "/api/span", svc.apiPutSpanHandler())
	svr.Router.RegisterRoute(util.MethodGet, "/api/traces/{traceId}", svc.apiGetTraceHandler())

	return svc
}

func (svc *TraceManager) Stop() {
	svc.Server.Stop()
}

func (svc *TraceManager) Start() {
	svc.Server.Run()
}
