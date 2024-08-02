package kube

import (
	"context"
	"fmt"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetKubeConfig() (*rest.Config, error) {
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	return clientConfig.ClientConfig()
}

func NewClient(cfg *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(cfg)
}

func ListServices(namespace string, client kubernetes.Interface) (*apiv1.ServiceList, error) {
	return client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
}

func GetDynamicClient(kubeconfig *rest.Config) (*dynamic.DynamicClient, error) {
	client, err := dynamic.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println(err)
	}
	return client, nil
}

func ApplyManifest(resource, namespace string, client dynamic.Interface, obj *unstructured.Unstructured, gvk schema.GroupVersionKind) {
	gvr := schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: resource,
	}

	client.Resource(gvr).Namespace(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
}
