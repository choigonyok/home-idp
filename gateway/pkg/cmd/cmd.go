package cmd

import (
	"log"

	gatewayconfig "github.com/choigonyok/home-idp/gateway/pkg/config"
	"github.com/choigonyok/home-idp/gateway/pkg/server"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
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
			cfg := gatewayconfig.New()
			log.Printf("Start installing gateway...")
			svr := server.New(cfg)
			defer svr.Close()

			log.Printf("Installing rbac-manager server is completed successfully!")
			log.Printf("Every installation has been finished successfully!\n")

			svr.Run()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
