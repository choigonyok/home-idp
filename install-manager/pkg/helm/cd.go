package helm

import (
	"strconv"

	"github.com/choigonyok/home-idp/pkg/helm"
	"gopkg.in/yaml.v3"
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
	RedisHA                bool
	ControllerReplicas     int
	ServerReplicas         int
	RepoServerReplicas     int
	ApplicationSetReplicas int
	Domain                 string
	Ingress                *ArgoCDIngressOption
}

type ArgoCDIngressOption struct {
	Enabled          bool
	IngressClassName string
	Annotation       *map[string]string
	Tls              bool
}

func NewArgoCDClient(namespace, releaseName string) *ArgoCD {
	return &ArgoCD{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *ArgoCD) Install(h *helm.HelmClient, opt *ArgoCDOption) error {
	h.Install("argo/argo-cd", c.Namespace, c.ReleaseName, c.GetOverrideValuesMap(opt))
	return nil
}

func (c *ArgoCD) Uninstall(h *helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func (c *ArgoCD) MakeVaulesWithOption(opt *ArgoCDOption) string {
	return `
redis-ha:
	enabled: ` + strconv.FormatBool(opt.RedisHA) + `

controller:
	replicas: ` + strconv.Itoa(opt.ControllerReplicas) + `

server:
	replicas: ` + strconv.Itoa(opt.ServerReplicas) + `

repoServer:
	replicas: ` + strconv.Itoa(opt.RepoServerReplicas) + `

applicationSet:
	replicas: ` + strconv.Itoa(opt.ApplicationSetReplicas) + `

global:
	domain: ` + opt.Domain + `

server:
	ingress:
		enabled: ` + strconv.FormatBool(opt.Ingress.Enabled) + `
		ingressClassName: ` + opt.Ingress.IngressClassName + `
		annotations:
			nginx.ingress.kubernetes.io/force-ssl-redirect: "true"
			nginx.ingress.kubernetes.io/ssl-passthrough: "true"
		tls: ` + strconv.FormatBool(opt.Ingress.Tls)
}

func (c *ArgoCD) GetOverrideValuesMap(opt *ArgoCDOption) map[string]interface{} {
	valuesYaml := c.MakeVaulesWithOption(opt)

	m := make(map[string]interface{})
	yaml.Unmarshal([]byte(valuesYaml), &m)

	return m
}
