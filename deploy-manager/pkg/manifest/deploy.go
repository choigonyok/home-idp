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

// 직접 개발 시, 도커파일 생성하면 자동으로 레지스트리 푸시 후 웹훅을 레지스트리에서 받아서 DeployManager가 업데이트된 이미지를 배포
// 퍼블릭 레지스트리 이미지 배포시, 대시보드에서 이미지, 레지스트리, 태그, 레플리카 개수, Deployment/statefulset/daemonset 등을 설정해서 배포
// 추가적으로 리소스 배포 기능
// - 도메인 설정 (Ingress)
// - Secret/ConfigMap
