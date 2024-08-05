package grpc

import (
	"context"
	"time"

	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
)

func (g *GrpcClient) InstallArgoCD(opt *helm.ArgoCDOption, metadata *helm.ArgoCD) (*pb.InstallArgoCDChartReply, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()
	r, err := pb.NewArgoCDClient(g.InstallConn).InstallArgoCDChart(ctx, &pb.InstallArgoCDChartRequest{
		Opt: &pb.Option{
			RedisHa:            opt.RedisHA,
			ControllerRepl:     int32(opt.ControllerReplicas),
			ServerRepl:         int32(opt.ServerReplicas),
			RepoServerRepl:     int32(opt.RepoServerReplicas),
			ApplicationSetRepl: int32(opt.ApplicationSetReplicas),
			Domain:             opt.Domain,
			Ingress: &pb.Option_OptionIngress{
				Enabled:    opt.Ingress.Enabled,
				ClassName:  opt.Ingress.IngressClassName,
				Tls:        opt.Ingress.Tls,
				Annotation: *opt.Ingress.Annotation,
			},
			Argocd: &pb.Option_ArgoCD{
				Namespace:   metadata.Namespace,
				ReleaseName: metadata.ReleaseName,
			},
		},
	})

	return r, err
}
