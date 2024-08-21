package client

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/http"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/util"
)

type DeployManagerClientSet struct {
	GrpcClient   map[util.Components]client.GrpcClient
	MailClient   mail.MailClient
	KubeClient   *kube.DeployManagerKubeClient
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
		tmp := &kube.DeployManagerKubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	default:
		return
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
