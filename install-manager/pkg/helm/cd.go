package helm

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type CDOption struct {
	Values map[string]interface{}
}

type CD interface {
	Install(h *helm.HelmClient, opt CDOption) error
	Uninstall(h *helm.HelmClient) error
	MakeVaulesWithOption(CDOption) map[string]interface{}
}

type ArgoCD struct {
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
}

func NewArgoCDClient(namespace, releaseName string) *ArgoCD {
	return &ArgoCD{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *ArgoCD) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("argo/argo-cd:7.4.0", c.Namespace, c.ReleaseName, c.MakeVaulesWithOption())
	return nil
}

// func (c *ArgoCD) Upgrade(h *client.InstallManagerHelmClient, opt *ArgoCDOption) error {
// 	fmt.Println("RELEASE: ", c.ReleaseName)
// 	fmt.Println("NAMESPACE: ", c.Namespace)
// h.Upgrade("argo/argo-cd:7.4.0", c.Namespace, c.ReleaseName, c.MakeVaulesWithOption())

// 	return nil
// }

func (c *ArgoCD) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func (c *ArgoCD) MakeVaulesWithOption() map[string]interface{} {
	return map[string]interface{}{
		"crds": map[string]interface{}{
			"keep": "false",
		},
		"redis-ha": map[string]interface{}{
			"enabled": false,
		},
		"controller": map[string]interface{}{
			"replicas": 1,
		},
		"repoServer": map[string]interface{}{
			"replicas": 1,
		},
		"applicationSet": map[string]interface{}{
			"replicas": 1,
		},
		"configs": map[string]interface{}{
			"secret": map[string]interface{}{
				"argocdServerAdminPassword": env.Get("DEFAULT_CD_ADMIN_PASSWORD"),
			},
		},
		"server": map[string]interface{}{
			"replicas": 1,
			"ingress": map[string]interface{}{
				"enabled": false,
			},
		},
	}
}
