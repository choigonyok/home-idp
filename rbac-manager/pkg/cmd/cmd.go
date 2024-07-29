package cmd

import (
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"

	"github.com/choigonyok/home-idp/pkg/server"
	rbacmanagerconfig "github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/policy"
	rbacserver "github.com/choigonyok/home-idp/rbac-manager/server"
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
			svr := server.New(config.RbacManager)
			rbacserver.IntegrateGrpcServerToServer(svr)
			defer svr.Server.CloseListner()
			defer svr.StorageClient.Close()

			log.Printf("Installing rbac-manager server is completed successfully!")
			log.Printf("Every installation has been finished successfully!\n")

			p, _ := policy.ParseToStruct([]byte(`
			{
				"policy": {
					"name": "example-policy",
					"effect": "Ask/Allow/Deny",
					"target": {
						"deploy": {
							"namespace": [
								"default",
								"test"
							],
							"resource": {
								"cpu": "500m",
								"memory": "1024Mi",
								"disk": "200Gi"
							},
							"gvk": [
								"apps/v1/Deployments",
								"networking.k8s.io/v1/Ingress",
								"/vi/Pod"
							]
						},
						"secret": {
							"path": [
								"/path1/to/secret/*",
								"/path2/to/secret/*"
							]
						}			
					},
					"action": [
						"Get",
						"Put",
						"Delete",
						"List"
					]
				}
			}
			`))
			fmt.Println("Name:", p.Name)
			fmt.Println("Effect:", p.Effect)
			fmt.Println("Actions:", p.Action)
			fmt.Println("Deploy.GVKs:", p.Target.Deploy.GVK)
			fmt.Println("Deploy.Namespace:", p.Target.Deploy.Namespace)
			fmt.Println("Deploy.CPU:", p.Target.Deploy.Resource.CPU)
			fmt.Println("Deploy.Memory:", p.Target.Deploy.Resource.Memory)
			fmt.Println("Deploy.Disk:", p.Target.Deploy.Resource.Disk)
			fmt.Println("Secret.path:", p.Target.Secret.Path)
			svr.Server.Serve()
			return nil
		},
	}

	serverStartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace where deploy-manager server is installed")
	serverStartCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Configuration file path for deploy-manager")

	return serverStartCmd
}
