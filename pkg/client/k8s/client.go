package client

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/choigonyok/idp/pkg/model"
	"github.com/choigonyok/idp/pkg/store"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func ptrInt64(v int64) *int64 { return &v }

type K8s struct {
	clients map[string]*kubernetes.Clientset
	db      *store.SQLiteStore
}

func NewK8sClient(v *viper.Viper, db *store.SQLiteStore) *K8s {
	auth := v.GetString("k8s.auth")
	switch auth {
	case "kubeconfig":
		k := K8s{clients: make(map[string]*kubernetes.Clientset)}
		infra, _ := db.ListInfra()
		for _, i := range *infra {
			if i.Platform.Kind == "K8s" {
				fmt.Println(i.Platform.K8sCluster.Context)
				k.AddClient(i.Platform.K8sCluster.Context)
				go k.StartEventCheck(i.Platform.K8sCluster.Context)
			}
		}
		return &k

	default:
		zap.S().Fatalf("%s: %s\n", "Invalid k8s auth method", auth)
		return nil
	}
}

func (k *K8s) AddClient(context string) {
	kubeconfig := filepath.Join("", "/Users/choigonyok/.kube/config")

	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}
	overrides := &clientcmd.ConfigOverrides{CurrentContext: context}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, overrides)
	config, err := clientConfig.ClientConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	k.clients[context] = cs
	return
}

func (k *K8s) LogService(conn *websocket.Conn, serviceName string, serviceID string, env string) ([]string, error) {
	ctx, err := k.db.GetContextFromServiceID(serviceID)
	if err != nil {
		return nil, err
	}
	req := k.clients[ctx].CoreV1().Pods(env).GetLogs(serviceName, &corev1.PodLogOptions{
		Follow: true,
	})

	stream, err := req.Stream(context.Background())
	if err != nil {
		fmt.Println("Log stream error:", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Log stream failed"))
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
	}()

	go func() {
		buf := make([]byte, 2000)
		for {
			n, err := stream.Read(buf)
			if err != nil {
				fmt.Println("Stream read error:", err)
				break
			}
			if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				fmt.Println("WebSocket write error:", err)
				break
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("메시지 읽기 에러:", err)
			break
		}
		fmt.Printf("클라이언트로부터 수신: %s", msg)
	}

	return nil, nil
}

func (k *K8s) StartEventCheck(context string) {
	factory := informers.NewSharedInformerFactory(k.clients[context], 0)
	serviceInformer := factory.Core().V1().Services().Informer()
	serviceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			svc := obj.(*v1.Service)
			fmt.Printf("Service create detected: %s/%s\n", svc.Namespace, svc.Name)
			infra, err := k.db.GetInfraFromContext(context)
			if err != nil {
				fmt.Println(err)
				return
			}
			k.db.CreateService(&model.Service{
				Name:         svc.Name,
				Status:       svc.Status.String(),
				Version:      "todo",
				Image:        "todo",
				Tag:          "todo",
				Namespace:    svc.Namespace,
				Env:          infra.Env,
				Assignee:     "todo",
				Git:          "todo",
				PlatformID:   infra.ID,
				PlatformKind: infra.Platform.Kind,
				PlatformName: infra.Platform.K8sCluster.Name,
			})
		},
		DeleteFunc: func(obj interface{}) {
			svc := obj.(*v1.Service)
			fmt.Printf("Service delete detected: %s/%s\n", svc.Namespace, svc.Name)
			// db.Create(&InfraService{
			// 	Name:      svc.Name,
			// 	Namespace: svc.Namespace,
			// 	Cluster:   clusterName,
			// 	CreatedAt: time.Now(),
			// })
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			oldSvc := oldObj.(*v1.Service)
			newSvc := newObj.(*v1.Service)
			fmt.Printf("Service update detected: %s -> %s\n", oldSvc.Name, newSvc.Name)
			// db.Create(&InfraService{
			// 	Name:      svc.Name,
			// 	Namespace: svc.Namespace,
			// 	Cluster:   clusterName,
			// 	CreatedAt: time.Now(),
			// })
		},
	})

	stopCh := make(chan struct{})
	factory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, serviceInformer.HasSynced) {
		log.Fatal("Timed out waiting for caches to sync")
	}

	<-stopCh
}
