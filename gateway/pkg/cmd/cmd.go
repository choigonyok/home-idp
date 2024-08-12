package cmd

import (
	"fmt"

	"github.com/choigonyok/home-idp/gateway/pkg/config"
	"github.com/choigonyok/home-idp/gateway/pkg/server"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
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

			svc := server.New(
				5050,
				client.WithGrpcInstallManagerClient(5051),
				client.WithGrpcDeployManagerClient(5052),
			)
			defer svc.Stop()

			fmt.Println(svc.ClientSet.GrpcClient[util.InstallManager].GetConnection().Target())
			fmt.Println(svc.ClientSet.GrpcClient[util.InstallManager].GetConnection().GetState().String())
			fmt.Println()
			fmt.Println(svc.ClientSet.GrpcClient[util.DeployManager].GetConnection().Target())
			fmt.Println(svc.ClientSet.GrpcClient[util.DeployManager].GetConnection().GetState().String())

			svc.Start()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
