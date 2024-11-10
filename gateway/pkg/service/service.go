package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/choigonyok/home-idp/gateway/pkg/client"
	gatewayhttp "github.com/choigonyok/home-idp/gateway/pkg/http"
	"github.com/choigonyok/home-idp/gateway/pkg/progress"
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
	svr.Router.RegisterRoute(http.MethodGet, "/test2/", svc.TestHandler2())
	svr.Router.RegisterRoute(http.MethodPost, "/webhooks/harbor", svc.HarborWebhookHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/webhooks/github", svc.GithubWebhookHandler())
	svr.Router.RegisterRoutePrefix(http.MethodOptions, "/", svc.ApiOptionsHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/projects", svc.apiGetProjectListHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/roles", svc.apiGetRoleListHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/policies", svc.apiGetPoliciyListHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/users", svc.apiGetUserListHandler())

	svr.Router.RegisterRoute(http.MethodPost, "/api/policy", svc.apiPostPolicyHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/project", svc.apiPostProjectHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/role", svc.apiPostRoleHandler())
	svr.Router.RegisterRoute(http.MethodPut, "/api/users/{username}/role", svc.apiUpdateUserRoleHandler())

	svr.Router.RegisterRoute(http.MethodDelete, "/api/policies/{policyId}", svc.apiDeletePolicyHandler())
	svr.Router.RegisterRoute(http.MethodPut, "/api/policies/{policyId}", svc.apiUpdatePolicyHandler())

	svr.Router.RegisterRoute(http.MethodDelete, "/api/role", svc.apiDeleteRoleHandler())
	svr.Router.RegisterRoute(http.MethodDelete, "/api/user", svc.apiDeleteUserHandler())
	svr.Router.RegisterRoute(http.MethodDelete, "/api/project", svc.apiDeleteProjectHandler())

	svr.Router.RegisterRoute(http.MethodDelete, "/test0", svc.UninstallArgoCDHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/roles/{roleId}/policies", svc.apiGetRolePoliciesHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/policies/{policyId}", svc.apiGetPolicyHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/projects/{projectName}/users", svc.apiGetUsersInProjectHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/users/{userName}/role", svc.apiGetRoleHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/users/{userName}/dockerfiles", svc.apiGetDockerfilesHandler())

	svr.Router.RegisterRoute(http.MethodPost, "/api/dockerfile", svc.apiPostDockerfileHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/pod", svc.apiPostPodHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/traces/{traceId}", svc.apiGetTraceHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/dockerfiles/{dockerfileId}/trace", svc.apiGetDockerTraceHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/api/projects/{projectName}/secrets", svc.apiGetSecretsHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/projects/{projectName}/secret", svc.apiPostSecretHandler())
	svr.Router.RegisterRoute(http.MethodPost, "/api/projects/{projectName}/configmap", svc.apiPostConfigmapHandler())

	svr.Router.RegisterRoute(http.MethodPost, "/api/manifest", svc.apiPostManifestHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/projects/{projectName}/resources", svc.apiGetResourcesHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/projects/{projectName}/resources/configmaps/{configmapName}", svc.apiGetConfigmapHandler())
	svr.Router.RegisterRoute(http.MethodDelete, "/api/projects/{projectName}/resources/{resourceName}", svc.apiDeleteResourcesHandler())
	svr.Router.RegisterRoute(http.MethodPut, "/api/projects/{projectName}/users/{userName}", svc.apiPutUserHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/api/projects/{projectName}/configmaps", svc.apiGetConfigmapsHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/progress/{image}", svc.GetProgressHandler())
	svr.Router.RegisterRoutePrefix(http.MethodOptions, "/progress/", svc.ApiOptionsHandler())

	svr.Router.RegisterRoute(http.MethodGet, "/login", svc.SignHandler())
	svr.Router.RegisterRoute(http.MethodGet, "/github/callback", svc.CallbackHandler())
	return svc
}

func (svc *Gateway) Stop() {
	svc.ClientSet.DeployGrpcClient.Close()
	svc.ClientSet.RbacGrpcClient.Close()
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

	progress.Map = make(map[string][]*progress.Step)
	svc.Server.Run()
}

func (svc *Gateway) waitGatewayRunning() {
	for !svc.ClientSet.KubeClient.IsGatewayHealthy(env.Get("HOME_IDP_NAMESPACE")) {
		time.Sleep(time.Millisecond * 10)
		fmt.Println("TEST WAIT GATEWAY RUNNING")
	}
}

// {
// 	"policy": {
// 		"effect": "Allow",
// 		"target": "roles",
// 		"action": "CREATE"
// 	}
// }
