package kube

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetKubeConfig() (*kubernetes.Clientset, *rest.Config, string, error) {
	var err error

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return nil, nil, "", err
	}

	config, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, nil, "", err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, "", err
	}
	return client, config, namespace, nil
}

func ListServices(namespace string, client kubernetes.Interface) (*apiv1.ServiceList, error) {
	return client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
}
