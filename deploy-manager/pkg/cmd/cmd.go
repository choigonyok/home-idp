package cmd

import (
	"log"

	deploymanagerconfig "github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

// type grpcServer struct {
// 	pb.UnimplementedGreeterServer
// }

// func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

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
	serverCmd := cmd.GetServerCmd(config.DeployManager)
	c.AddCommand(serverCmd)
	serverCmd.AddCommand(getServerStartCmd())
}

// func getServerCmd() *cobra.Command {
// 	serverCmd := &cobra.Command{
// 		Use:   "server",
// 		Short: "server",
// 		Args:  cobra.ExactArgs(0),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			return nil
// 		},
// 	}

// 	serverCmd.AddCommand(getServerStartCmd())

// 	return serverCmd
// }

func getServerStartCmd() *cobra.Command {
	var filepath string
	var namespace string

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start deploy-manager server",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			dm := deploymanagerconfig.New()
			dm.Init(filepath)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Start installing deploy-manager server...")
			svr := server.New(config.DeployManager)
			defer svr.Listener.Close()
			defer svr.StorageClient.Close()
			log.Printf("Installing deploy-manager server is completed successfully!")

			log.Printf("Every installation has been finished successfully!\n")
			svr.Server.GrpcServer.Serve(svr.Listener)
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
