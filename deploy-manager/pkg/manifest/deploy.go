package manifest

import "k8s.io/apimachinery/pkg/runtime/schema"

type DeploySpec struct {
	Type     string // default/stateful/replica/daemon
	Replicas int
	Image    string
}

func (p *DeploySpec) New(name, namespace string, gvk schema.GroupVersionKind) *KubeManifest {
	return &KubeManifest{
		Spec: &DeploySpec{
			Type:     p.Type,
			Replicas: p.Replicas,
			Image:    p.Image,
		},
		Name:      name,
		Namespace: namespace,
	}
}

func (p *DeploySpec) Get() string {
	return ""
}
