package helm

import (
	"fmt"
	"strconv"

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

type ArgoCDOption struct {
	Name                   string `json:"name"`
	Namespace              string `json:"namespace"`
	RedisHA                bool   `json:"redis_ha"`
	ControllerReplicas     int    `json:"controller_repl"`
	ServerReplicas         int    `json:"server_repl"`
	RepoServerReplicas     int    `json:"repo_server_repl"`
	ApplicationSetReplicas int    `json:"application_set_repl"`
	Domain                 string `json:"domain"`
	Ingress                *ArgoCDIngressOption
}

type ArgoCDIngressOption struct {
	Enabled          bool               `json:"enabled"`
	IngressClassName string             `json:"class_name"`
	Tls              bool               `json:"tls"`
	Annotation       *map[string]string `json:"annotation"`
}

type ArgoCDData struct {
	MetadataNamespace         string            `json:"namespace"`
	MetadataReleaseName       string            `json:"release_name"`
	IngressEnabled            bool              `json:"enabled"`
	IngressClassName          string            `json:"class_name"`
	IngressTls                bool              `json:"tls"`
	IngressAnnotation         map[string]string `json:"annotation"`
	OptRedisHA                bool              `json:"redis_ha"`
	OptControllerReplicas     int               `json:"controller_repl"`
	OptServerReplicas         int               `json:"server_repl"`
	OptRepoServerReplicas     int               `json:"repo_server_repl"`
	OptApplicationSetReplicas int               `json:"application_set_repl"`
	OptDomain                 string            `json:"domain"`
}

func NewArgoCDClient(namespace, releaseName string) *ArgoCD {
	return &ArgoCD{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *ArgoCD) Install(h *helm.HelmClient, opt *ArgoCDOption) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("argo/argo-cd:7.4.0", c.Namespace, c.ReleaseName, c.MakeVaulesWithOption(opt))
	return nil
}

func (c *ArgoCD) Upgrade(h *helm.HelmClient, opt *ArgoCDOption) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Upgrade("argo/argo-cd:7.4.0", c.Namespace, c.ReleaseName, c.MakeVaulesWithOption(opt))

	return nil
}

func (c *ArgoCD) Uninstall(h *helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func (c *ArgoCD) MakeVaulesWithOption(opt *ArgoCDOption) map[string]interface{} {
	return map[string]interface{}{
		"crds": map[string]interface{}{
			"keep": "false",
		},
		"redis-ha": map[string]interface{}{
			"enabled": strconv.FormatBool(opt.RedisHA),
		},
		"controller": map[string]interface{}{
			"replicas": strconv.Itoa(opt.ControllerReplicas),
		},
		"repoServer": map[string]interface{}{
			"replicas": strconv.Itoa(opt.RepoServerReplicas),
		},
		"applicationSet": map[string]interface{}{
			"replicas": strconv.Itoa(opt.ApplicationSetReplicas),
		},
		"global": map[string]interface{}{
			"domain": opt.Domain,
		},
		"server": map[string]interface{}{
			"replicas": strconv.Itoa(opt.ServerReplicas),
			"ingress": map[string]interface{}{
				"enabled":          strconv.FormatBool(opt.Ingress.Enabled),
				"ingressClassName": opt.Ingress.IngressClassName,
				"annotations": map[string]interface{}{
					"nginx.ingress.kubernetes.io/force-ssl-redirect": "true",
					"nginx.ingress.kubernetes.io/ssl-passthrough":    "true",
				},
				"tls": strconv.FormatBool(opt.Ingress.Tls),
			},
		},
	}
}
