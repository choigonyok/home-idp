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

// var (
//
//	testPredicates = predicate.TypedFuncs[*unstructured.Unstructured]{
//		CreateFunc: func(_ event.TypedCreateEvent[*unstructured.Unstructured]) bool {
//			// no action
//			return true
//		},
//		GenericFunc: func(_ event.TypedGenericEvent[*unstructured.Unstructured]) bool {
//			// no action
//			return false
//		},
//		// DeleteFunc: func(e event.TypedDeleteEvent[*unstructured.Unstructured]) bool {}
//		UpdateFunc: func(e event.TypedUpdateEvent[*unstructured.Unstructured]) bool {
//			// no action
//			return false
//		},
//	}
//
// )

// func run() {
// 	cfg, _ := kube.GetKubeConfig()
// 	var mgrOpt manager.Options

// 	mgrOpt = manager.Options{
// 		// LeaderElection:          leaderElectionEnabled,
// 		// LeaderElectionNamespace: leaderElectionNS,
// 		// LeaderElectionID:        leaderElectionID,
// 		// LeaseDuration:           &leaseDuration,
// 		// RenewDeadline:           renewDeadline,
// 	}

// 	mgr, _ := manager.New(cfg, mgrOpt)
// 	client := mgr.GetClient()
// 	// apis.AddToScheme(mgr.GetScheme())
// 	// controller.AddToManager(mgr, nil)
// 	r := &Reconciler{
// 		Client: &client,
// 	}

// 	opt := controller.Options{
// 		Reconciler: r,
// 	}
// 	c, _ := controller.New("test-controller", mgr, opt)

// 	for _, resource := range watchedResources() {
// 		obj := &unstructured.Unstructured{}
// 		obj.SetGroupVersionKind(schema.GroupVersionKind{
// 			Kind:    resource.Kind,
// 			Group:   resource.Group,
// 			Version: resource.Version,
// 		})

// 		handlerFunc := handler.TypedEnqueueRequestsFromMapFunc[*unstructured.Unstructured](func(_ context.Context, a *unstructured.Unstructured) []reconcile.Request {
// 			return []reconcile.Request{
// 				{NamespacedName: types.NamespacedName{
// 					Name:      a.GetName(),
// 					Namespace: a.GetNamespace(),
// 				}},
// 			}
// 		})

// 		c.Watch(source.Kind(mgr.GetCache(), obj, handlerFunc, testPredicates))
// 	}
// 	err := mgr.Start(signals.SetupSignalHandler())
// 	fmt.Println("MGR START ERR:", err)
// }

// type Reconciler struct {
// 	Client *client.Client
// }

// func (c *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
// 	name := req.Name
// 	ns := req.Namespace
// 	fmt.Println()
// 	fmt.Println("-----RECONCILE START-----")
// 	fmt.Println("Name:", name)
// 	fmt.Println("NameSpace:", ns)

// 	fmt.Println("-----RECONCILE END-----")
// 	fmt.Println()

// 	return reconcile.Result{}, nil

// }

// func watchedResources() []schema.GroupVersionKind {
// 	gvks := []schema.GroupVersionKind{
// 		{Group: "apps", Version: "v1", Kind: "Deployment"},
// 		{Group: "", Version: "v1", Kind: "Pod"},
// 		// {Group: "", Version: "v1", Kind: name.SecretStr},
// 		// {Group: "", Version: "v1", Kind: name.SAStr},
// 		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.RoleBindingStr},
// 		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.RoleStr},
// 		// {Group: "admissionregistration.k8s.io", Version: "v1", Kind: name.MutatingWebhookConfigurationStr},
// 		// {Group: "admissionregistration.k8s.io", Version: "v1", Kind: name.ValidatingWebhookConfigurationStr},
// 		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.ClusterRoleStr},
// 		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.ClusterRoleBindingStr},
// 		// {Group: "apiextensions.k8s.io", Version: "v1", Kind: name.CRDStr},
// 		// {Group: "policy", Version: "v1", Kind: name.PDBStr},
// 		// {Group: "autoscaling", Version: "v2", Kind: name.HPAStr},
// 	}
// 	return gvks
// }
