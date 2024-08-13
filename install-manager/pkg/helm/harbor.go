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

// type HarborData struct {
// 	MetadataNamespace         string            `json:"namespace"`
// 	MetadataReleaseName       string            `json:"release_name"`
// 	IngressEnabled            bool              `json:"enabled"`
// 	IngressClassName          string            `json:"class_name"`
// 	IngressTls                bool              `json:"tls"`
// 	IngressAnnotation         map[string]string `json:"annotation"`
// 	OptRedisHA                bool              `json:"redis_ha"`
// 	OptControllerReplicas     int               `json:"controller_repl"`
// 	OptServerReplicas         int               `json:"server_repl"`
// 	OptRepoServerReplicas     int               `json:"repo_server_repl"`
// 	OptApplicationSetReplicas int               `json:"application_set_repl"`
// 	OptDomain                 string            `json:"domain"`
// }

func NewHarborClient(namespace, releaseName string) *Harbor {
	return &Harbor{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *Harbor) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("harbor/harbor:1.15.0", c.Namespace, c.ReleaseName, c.MakeVaulesWithOption())
	return nil
}

func (c *Harbor) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func (c *Harbor) MakeVaulesWithOption() map[string]interface{} {
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
