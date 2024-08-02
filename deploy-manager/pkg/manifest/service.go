package manifest

import (
	"strconv"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ServiceSpec struct {
	Selector   string // "group/version/kind/name/namespace"
	Type       string // ClusterIP/NodePort/LoadBalancer
	NodePort   int
	Port       int
	TargetPort int
}

func (s *ServiceSpec) New(name, namespace string, gvk schema.GroupVersionKind) *KubeManifest {
	return &KubeManifest{
		GVK: gvk,
		Spec: &ServiceSpec{
			Selector:   s.Selector,
			Type:       s.Type,
			NodePort:   s.NodePort,
			Port:       s.Port,
			TargetPort: s.TargetPort,
		},
		Name:      name,
		Namespace: namespace,
	}
}

func (p *ServiceSpec) Get() string {
	spec :=
		`type: ` + p.Type + `
		ports:
			port: ` + strconv.Itoa(p.Port) + `
			targetPort ` + strconv.Itoa(p.TargetPort) + `
			nodePort: ` + strconv.Itoa(p.NodePort) + `
		selector:	
		`

	return spec
}
