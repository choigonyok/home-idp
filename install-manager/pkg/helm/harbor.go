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
	return map[string]interface{}{
		"harborAdminPassword": env.Get("DEFAULT_REGISTRY_ADMIN_PASSWORD"),
		"expose": map[string]interface{}{
			"type": "clusterIP",
			"tls": map[string]interface{}{
				"enabled": true,
				"auto": map[string]interface{}{
					"commonName": "harbor",
				},
			},
		},
		"persistence": map[string]interface{}{
			"enabled": false,
			"persistentVolumeClaim": map[string]interface{}{
				"registry": map[string]interface{}{
					"storageClass": env.Get("GLOBAL_STORAGE_CLASS_NAME"),
					"size":         env.Get("GLOBAL_STORAGE_CLASS_SIZE"),
				},
			},
		},
		"trivy": map[string]interface{}{
			"enabled": false,
		},
	}
}