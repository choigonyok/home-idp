package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/choigonyok/home-idp/pkg/object"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type KubeClient struct {
	Dynamic   *dynamic.DynamicClient
	ClientSet *kubernetes.Clientset
}

func NewKubeClient() *KubeClient {
	kubeconfig, _ := getKubeConfig()
	dc, _ := getDynamicClient(kubeconfig)
	cs, _ := kubernetes.NewForConfig(kubeconfig)
	return &KubeClient{
		Dynamic:   dc,
		ClientSet: cs,
	}
}

func (c *KubeClient) Set(i interface{}) {
	c.Dynamic = parseKubeClientFromInterface(i).Dynamic
	c.ClientSet = parseKubeClientFromInterface(i).ClientSet
}

func parseKubeClientFromInterface(i interface{}) *KubeClient {
	client := i.(*KubeClient)
	return client
}

func (c *KubeClient) ApplyManifest(manifest, resource, namespace string) error {
	gvk, obj := object.ParseObjectsFromManifest(manifest)

	mapIOP := make(map[string]any)
	yaml.Unmarshal([]byte(manifest), &mapIOP)
	gvr := schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: resource,
	}

	_, err := c.Dynamic.Resource(gvr).Namespace(namespace).Create(context.TODO(), &obj, metav1.CreateOptions{})
	return err
}

func getKubeConfig() (*rest.Config, error) {
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	return clientConfig.ClientConfig()
}

func ListServices(namespace string, client kubernetes.Interface) (*apiv1.ServiceList, error) {
	return client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
}

func getDynamicClient(kubeconfig *rest.Config) (*dynamic.DynamicClient, error) {
	client, err := dynamic.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println(err)
	}
	return client, nil
}

func (c *KubeClient) GetPodsByLabel(namespace, label string) []apiv1.Pod {
	pods, _ := c.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: label})

	return pods.Items
}

func (c *KubeClient) GetPods(namespace string) (*[]apiv1.Pod, error) {
	pods, err := c.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return &pods.Items, err
}

func (c *KubeClient) GetServices(namespace string) (*[]apiv1.Service, error) {
	svcs, err := c.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	return &svcs.Items, err
}

func (c *KubeClient) GetIngresses(namespace string) (*[]v1.Ingress, error) {
	ingresses, err := c.ClientSet.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	return &ingresses.Items, err
}

func (c *KubeClient) GetConfigmaps(namespace string) (*[]corev1.ConfigMap, error) {
	cms, err := c.ClientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	return &cms.Items, err
}
func (c *KubeClient) GetConfigmap(name, namespace string) (*corev1.ConfigMap, error) {
	return c.ClientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *KubeClient) GetSecrets(namespace string) (*[]corev1.Secret, error) {
	secrets, err := c.ClientSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	return &secrets.Items, err
}

func (c *KubeClient) GetNamespaces() (*[]apiv1.Namespace, error) {
	namespaces, err := c.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return &namespaces.Items, nil
}

func (c *KubeClient) GetServiceSelectors(name, namespace string) string {
	svc, err := c.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println("TEST GET SERVICE ERR:", err)
	}

	l := []string{}
	for k, v := range svc.Spec.Selector {
		l = append(l, k+"="+v)
	}

	return strings.Join(l, ",")
}

func (c *KubeClient) GetSecret(name, namespace, key string) []byte {
	secret, err := c.ClientSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	fmt.Println("TEST GET ARGOCD PASSWORD SECRET ERR:", err)
	return secret.Data[key]
}

func (c *KubeClient) IsServiceHealthy(service, namespace string) bool {
	label := c.GetServiceSelectors(service, namespace)
	pods := c.GetPodsByLabel(namespace, label)
	for _, pod := range pods {
		if pod.Status.Phase != corev1.PodRunning || !pod.Status.ContainerStatuses[0].Ready {
			return false
		}
	}

	return true
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
