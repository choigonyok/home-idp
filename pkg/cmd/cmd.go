package cmd

import (
	"github.com/choigonyok/home-idp/pkg/config"

	"github.com/spf13/cobra"
)

func GetServerCmd(component config.Components) *cobra.Command {
	c := ""
	switch component {
	case 0:
		c = "secret-manager"
	case 1:
		c = "deploy-manager"
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
