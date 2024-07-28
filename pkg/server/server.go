package server

import (
	"context"
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/storage"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	testPredicates = predicate.TypedFuncs[*unstructured.Unstructured]{
		CreateFunc: func(_ event.TypedCreateEvent[*unstructured.Unstructured]) bool {
			// no action
			return true
		},
		GenericFunc: func(_ event.TypedGenericEvent[*unstructured.Unstructured]) bool {
			// no action
			return false
		},
		// DeleteFunc: func(e event.TypedDeleteEvent[*unstructured.Unstructured]) bool {}
		UpdateFunc: func(e event.TypedUpdateEvent[*unstructured.Unstructured]) bool {
			// no action
			return false
		},
	}
)

type Server struct {
	Server        *grpc.Server
	StorageClient storage.StorageClient
}

func New(component config.Components) *Server {
	svr := &Server{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
	}
	log.Printf("Start configuring storage backend for the server...")
	client, _ := storage.NewClient(component)
	svr.StorageClient = client

	return svr
}

// func (svr *Server) Run() error {
// fmt.Println("SERVER IS RUNING NOW")
// fmt.Println()
// fmt.Println("LOOK FOR port")
// result1, found1, err := svr.Config.Get("port")
// if err != nil {
// 	fmt.Println("ERROR:", err)
// }
// if !found1 {
// 	fmt.Println("KEY_VALUE NOT FOUND")
// } else {
// 	fmt.Println("RESULT: ", result1)
// }

// fmt.Println()
// fmt.Println("LOOK FOR replicas")
// result2, found2, err := svr.Config.Get("replicas")
// if err != nil {
// 	fmt.Println("ERROR:", err)
// }
// if !found2 {
// 	fmt.Println("KEY_VALUE NOT FOUND")
// } else {
// 	fmt.Println("RESULT: ", result2)
// }

// fmt.Println()
// fmt.Println("LOOK FOR Port")
// result3, found3, err := svr.Config.Get("Port")
// if err != nil {
// 	fmt.Println("ERROR:", err)
// }
// if !found3 {
// 	fmt.Println("KEY_VALUE NOT FOUND")
// } else {
// 	fmt.Println("RESULT: ", result3)
// }

// fmt.Println()
// fmt.Println("LOOK FOR Replicas")
// result4, found4, err := svr.Config.Get("Replicas")
// if err != nil {
// 	fmt.Println("ERROR:", err)
// }
// if !found4 {
// 	fmt.Println("KEY_VALUE NOT FOUND")
// } else {
// 	fmt.Println("RESULT: ", result4)
// }
// run()
// 	return nil
// }

func run() {
	cfg, _ := kube.GetKubeConfig()
	var mgrOpt manager.Options

	mgrOpt = manager.Options{
		// LeaderElection:          leaderElectionEnabled,
		// LeaderElectionNamespace: leaderElectionNS,
		// LeaderElectionID:        leaderElectionID,
		// LeaseDuration:           &leaseDuration,
		// RenewDeadline:           renewDeadline,
	}

	mgr, _ := manager.New(cfg, mgrOpt)
	client := mgr.GetClient()
	// apis.AddToScheme(mgr.GetScheme())
	// controller.AddToManager(mgr, nil)
	r := &Reconciler{
		Client: &client,
	}

	opt := controller.Options{
		Reconciler: r,
	}
	c, _ := controller.New("test-controller", mgr, opt)

	for _, resource := range watchedResources() {
		obj := &unstructured.Unstructured{}
		obj.SetGroupVersionKind(schema.GroupVersionKind{
			Kind:    resource.Kind,
			Group:   resource.Group,
			Version: resource.Version,
		})

		handlerFunc := handler.TypedEnqueueRequestsFromMapFunc[*unstructured.Unstructured](func(_ context.Context, a *unstructured.Unstructured) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      a.GetName(),
					Namespace: a.GetNamespace(),
				}},
			}
		})

		c.Watch(source.Kind(mgr.GetCache(), obj, handlerFunc, testPredicates))
	}
	err := mgr.Start(signals.SetupSignalHandler())
	fmt.Println("MGR START ERR:", err)
}

type Reconciler struct {
	Client *client.Client
}

func (c *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	name := req.Name
	ns := req.Namespace
	fmt.Println()
	fmt.Println("-----RECONCILE START-----")
	fmt.Println("Name:", name)
	fmt.Println("NameSpace:", ns)

	fmt.Println("-----RECONCILE END-----")
	fmt.Println()

	return reconcile.Result{}, nil

}

func watchedResources() []schema.GroupVersionKind {
	gvks := []schema.GroupVersionKind{
		{Group: "apps", Version: "v1", Kind: "Deployment"},
		{Group: "", Version: "v1", Kind: "Pod"},
		// {Group: "", Version: "v1", Kind: name.SecretStr},
		// {Group: "", Version: "v1", Kind: name.SAStr},
		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.RoleBindingStr},
		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.RoleStr},
		// {Group: "admissionregistration.k8s.io", Version: "v1", Kind: name.MutatingWebhookConfigurationStr},
		// {Group: "admissionregistration.k8s.io", Version: "v1", Kind: name.ValidatingWebhookConfigurationStr},
		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.ClusterRoleStr},
		// {Group: "rbac.authorization.k8s.io", Version: "v1", Kind: name.ClusterRoleBindingStr},
		// {Group: "apiextensions.k8s.io", Version: "v1", Kind: name.CRDStr},
		// {Group: "policy", Version: "v1", Kind: name.PDBStr},
		// {Group: "autoscaling", Version: "v2", Kind: name.HPAStr},
	}
	return gvks
}
