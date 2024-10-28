package cmd

import (
	"strconv"

	"github.com/choigonyok/home-idp/install-manager/pkg/config"
	"github.com/choigonyok/home-idp/install-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "install-manager",
		Short: "HOME-IDP INSTALL-MANAGER",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.InstallManager)
	c.AddCommand(serverCmd)
	serverCmd.AddCommand(getServerStartCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start install-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New()
			cfg.SetEnvVars()
			port, _ := strconv.Atoi(env.Get("INSTALL_MANAGER_SERVICE_PORT"))

			svc := service.New(
				port,
				client.WithHelmClient(),
				client.WithKubeClient(),
				client.WithHttpClient(),
				client.WithGitClient(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_GIT_TOKEN")),
			)
			defer svc.Stop()

			svc.Start()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where install-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for install-manager")

	return serverStartCmd
}
