package grpc

import (
	"context"

	installhelm "github.com/choigonyok/home-idp/install-manager/pkg/helm"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type ArgoCDServer struct {
	pb.UnimplementedArgoCDServer
	HelmClient *helm.HelmClient
}

func (s *ArgoCDServer) InstallArgoCDChart(ctx context.Context, in *pb.InstallArgoCDChartRequest) (*pb.InstallArgoCDChartReply, error) {
	tls := in.GetOpt().GetIngress().GetAnnotation()

	opt := &installhelm.ArgoCDOption{
		RedisHA:                in.GetOpt().GetRedisHa(),
		ControllerReplicas:     int(in.GetOpt().GetControllerRepl()),
		ServerReplicas:         int(in.GetOpt().GetServerRepl()),
		RepoServerReplicas:     int(in.GetOpt().GetRepoServerRepl()),
		ApplicationSetReplicas: int(in.GetOpt().GetApplicationSetRepl()),
		Domain:                 in.Opt.Domain,
		Ingress: &installhelm.ArgoCDIngressOption{
			Enabled:          in.GetOpt().GetIngress().GetEnabled(),
			IngressClassName: in.GetOpt().GetIngress().GetClassName(),
			Annotation:       &tls,
			Tls:              in.GetOpt().GetIngress().GetTls(),
		},
	}

	client := installhelm.NewArgoCDClient(in.GetOpt().GetArgocd().GetNamespace(), in.GetOpt().GetArgocd().GetReleaseName())

	client.Install(*s.HelmClient, opt)

	return &pb.InstallArgoCDChartReply{
		Succeed: true,
	}, nil
}

func (s *ArgoCDServer) UninstallArgoCDChart(ctx context.Context, in *pb.UninstallArgoCDChartRequest) (*pb.UninstallArgoCDChartReply, error) {
	return &pb.UninstallArgoCDChartReply{
		Succeed: true,
	}, nil
}
