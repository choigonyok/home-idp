package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/git"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
)

type DeployServer struct {
	pb.DeployServer
	KubeClient *kube.DeployManagerKubeClient
	GitClient  *git.DeployManagerGitClient
}

func (svr *DeployServer) Deploy(ctx context.Context, in *pb.DeployRequest) (*pb.DeployReply, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()

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
