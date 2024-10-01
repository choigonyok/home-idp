package helm

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type Harbor struct {
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
}

func NewHarborClient(namespace, releaseName string) *Harbor {
	return &Harbor{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *Harbor) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("harbor/harbor:1.15.0", c.Namespace, c.ReleaseName, harborOverrideValues())

	return nil
}

func (c *Harbor) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func harborOverrideValues() map[string]interface{} {
	host := env.Get("HOME_IDP_HARBOR_HOST")
	port := env.Get("HOME_IDP_HARBOR_PORT")
	schema := "http"
	tls := false
	if env.Get("HOME_IDP_HARBOR_TLS_ENABLED") == "true" {
		schema = "https"
		tls = true
	}
	url := schema + "://" + host + ":" + port

	return map[string]interface{}{
		"harborAdminPassword": env.Get("HOME_IDP_ADMIN_PASSWORD"),
		"externalURL":         url,
		"expose": map[string]interface{}{
			"tls": map[string]interface{}{
				"enabled": tls,
			},
			"ingress": map[string]interface{}{
				"hosts": map[string]interface{}{
					"core": env.Get("HOME_IDP_HARBOR_HOST"),
				},
				"className": "nginx",
			},
		},
		"persistence": map[string]interface{}{
			"enabled": false,
			// "persistentVolumeClaim": map[string]interface{}{
			// 	"registry": map[string]interface{}{
			// 		"storageClass": env.Get("HOME_IDP_STORAGE_CLASS_NAME"),
			// 		"size":         env.Get("HOME_IDP_STORAGE_CLASS_SIZE"),
			// 	},
			// },
		},
		"trivy": map[string]interface{}{
			"enabled": false,
		},
	}
}
