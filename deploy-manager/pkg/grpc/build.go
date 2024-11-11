package grpc

import (
	"context"
	"fmt"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/trace"
)

type BuildServer struct {
	pb.BuildServer
	KubeClient  *kube.DeployManagerKubeClient
	TraceClient *trace.TraceClient
}

func (svr *BuildServer) BuildDockerfile(ctx context.Context, in *pb.BuildDockerfileRequest) (*pb.BuildDockerfileReply, error) {
	deployKanikoSpan := svr.TraceClient.NewSpanFromIncomingContext(ctx)
	fmt.Println("[deployKanikoSpan ID]", deployKanikoSpan.SpanID)
	err := deployKanikoSpan.Start(ctx)
	if err != nil {
		fmt.Println("DEPLOY KANIKO SPAN START ERR:", err)
	}

	if err := svr.KubeClient.ApplyKanikoBuildJob(in.Img.GetName()+":"+in.Img.GetVersion(), in.Img.GetPusher(), in.Img.GetRepository(), in.GetProject()); err != nil {
		fmt.Println("TEST APPLY KANIKO BUILD MANIFEST ERR:", err)
		return &pb.BuildDockerfileReply{
			Succeed: false,
		}, nil
	}

	err = deployKanikoSpan.Stop()
	if err != nil {
		fmt.Println("DEPLOY KANIKO SPAN STOP ERR:", err)
	}

	return &pb.BuildDockerfileReply{
		Succeed: true,
	}, nil
}
