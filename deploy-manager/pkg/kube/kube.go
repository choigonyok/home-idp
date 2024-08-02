package kube

import (
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/object"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type KubeClient struct {
	Dynamic *dynamic.DynamicClient
	Set     *kubernetes.Clientset
}

func NewKubeClient() *KubeClient {
	kubeconfig, _ := kube.GetKubeConfig()
	dc, _ := kube.GetDynamicClient(kubeconfig)
	cs, _ := kube.NewClient(kubeconfig)
	return &KubeClient{
		Dynamic: dc,
		Set:     cs,
	}
}

func (c *KubeClient) ApplyManifest(manifest string) {
	gvk, obj := object.ParseObjectsFromManifest(manifest)

	mapIOP := make(map[string]any)
	yaml.Unmarshal([]byte(manifest), &mapIOP)
	kube.ApplyManifest("pods", "default", c.Dynamic, obj, gvk)
}
