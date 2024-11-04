package grpc

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/git"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type DeployServer struct {
	pb.DeployServer
	KubeClient *kube.DeployManagerKubeClient
	GitClient  *git.DeployManagerGitClient
}

func (svr *DeployServer) Deploy(ctx context.Context, in *pb.DeployRequest) (*pb.DeployReply, error) {
	fmt.Println("TEST START DEPLOY ARGOCD APPLICATION")
	manifest := svr.GitClient.GetManifestFromGithub(in.Filepath)

	if err := svr.KubeClient.DeployManifest(manifest); err != nil {
		fmt.Println("TEST DEPLOY MANIFEST ERR:", err)
		return &pb.DeployReply{
			Succeed: false,
		}, nil
	}

	fmt.Println("TEST END DEPLOY ARGOCD APPLICATION")
	return &pb.DeployReply{
		Succeed: true,
	}, nil
}

func (svr *DeployServer) DeployPod(ctx context.Context, in *pb.DeployPodRequest) (*emptypb.Empty, error) {

	port := in.GetPod().GetContainerPort()
	img := in.GetPod().GetImage()
	name := in.GetPod().GetName()
	ns := in.GetPod().GetNamespace()

	svr.Test2(name, ns, img, port)

	return nil, nil
}

func (svr *DeployServer) Test(name, namespace, image, port string) unstructured.Unstructured {
	fmt.Println("TEST START")
	obj := unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	})

	p, _ := strconv.Atoi(port)

	// Deployment를 표현하는 unstructured 객체 생성
	obj.Object = map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"name":      name,
			"namespace": namespace,
		},
		"spec": map[string]interface{}{
			"replicas": int64(1),
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app": "example-app",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app": "example-app",
					},
				},
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{
							"name":  "app-container",
							"image": image,
							"ports": []interface{}{
								map[string]interface{}{
									"containerPort": int64(p),
								},
							},
						},
						map[string]interface{}{
							"name":  "envoy",
							"image": "envoyproxy/envoy:v1.18.3",
							"ports": []interface{}{
								map[string]interface{}{
									"containerPort": int64(10000),
								},
								map[string]interface{}{
									"containerPort": int64(9901),
								},
							},
							"volumeMounts": []interface{}{
								map[string]interface{}{
									"name":      "envoy-config",
									"mountPath": "/etc/envoy",
								},
							},
							"args": []interface{}{
								"-c",
								"/etc/envoy/envoy.yaml",
							},
						},
					},
					"volumes": []interface{}{
						map[string]interface{}{
							"name": "envoy-config",
							"configMap": map[string]interface{}{
								"name": "envoy-config",
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("TEST END")
	return obj
}

func (svr *DeployServer) Test2(name, namespace, image, port string) {
	fmt.Println("TEST2 START")
	obj := svr.Test(name, namespace, image, port)
	gvr := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	_, err := svr.KubeClient.Client.Dynamic.Resource(gvr).Namespace(namespace).Apply(context.TODO(), name, &obj, v1.ApplyOptions{
		FieldManager: "deploy-manager",
	})
	if err != nil {
		fmt.Println("TEST2 ERR:", err)
	}

	fmt.Println("TEST2 END")
}

func (svr *DeployServer) DeploySecret(ctx context.Context, in *pb.DeploySecretRequest) (*emptypb.Empty, error) {
	ns := in.GetNamespace()
	pusher := in.GetPusher()
	secrets := in.GetSecrets()

	obj := svr.GetSecretObject(pusher, ns, secrets)
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "secrets",
	}

	_, err := svr.KubeClient.Client.Dynamic.Resource(gvr).Namespace(ns).Apply(context.TODO(), "secret-"+pusher, &obj, v1.ApplyOptions{
		FieldManager: "deploy-manager",
	})
	if err != nil {
		fmt.Println("ERR DEPLOY SECRET :", err)
		return nil, err
	}

	return nil, nil
}

func (svr *DeployServer) GetSecretObject(pusher, namespace string, kvs []*pb.Secret) unstructured.Unstructured {
	obj := unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Secret",
	})

	obj.Object = map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Secret",
		"metadata": map[string]interface{}{
			"name":      "secret-" + pusher,
			"namespace": namespace,
		},
		"type": "Opaque",
		"data": map[string]string{},
	}

	for _, kv := range kvs {
		m := obj.Object["data"].(map[string]string)
		m[kv.Key] = base64.StdEncoding.EncodeToString([]byte(kv.Value))
	}

	return obj
}
