package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/grpc"
	"github.com/choigonyok/home-idp/pkg/util"
	rbacmanagerconfig "github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
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
	c.AddCommand(getTestClientCmd())

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
			svr := server.New(util.RbacManager)
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
		// Args:  cobra.ExactArgs(1),
		// PreRunE: func(cmd *cobra.Command, args []string) error {
		// 	dm := rbacmanagerconfig.New()
		// 	dm.Init(filepath)
		// 	return nil
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Start installing rbac-manager server...")
			// svr := server.New(util.RbacManager)
			// defer svr.Close()

			// log.Printf("Installing rbac-manager server is completed successfully!")
			// log.Printf("Every installation has been finished successfully!\n")

			// svr.Run()
			fmt.Println("T1")
			conn1 := grpc.NewClientConn("localhost", "5105")
			fmt.Println("T2")
			defer conn1.Close()

			c1 := pb.NewUserServiceClient(conn1)
			fmt.Println("T3")
			ctx1, cancel := context.WithTimeout(context.Background(), time.Second*1)
			fmt.Println("T4")
			defer cancel()
			fmt.Println("T5")

			r1, err := c1.GetUserInfo(ctx1, &pb.GetUserInfoRequest{
				Id: int32(1),
			})
			fmt.Println("T6")
			fmt.Println(err)
			fmt.Println(r1.GetPassword())
			fmt.Println(r1.GetName())
			fmt.Println(r1.GetEmail())
			return nil
		},
	}

	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

	return testCmd
}
