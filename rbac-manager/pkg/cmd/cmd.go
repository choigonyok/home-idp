package cmd

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	rbacmanagerconfig "github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

// func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

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
	serverCmd := cmd.GetServerCmd(config.RbacManager)
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			dm := rbacmanagerconfig.New()
			dm.Init(filepath)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Start installing rbac-manager server...")
			svr := server.New(config.DeployManager)
			defer svr.Listener.Close()
			defer svr.StorageClient.Close()

			log.Printf("Installing rbac-manager server is completed successfully!")

			log.Printf("Every installation has been finished successfully!\n")
			svr.Server.GrpcServer.Serve(svr.Listener)
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
