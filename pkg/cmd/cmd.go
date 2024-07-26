package cmd

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/spf13/cobra"
)

func GetServerCmd(component config.Components) *cobra.Command {
	c := ""
	switch component {
	case 0:
		c = "secret-manager"
	}

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start home-idp " + c + " server",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("INSTALL " + c + " START")

			cfg := config.New(component)
			cfg.Set()

			svr := server.New(cfg)
			svr.Run()

			return nil
		},
	}

	return serverCmd
}
