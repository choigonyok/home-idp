package cmd

import (
	"flag"

	"github.com/spf13/cobra"
)

// AddFlags adds all command line flags to the given command.
func AddFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
