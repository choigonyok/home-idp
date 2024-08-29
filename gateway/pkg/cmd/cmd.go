package cmd

import (
	"strconv"

	"github.com/choigonyok/home-idp/gateway/pkg/config"
	"github.com/choigonyok/home-idp/gateway/pkg/service"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = ".idpctl/config.yaml"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "gateway",
		Short: "HOME-IDP GATEWAY",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.Gateway)
	c.AddCommand(serverCmd)

	serverCmd.AddCommand(getServerStartCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start gateway server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New()
			cfg.SetEnvVars()
			port, _ := strconv.Atoi(env.Get("GATEWAY_SERVICE_PORT"))
			svc := service.New(
				port,
				client.WithGrpcInstallManagerClient(5051),
				client.WithGrpcDeployManagerClient(5104),
				client.WithGitClient(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_GIT_TOKEN")),
				client.WithKubeClient(),
			)

			defer svc.Stop()

			svc.Start()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}

// 도커파일을 전송
// 도커파일을 깃허브에 저장
// 깃허브에서 도커파일을 가져와서 빌드 후 하버에 푸시
// 하버에서 이미지를 가져와서 이미지 이름과 태그로 분리
// 이미지 이름과 태그를 가지고 깃허브 레포의 manifest 들의 이미지와 버전을 수정해서 푸시
// ArgoCD가 감지하고 변경된 이미지로 배포

// argocd 배포 후 깃허브 레포지토리 연결
// application 파일 생성 후 푸시
// /cd 에 푸시된 깃허브가 웹훅을 보내고 이 웹훅을 받으면 application 업데이트된 내용으로 kubectl apply 해서 배포
