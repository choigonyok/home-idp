package client

import (
	"fmt"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/manifest"
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/http"
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/util"
	"sigs.k8s.io/yaml"
)

type DeployManagerClientSet struct {
	GrpcClient   map[util.Components]client.GrpcClient
	MailClient   mail.MailClient
	KubeClient   *kube.KubeClient
	DockerClient *docker.DockerClient
	HttpClient   *http.HttpClient
}

func EmptyClientSet() *DeployManagerClientSet {
	return &DeployManagerClientSet{
		GrpcClient: make(map[util.Components]client.GrpcClient, client.ClientTotalCount),
	}
}

func (cs *DeployManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.GrpcRbacManagerClient:
		tmp := &grpc.InstallManagerGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.RbacManager] = tmp
		return
	case util.DockerClient:
		tmp := &docker.DockerClient{}
		tmp.Set(i)
		cs.DockerClient = tmp
		return
	case util.HttpClient:
		tmp := &http.HttpClient{}
		tmp.Set(i)
		cs.HttpClient = tmp
		return
	case util.KubeClient:
		tmp := &kube.KubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	default:
		return
	}
}

func (cs *DeployManagerClientSet) TestCreatHarborCredSecret() {
	s := manifest.GetHarborCredManifest(env.Get("DEPLOY_MANAGER_REGISTRY_PASSWORD"))
	b, _ := yaml.Marshal(s)

	if err := cs.KubeClient.ApplyManifest(string(b), "secrets", s.GetNamespace()); err != nil {
		fmt.Println("ERROR:", err)
	}

	// data := make(map[string]string, 1)
	// data[".dockerconfigjson"] = "eyJhdXRocyI6eyJoYXJib3IuaWRwLXN5c3RlbS5zdmMuY2x1c3Rlci5sb2NhbDo4MCI6eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJ0ZXN0ZXIxMjM0IiwiYXV0aCI6IllXUnRhVzQ2ZEdWemRHVnlNVEl6TkE9PSJ9fX0="

	// m := &manifest.KubeManifest{
	// 	GVK: s    "k8s.io/apimachinery/pkg/apis/meta/v1"
	//   "k8s.io/client-go/kubernetes/scheme"
	//   corev1 "k8s.io/api/core/v1"chema.GroupVersionKind{
	// 		Group:   "",
	// 		Version: "v1",
	// 		Kind:    "Secret",
	// 	},
	// 	Spec: &manifest.SecretSpec{
	// 		Name:      "harborcred-test",
	// 		Namespace: ns,
	// 		Type:      "kubernetes.io/dockerconfigjson",
	// 		Data:      &data,
	// 	},
	// }

	// if err := cs.KubeClient.ApplyManifest(m.GenerateManifest(), "secrets", ns); err != nil {
	// 	fmt.Println("ERROR:", err)
	// }
}

func (cs *DeployManagerClientSet) TestBuildWithKaniko(repo, imageName, imageTag string) {
	j := manifest.GetKanikoJobManifest(docker.NewDockerImage(imageName, imageTag), repo)

	b, _ := yaml.Marshal(j)

	if err := cs.KubeClient.ApplyManifest(string(b), "jobs", j.GetNamespace()); err != nil {
		fmt.Println("ERROR:", err)
	}
}

// func (c *DockerClient) Test(namespace string) {
// 	cfg := registry.AuthConfig{
// 		// ServerAddress: "harbor." + namespace + ".svc.cluster.local:80",
// 		ServerAddress: "harbor",
// 	}
// 	err := c.LoginWithEnv(cfg)
// 	fmt.Println(err)

// 	err = c.Client.ImageTag(context.TODO(), "nginx:latest", cfg.ServerAddress+"/nginx:latest")

// 	_, err = c.Client.ImagePull(context.TODO(), "nginx:latest", image.PullOptions{})

// 	tmp, _ := json.Marshal(cfg)
// 	_, err = c.Client.ImagePush(context.TODO(), cfg.ServerAddress+"/nginx:latest", image.PushOptions{
// 		RegistryAuth: base64.URLEncoding.EncodeToString(tmp),
// 	})
// }
