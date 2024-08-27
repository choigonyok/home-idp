package cmd

import (
	"fmt"
	"strconv"
	"time"

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
				client.WithHttpClient(),
				client.WithGitClient(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_GIT_TOKEN")),
			)

			defer svc.Stop()

			fmt.Println("TEST START UPDATE")
			svc.ClientSet.GitClient.UpdateArgoCDApplicationManifest("testuser", "tester@naver.com", "test:v1.3", "test:v1.4")
			time.Sleep(time.Second * 5)
			svc.ClientSet.GitClient.UpdateArgoCDApplicationManifest("testuser", "tester@naver.com", "test:v1.4", "test:v1.5")
			fmt.Println("TEST END UPDATE")

			svc.Start()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
