package cmd

import (
	"log"
	"strconv"
	"time"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "deploy-manager",
		Short: "Home-idp Deploy-Manager",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.DeployManager)
	c.AddCommand(serverCmd)
	serverCmd.AddCommand(getServerStartCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start deploy-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New()
			cfg.SetEnvVars()
			port, _ := strconv.Atoi(env.Get("DEPLOY_MANAGER_SERVICE_PORT"))

			log.Printf("Start installing deploy-manager server...")
			svc := service.New(
				port,
				client.WithDockerClient(),
			)
			defer svc.Stop()

			time.Sleep(time.Second * 60)
			svc.ClientSet.DockerClient.Test(env.Get("GLOBAL_NAMESPACE"))

			// fmt.Println(*svc.ClientSet.DockerClient.AuthCredential)
			// fmt.Println(*svc.ClientSet.DockerClient.AuthCredential)
			// err := svc.ClientSet.DockerClient.Build("testtag", `
			// FROM nginx:latest

			// # Nginx가 설치된 Debian 기반 이미지에 필요한 패키지 설치
			// RUN apt-get update \
			// 		&& apt-get install -y git \
			// 		&& rm -rf /var/lib/apt/lists/*

			// # Nginx 설정 파일 복사
			// COPY nginx.conf /etc/nginx/nginx.conf

			// # 포트 80 오픈
			// EXPOSE 80

			// # Nginx 실행
			// CMD ["nginx", "-g", "daemon off;"]
			// `)

			// fmt.Println(err)

			svc.Start()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
