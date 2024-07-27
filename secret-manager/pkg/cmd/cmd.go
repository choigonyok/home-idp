package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/grpc"
	pb "github.com/choigonyok/home-idp/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/server"
	secretmanagercfg "github.com/choigonyok/home-idp/secret-manager/pkg/config"
	"github.com/spf13/cobra"
)

const (
	defaultHomeIdpConfig = "$HOME/.home-idp/config.yaml"
)

type grpcServer struct {
	pb.UnimplementedGreeterServer
}

func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "secret-manager",
		Short: "Home-idp Secret-Manager",
		Args:  cobra.ExactArgs(0),
	}

	// c.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// addRootFlags(c)
	addSubCmds(c)

	return c
}

func addRootFlags(c *cobra.Command) {
	c.PersistentFlags().StringP("namespace", "n", "", "Kubernetes namespace")
}

func addSubCmds(c *cobra.Command) {
	c.AddCommand(cmd.GetServerCmd(config.SecretManager))
	c.AddCommand(getTestCmd())
	c.AddCommand(getTestClientCmd())

	getTestCmd().Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func getTestCmd() *cobra.Command {
	var filepath string

	testCmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			sm := secretmanagercfg.New()
			return sm.Init(filepath)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TEST COMMAND START")

			fmt.Println("SERVER CONFIGURATION START")
			s := server.New(config.SecretManager)
			defer s.StorageClient.Close()

			fmt.Println("REGISTER GRPC")
			pb.RegisterGreeterServer(s.Server, &grpcServer{})
			l := grpc.NewListener(env.Get("SECRET_MANAGER_PORT"))
			defer l.Close()

			fmt.Println("LISTENER START")
			s.Server.Serve(l)
			return nil
		},
	}

	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

	return testCmd
}

func getTestClientCmd() *cobra.Command {
	var filepath string

	testCmd := &cobra.Command{
		Use:   "test-client",
		Short: "test-client",
		// Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			sm := secretmanagercfg.New()
			return sm.Init(filepath)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TEST COMMAND START")

			conn := grpc.NewClientConn("localhost", "5103")
			defer conn.Close()
			c := pb.NewGreeterClient(conn)

			name := "world"
			if len(os.Args) > 1 {
				name = os.Args[1]
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
			defer cancel()
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			fmt.Println(err)
			fmt.Println(r.GetMessage())

			return nil
		},
	}

	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

	return testCmd
}
