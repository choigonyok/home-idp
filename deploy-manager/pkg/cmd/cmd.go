package cmd

import (
	"fmt"
	"log"

	deploymanagerconfig "github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/grpc"
	pb "github.com/choigonyok/home-idp/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

type grpcServer struct {
	pb.UnimplementedGreeterServer
}

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
			fmt.Println("ENABELD:", dm.Config.Enabled)
			fmt.Println("REPLICAS:", dm.Config.Replicas)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Start installing deploy-manager server...")
			s := server.New(config.DeployManager)
			log.Printf("Installing deploy-manager server is completed successfully!")

			log.Printf("Start attach grpc server to deploy-manager server...")
			pb.RegisterGreeterServer(s.Server, &grpcServer{})
			l := grpc.NewListener("5104")
			defer l.Close()

			log.Printf("Every installation has finished successfully!\n")
			s.Server.Serve(l)
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
