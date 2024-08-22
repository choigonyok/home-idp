package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
)

type BuildServer struct {
	pb.BuildServer
	KubeClient *kube.DeployManagerKubeClient
}

func (svr *BuildServer) BuildDockerfile(ctx context.Context, in *pb.BuildDockerfileRequest) (*pb.BuildDockerfileReply, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()

	if err := svr.KubeClient.ApplyKanikoBuildJob(in.Img.GetName() + ":" + in.Img.GetVersion()); err != nil {
		fmt.Println("TEST APPLY KANIKO BUILD MANIFEST ERR:", err)
		return &pb.BuildDockerfileReply{
			Succeed: false,
		}, nil
	}
	return &pb.BuildDockerfileReply{
		Succeed: true,
	}, nil
}
