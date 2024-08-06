package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/grpc"
	"github.com/choigonyok/home-idp/pkg/util"
	rbacconfig "github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/server"
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

	c.AddCommand(getTestClientCmd())

}

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start rbac-manager server",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := rbacconfig.New()
			log.Printf("Start installing rbac-manager server...")
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

func getTestClientCmd() *cobra.Command {
	var filepath string

	testCmd := &cobra.Command{
		Use:   "test-client",
		Short: "test-client",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Start installing install-manager server...")
			conn1 := grpc.NewClient("localhost", "5107")
			defer conn1.Close()
			c1 := pb.NewArgoCDClient(conn1)
			ctx1, cancel := context.WithTimeout(context.Background(), time.Second*1)
			defer cancel()

			r2, err := c1.InstallArgoCDChart(ctx1, &pb.InstallArgoCDChartRequest{
				Opt: &pb.Option{
					RedisHa:            true,
					ControllerRepl:     2,
					ServerRepl:         2,
					RepoServerRepl:     2,
					ApplicationSetRepl: 2,
					Domain:             "test.slexn.com",
					Ingress:            &pb.Option_OptionIngress{},
					Argocd:             &pb.Option_ArgoCD{},
				},
			})

			fmt.Println("ERROR:", err)
			fmt.Println("SUCCESS: ", r2.GetSucceed())
			return nil
		},
	}

	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

	return testCmd
}
