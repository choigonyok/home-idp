package cmd

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/cmd"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/grpc"
	secretmanagercfg "github.com/choigonyok/home-idp/secret-manager/pkg/config"
	"github.com/spf13/cobra"
)

const (
	// Location to read istioctl defaults from
	defaultIstioctlConfig = "$HOME/.idpctl/config.yaml"
)

const (
	FlagCharts = "charts"
)

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

			fmt.Println("WATING GRPC CONNECTION")
			conn := grpc.NewListenerConn(env.Get("SECRET_MANAGER_PORT"))

			fmt.Println(conn.LocalAddr().String())
			fmt.Println(conn.RemoteAddr().String())

			// cfg := config.New(component)
			// cfg.Set()

			// svr := server.New(cfg)
			// svr.Run()

			return nil
		},
	}

	testCmd.PersistentFlags().StringVarP(&filepath, "config", "f", "", "Secret Manager Configuration File")

	return testCmd
}
