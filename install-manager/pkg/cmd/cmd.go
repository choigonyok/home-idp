package cmd

import (
	"log"

	installconfig "github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/install-manager/pkg/server"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/util"

	// pb "github.com/choigonyok/home-idp/instasll-manager/pkg/proto"

	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "rbac-manager",
		Short: "HOME-IDP INSTALL-MANAGER",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.RbacManager)
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
			cfg := installconfig.New()
			log.Printf("Start installing install-manager server...")
			svr := server.New(util.InstallManager, cfg)
			defer svr.Close()

			log.Printf("Installing install-manager server is completed successfully!")
			log.Printf("Every installation has been finished successfully!\n")

			h := helm.New()
			h.AddRepository("bitnami", "https://charts.bitnami.com/bitnami", true)
			h.Install("bitnami/nginx:17.3.0", "default", "nginx-tester")
			h.Uninstall("nginx-tester", "default")

			svr.Run()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where install-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for install-manager")

	return serverStartCmd
}
