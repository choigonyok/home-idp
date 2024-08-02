package cmd

import (
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/spf13/cobra"
)

func GetServerCmd(component util.Components) *cobra.Command {
	c := ""
	switch component {
	case 0:
		c = "secret-manager"
	case 1:
		c = "deploy-manager"
	case 2:
		c = "rbac-manager"
	case 3:
		c = "gateway"
	}

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "home-idp " + c + " server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return serverCmd
}
