package grpc

import (
	"context"

	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
)

type ArgoCDServer struct {
	pb.UnimplementedArgoCDServer
}

func (s *InstallManagerServer) InstallChart(ctx context.Context, in *pb.InstallArgoCDChartRequest) (*pb.InstallArgoCDChartReply, error) {
	tls := in.GetOpt().GetIngress().GetAnnotation()

	opt := &helm.ArgoCDOption{
		RedisHA:                in.GetOpt().GetRedisHa(),
		ControllerReplicas:     int(in.GetOpt().GetControllerRepl()),
		ServerReplicas:         int(in.GetOpt().GetServerRepl()),
		RepoServerReplicas:     int(in.GetOpt().GetRepoServerRepl()),
		ApplicationSetReplicas: int(in.GetOpt().GetApplicationSetRepl()),
		Domain:                 in.Opt.Domain,
		Ingress: &helm.ArgoCDIngressOption{
			Enabled:          in.GetOpt().GetIngress().GetEnabled(),
			IngressClassName: in.GetOpt().GetIngress().GetClassName(),
			Annotation:       &tls,
			Tls:              in.GetOpt().GetIngress().GetTls(),
		},
	}

	client := helm.NewArgoCDClient(in.GetOpt().GetArgocd().GetNamespace(), in.GetOpt().GetArgocd().GetReleaseName())
	client.Install(s.HelmClient, opt)

	return &pb.InstallArgoCDChartReply{
		Succeed: true,
	}, nil
}

func (s *InstallManagerServer) UninstallChart(ctx context.Context, in *pb.UninstallArgoCDChartRequest) (*pb.UninstallArgoCDChartReply, error) {
	return &pb.UninstallArgoCDChartReply{
		Succeed: true,
	}, nil
}
