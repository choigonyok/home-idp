package cmd

import (
	"github.com/choigonyok/home-idp/pkg/cmd"
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
	c.AddCommand(cmd.GetServerCmd(cmd.SecretManager))
}
