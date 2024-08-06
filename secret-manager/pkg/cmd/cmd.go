package cmd

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/util"
	secretconfig "github.com/choigonyok/home-idp/secret-manager/pkg/config"
	"github.com/choigonyok/home-idp/secret-manager/pkg/server"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "secret-manager",
		Short: "Home-idp Secret-Manager",
		Args:  cobra.ExactArgs(0),
	}
	addSubCmds(c)
	return c
}

func addSubCmds(c *cobra.Command) {
	serverCmd := cmd.GetServerCmd(util.SecretManager)
	c.AddCommand(serverCmd)

	serverCmd.AddCommand(getServerStartCmd())
}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start rbac-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := secretconfig.New()
			log.Printf("Start installing secret-manager server...")
			svr := server.New(util.RbacManager, cfg)
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

// func getTestCmd() *cobra.Command {
// 	var filepath string

// 	testCmd := &cobra.Command{
// 		Use:   "test",
// 		Short: "test",
// 		Args:  cobra.ExactArgs(0),
// 		PreRunE: func(cmd *cobra.Command, args []string) error {
// 			sm := secretmanagercfg.New()
// 			return sm.Init(filepath)
// 		},
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			log.Printf("Start installing secret-manager server...")
// 			svr := server.New(config.SecretManager)
// 			defer svr.StorageClient.Close()
// 			defer svr.Listener.Close()
// 			log.Printf("Installing secret-manager server is completed successfully!")

// 			log.Printf("Every installation has been finished successfully!\n")
// 			svr.Server.GrpcServer.Serve(svr.Listener)
// 			return nil
// 		},
// 	}

// 	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

// 	return testCmd
// }
