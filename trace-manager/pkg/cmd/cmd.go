package cmd

import (
	"strconv"

	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/trace-manager/pkg/service"

	"github.com/choigonyok/home-idp/trace-manager/pkg/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "trace-manager",
		Short: "HOME-IDP TRACE-MANAGER",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.TraceManager)
	c.AddCommand(serverCmd)
	serverCmd.AddCommand(getServerStartCmd())

	// c.AddCommand(getTestClientCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start trace-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New()
			cfg.SetEnvVars()
			port, _ := strconv.Atoi(env.Get("TRACE_MANAGER_SERVICE_PORT"))

			svc := service.New(
				port,
				client.WithStorageClient(
					"postgres",
					env.Get("HOME_IDP_STORAGE_USERNAME"),
					env.Get("HOME_IDP_STORAGE_PASSWORD"),
					env.Get("HOME_IDP_STORAGE_DATABASE"),
				),
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
