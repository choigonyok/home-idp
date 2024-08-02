package cmd

import (
	"log"

	deployconfig "github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/server"
	"github.com/choigonyok/home-idp/pkg/cmd"
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
			cfg := deployconfig.New()
			log.Printf("Start installing deploy-manager server...")
			svr := server.New(util.DeployManager, cfg)
			defer svr.Close()

			log.Printf("Installing deploy-manager server is completed successfully!")
			log.Printf("Every installation has been finished successfully!\n")
			svr.Run()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
