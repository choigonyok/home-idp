package cmd

import (
	"log"

	installconfig "github.com/choigonyok/home-idp/deploy-manager/pkg/config"
	"github.com/choigonyok/home-idp/install-manager/pkg/server"
	"github.com/choigonyok/home-idp/pkg/cmd"
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

			svr.Run()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where install-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for install-manager")

	return serverStartCmd
}

// func getTestClientCmd() *cobra.Command {
// 	var filepath string

// 	testCmd := &cobra.Command{
// 		Use:   "test-client",
// 		Short: "test-client",
// 		// Args:  cobra.ExactArgs(1),
// 		// PreRunE: func(cmd *cobra.Command, args []string) error {
// 		// 	dm := rbacmanagerconfig.New()
// 		// 	dm.Init(filepath)
// 		// 	return nil
// 		// },
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			log.Printf("Start installing rbac-manager server...")
// 			// svr := server.New(util.RbacManager)
// 			// defer svr.Close()

// 			// log.Printf("Installing rbac-manager server is completed successfully!")
// 			// log.Printf("Every installation has been finished successfully!\n")

// 			// svr.Run()
// 			conn1 := grpc.NewClient("localhost", "5105")
// 			defer conn1.Close()
// 			c1 := pb.NewUserServiceClient(conn1)
// 			ctx1, cancel := context.WithTimeout(context.Background(), time.Second*1)
// 			defer cancel()

// 			// r1, err := c1.GetUserInfo(ctx1, &pb.GetUserInfoRequest{
// 			// 	Id: int32(1),
// 			// })
// 			r2, err := c1.PutUser(ctx1, &pb.PutUserRequest{
// 				Email:     "tester1234@naver.com",
// 				Name:      "TESTER-choi",
// 				Password:  "HEWWL",
// 				ProjectId: 1,
// 			})

// 			fmt.Println("ERROR:", err)
// 			fmt.Println("SUCCESS: ", r2.GetSucceed())
// 			return nil
// 		},
// 	}

// 	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

// 	return testCmd
// }
