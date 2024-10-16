package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/choigonyok/home-idp/gateway/pkg/client"
	gatewayhttp "github.com/choigonyok/home-idp/gateway/pkg/http"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/env"
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

	svr.Router.RegisterRoute(http.MethodDelete, "/test0", svc.UninstallArgoCDHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/test2/", svc.TestHandler2())
	svr.Router.RegisterRoute(http.MethodPost, "/webhooks/harbor", svc.HarborWebhookHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/webhooks/github", svc.GithubWebhookHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/", svc.ApiPostHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/", svc.ApiGetHandler())
	svr.Router.RegisterRoute(http.MethodOptions, "/api/", svc.ApiOptionsHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/login", svc.LoginHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/login/callback", svc.LoginCallbackHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/signup", svc.SignUpHandler())
	return svc
}

func (svc *Gateway) Stop() {
	for _, cli := range svc.ClientSet.GrpcClient {
		cli.Close()
	}
	svc.Server.Stop()
}

func (svc *Gateway) Start() {
	go func() {
		svc.waitGatewayRunning()
		svc.ClientSet.GitClient.CreateAdminDir()
		if err := svc.ClientSet.GitClient.CreateGithubWebhook(); err != nil {
			fmt.Println("TEST GITHUB WEBHOOK CREATE ERR:", err)
		}
		fmt.Println("Clone URL:", svc.ClientSet.GitClient.GetRepositoryCloneURL())
	}()

	svc.Server.Run()
}

func (svc *Gateway) waitGatewayRunning() {
	for !svc.ClientSet.KubeClient.IsGatewayHealthy(env.Get("HOME_IDP_NAMESPACE")) {
		time.Sleep(time.Millisecond * 10)
		fmt.Println("TEST WAIT GATEWAY RUNNING")
	}
}
