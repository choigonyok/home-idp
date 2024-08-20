package manifest

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type KubeManifest struct {
	GVK       schema.GroupVersionKind
	Spec      KubeSpec
	Namespace string
	Name      string
}

type KubeSpec interface {
	Get() string
}

// resource := &manifest.Pod{} 구조체에 요청받은 리소스 내용을 정의하고
// manifest := resource.New()로 KubeManifest 구조체 생성 및 리턴
// m := manifest.GenerateManifest()로 실제 manifest 스트링 생성해서 전달
// kube.ApplyManifest(m)으로 클러스터에 배포

func (m *KubeManifest) GenerateManifest() string {
	// if m.GVK.Group == "" {
	// 	apiVersion := m.GVK.Version
	// 	fmt.Println("apiVersion: ", apiVersion)
	// } else {
	// 	apiVersion := m.GVK.Group + "/" + m.GVK.Version
	// 	fmt.Println("apiVersion: ", apiVersion)
	// }

	// kind := m.GVK.Kind
	// fmt.Println("kind:", kind)

	// metadataName := m.Name
	// metadataNamespace := m.Namespace
	// fmt.Printf("metadata:\n  name: %s\n  namespace: %s\n", metadataName, metadataNamespace)
	return m.Spec.Get()
}
