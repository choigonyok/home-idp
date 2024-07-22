package cli

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	kubescheme "k8s.io/client-go/kubernetes/scheme"
)

type CLIClient interface {
	GetPods(ctx context.Context, namespace string, opts metav1.ListOptions) ([]v1.Pod, error)
	ApplyManifests(ctx context.Context, namespace string, opts metav1.ListOptions) ([]v1.Pod, error)
}

type Client struct {
	Kube       kubernetes.Interface
	kubeClient *kubernetes.Clientset
}

// func (c *Client) GetPods(ctx context.Context, namespace string, opts metav1.ListOptions) ([]v1.Pod, error) {
// 	podList, err := c.Kube.CoreV1().Pods(namespace).List(ctx, opts)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to retrieve Pods: %v", err)
// 	}
// 	return podList, nil
// }

func (c *Client) ApplyManifests(ctx context.Context, namespace string, opts metav1.ListOptions) error {
	return nil
}

func idpScheme() *runtime.Scheme {

	scheme := runtime.NewScheme()
	utilruntime.Must(kubescheme.AddToScheme(scheme))
	// utilruntime.Must(clientnetworkingalpha.AddToScheme(scheme))
	return scheme
	// runtime.Must(mcs.AddToScheme(scheme))
	// utilruntime.Must(clientnetworkingbeta.AddToScheme(scheme))
	// utilruntime.Must(clientsecurity.AddToScheme(scheme))
	// utilruntime.Must(clienttelemetry.AddToScheme(scheme))
	// utilruntime.Must(clientextensions.AddToScheme(scheme))
	// utilruntime.Must(gatewayapi.AddToScheme(scheme))
	// utilruntime.Must(gatewayapibeta.AddToScheme(scheme))
	// utilruntime.Must(gatewayapiv1.AddToScheme(scheme))
	// utilruntime.Must(apiextensionsv1.AddToScheme(scheme))
}

// func KubernetesClients(kubeClient kube.CLIClient, l clog.Logger) (kube.CLIClient, client.Client, error) {
// 	client, err := client.New(kubeClient.RESTConfig(), client.Options{Scheme: kube.IstioScheme})
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	if err := k8sversion.IsK8VersionSupported(kubeClient, l); err != nil {
// 		return nil, nil, fmt.Errorf("check minimum supported Kubernetes version: %v", err)
// 	}
// 	return kubeClient, client, nil
// }
