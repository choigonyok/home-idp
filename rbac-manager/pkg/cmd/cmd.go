package cmd

import (
	"strconv"

	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/service"

	"github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "rbac-manager",
		Short: "HOME-IDP RBAC-MANAGER",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.RbacManager)
	c.AddCommand(serverCmd)
	serverCmd.AddCommand(getServerStartCmd())

	// c.AddCommand(getTestClientCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start rbac-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New()
			cfg.SetEnvVars()
			port, _ := strconv.Atoi(env.Get("RBAC_MANAGER_SERVICE_PORT"))

			svc := service.New(
				port,
				client.WithStorageClient("postgres"),
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
