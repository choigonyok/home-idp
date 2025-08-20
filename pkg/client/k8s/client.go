package client

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ptrInt64(v int64) *int64 { return &v }

type K8s struct {
	clientset *kubernetes.Clientset
}

func NewFromKubeconfig() *K8s {
	kubeconfig := filepath.Join("", "/Users/choigonyok/.kube/config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		// fallback to in-cluster
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
			&clientcmd.ConfigOverrides{}).ClientConfig()
		if err != nil {
			panic(err)
		}
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return &K8s{clientset: cs}
}

func (k *K8s) LogService(conn *websocket.Conn, serviceName string, env string) ([]string, error) {
	req := k.clientset.CoreV1().Pods(env).GetLogs(serviceName, &corev1.PodLogOptions{
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
