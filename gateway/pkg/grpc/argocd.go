package grpc

import (
	"context"
	"time"

	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	"google.golang.org/grpc"
)

func (g *GrpcClient) InstallArgoCD(data *helm.ArgoCDData) (*pb.InstallArgoCDChartReply, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()
	client := pb.NewArgoCDClient(g.InstallConn)

	r, err := client.InstallArgoCDChart(ctx, &pb.InstallArgoCDChartRequest{
		Opt: &pb.Option{
			RedisHa:            data.OptRedisHA,
			ControllerRepl:     int32(data.OptControllerReplicas),
			ServerRepl:         int32(data.OptServerReplicas),
			RepoServerRepl:     int32(data.OptRepoServerReplicas),
			ApplicationSetRepl: int32(data.OptApplicationSetReplicas),
			Domain:             data.OptDomain,
			Ingress: &pb.Option_OptionIngress{
				Enabled:    data.IngressEnabled,
				ClassName:  data.IngressClassName,
				Tls:        data.IngressTls,
				Annotation: data.IngressAnnotation,
			},
			Argocd: &pb.Option_ArgoCD{
				Namespace:   data.MetadataNamespace,
				ReleaseName: data.MetadataReleaseName,
			},
		},
	},
		grpc.FailFastCallOption{
			FailFast: false,
		},
	)
	return r, err
}

// curl -X POST localhost:5106/charts/argocd \
// -d '{"namespace": "test-ns", "release_name": "testargocd-one", "controller_repl": 3, "server_repl": 3, "repo_server_repl": 3, "application_set_repl": 5, "domain": "test.choigonyok.com", "class_name": "test", "annotation": {"key": "values"}}'

func (g *GrpcClient) UpgradeArgoCD(data *helm.ArgoCDData) (*pb.InstallArgoCDChartReply, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()
	client := pb.NewArgoCDClient(g.InstallConn)

	r, err := client.InstallArgoCDChart(ctx, &pb.InstallArgoCDChartRequest{
		Opt: &pb.Option{
			RedisHa:            data.OptRedisHA,
			ControllerRepl:     int32(data.OptControllerReplicas),
			ServerRepl:         int32(data.OptServerReplicas),
			RepoServerRepl:     int32(data.OptRepoServerReplicas),
			ApplicationSetRepl: int32(data.OptApplicationSetReplicas),
			Domain:             data.OptDomain,
			Ingress: &pb.Option_OptionIngress{
				Enabled:    data.IngressEnabled,
				ClassName:  data.IngressClassName,
				Tls:        data.IngressTls,
				Annotation: data.IngressAnnotation,
			},
			Argocd: &pb.Option_ArgoCD{
				Namespace:   data.MetadataNamespace,
				ReleaseName: data.MetadataReleaseName,
			},
		},
	},
		grpc.FailFastCallOption{
			FailFast: false,
		},
	)
	return r, err
}
