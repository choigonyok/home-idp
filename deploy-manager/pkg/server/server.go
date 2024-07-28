package server

import (
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/object"
	"gopkg.in/yaml.v2"
)

func Test() {
	restConfig, _ := kube.GetKubeConfig()
	dc, _ := kube.GetDynamicClient(restConfig)

	gvk, obj := object.ParseObjectsFromManifest("RECEIVED_MANIFEST")

	// 1. 요청 수신
	// 2. 리소스 생성 권한을 체크
	// 3. YAML 파일이 유효한지 체크
	// 4. 베포

	mapIOP := make(map[string]any)
	yaml.Unmarshal([]byte("RECEIVED_MANIFES"), &mapIOP)
	kube.ApplyManifest("pods", "default", dc, obj, gvk)
}
