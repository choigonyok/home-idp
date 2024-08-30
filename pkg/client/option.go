package client

import (
	"os"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/http"
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	helmcli "helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type GrpcClientOption struct {
	F func(ClientSet)
}

func (opt *GrpcClientOption) Apply(cli ClientSet) error {
	opt.F(cli)
	return nil
}

func WithGrpcRbacManagerClient(port int) ClientOption {
	return useGrpcRbacManagerClient(port)
}

func useGrpcRbacManagerClient(port int) ClientOption {
	return newGrpcRbacManagerClientOption(func(cli ClientSet) {
		grpcOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithTransportCredentials(tlsOpt),
		}

		conn, _ := grpc.NewClient("home-idp-rbac-manager:"+strconv.Itoa(port), grpcOptions...)
		cli.Set(util.GrpcRbacManagerClient, conn)
	})
}

func newGrpcRbacManagerClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithGrpcSecretManagerClient(port int) ClientOption {
	return useGrpcSecretManagerClient(port)
}

func useGrpcSecretManagerClient(port int) ClientOption {
	return newGrpcSecretManagerClientOption(func(cli ClientSet) {
		grpcOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithTransportCredentials(tlsOpt),
		}

		conn, _ := grpc.NewClient("home-idp-secret-manager:"+strconv.Itoa(port), grpcOptions...)
		cli.Set(util.GrpcSecretManagerClient, conn)
	})
}

func newGrpcSecretManagerClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithGrpcDeployManagerClient(port int) ClientOption {
	return useGrpcDeployManagerClient(port)
}

func useGrpcDeployManagerClient(port int) ClientOption {
	return newGrpcDeployManagerClientOption(func(cli ClientSet) {
		grpcOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithTransportCredentials(tlsOpt),
		}

		conn, _ := grpc.NewClient("home-idp-deploy-manager:"+strconv.Itoa(port), grpcOptions...)
		cli.Set(util.GrpcDeployManagerClient, conn)
	})
}

func newGrpcDeployManagerClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithGrpcInstallManagerClient(port int) ClientOption {
	return useGrpcInstallManagerClient(port)
}

func useGrpcInstallManagerClient(port int) ClientOption {
	return newGrpcInstallManagerClientOption(func(cli ClientSet) {
		grpcOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithTransportCredentials(tlsOpt),
		}

		conn, _ := grpc.NewClient("home-idp-install-manager:"+strconv.Itoa(port), grpcOptions...)
		cli.Set(util.GrpcInstallManagerClient, conn)
	})
}

func newGrpcInstallManagerClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithHelmClient() ClientOption {
	return useHelmClient()
}

func useHelmClient() ClientOption {
	return newHelmClientOption(func(cli ClientSet) {
		settings := helmcli.New()

		dl := &downloader.ChartDownloader{
			Out:              os.Stdout,
			RepositoryConfig: settings.RepositoryConfig,
			RepositoryCache:  settings.RepositoryCache,
			Getters:          getter.All(settings),
		}

		cli.Set(util.HelmClient, &helm.HelmClient{
			Downloader: dl,
			Setting:    settings,
			Repository: make(map[string]*repo.ChartRepository),
		})
	})
}

func newHelmClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithDockerClient() ClientOption {
	return useDockerClient()
}

func useDockerClient() ClientOption {
	return newDockerClientOption(func(cli ClientSet) {
		client, _ := dockercli.NewClientWithOpts(dockercli.FromEnv, dockercli.WithVersion("27.1.2"))

		cfg := registry.AuthConfig{
			Username: env.Get("DEPLOY_MANAGER_DOCKER_USERNAME"),
			Password: env.Get("DEPLOY_MANAGER_DOCKER_PASSWORD"),
		}

		config, _ := registry.EncodeAuthConfig(cfg)

		i := &docker.DockerClient{
			Client:         client,
			AuthCredential: &config,
		}

		cli.Set(util.DockerClient, i)
		return
	})
}

func newDockerClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithHttpClient() ClientOption {
	return useHttpClient()
}

func useHttpClient() ClientOption {
	// return newHttpClientOption(func(cli ClientSet) {
	// 	client := http.DefaultClient
	// 	i := &pkghttp.HttpClient{
	// 		Client: client,
	// 	}
	// 	cli.Set(util.HttpClient, i)
	// 	return
	// })
	return newHttpClientOption(func(cli ClientSet) {
		i := http.NewClient()
		cli.Set(util.HttpClient, i)
		return
	})
}

func newHttpClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithKubeClient() ClientOption {
	return useKubeClient()
}

func useKubeClient() ClientOption {
	return newKubeClientOption(func(cli ClientSet) {
		i := kube.NewKubeClient()
		cli.Set(util.KubeClient, i)
		return
	})
}

func newKubeClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}

func WithGitClient(owner, email, token string) ClientOption {
	return useGitClient(owner, email, token)
}

func useGitClient(owner, email, token string) ClientOption {
	return newGitClientOption(func(cli ClientSet) {
		i := git.NewGitClient(owner, email, token)
		cli.Set(util.GitClient, i)
		return
	})
}

func newGitClientOption(f func(cli ClientSet)) *GrpcClientOption {
	return &GrpcClientOption{
		F: f,
	}
}
