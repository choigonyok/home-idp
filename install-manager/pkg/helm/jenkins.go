package helm

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type Jenkins struct {
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
}

func NewJenkinsClient(namespace, releaseName string) *Jenkins {
	return &Jenkins{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *Jenkins) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("jenkins/jenkins:5.4.2", c.Namespace, c.ReleaseName, jenkinsOverrideValues())
	return nil
}

func (c *Jenkins) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func jenkinsOverrideValues() map[string]interface{} {
	return map[string]interface{}{
		"fullnameOverride:": "home-idp-jenkins",
		"namespaceOverride": env.Get("GLOBAL_NAMESPACE"),
		"persistence": map[string]interface{}{
			"enabled": true,
		},
		"controller": map[string]interface{}{
			"admin": map[string]interface{}{
				"password": env.Get("DEFAULT_CI_ADMIN_PASSWORD"),
			},
			"csrf": map[string]interface{}{
				"defaultCrumbIssuer": map[string]interface{}{
					"enabled":            true,
					"proxyCompatability": true,
				},
			},
		},
	}
}
