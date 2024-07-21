package main

import (
	"os"

	"github.com/choigonyok/home-idp/idpctl/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	// if err := cmd.ConfigAndEnvProcessing(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Could not initialize: %v\n", err)
	// 	exitCode := cmd.GetExitCode(err)
	// 	os.Exit(exitCode)
	// }

	rootCmd := cmd.GetRootCmd(os.Args[1:])
	// log.EnableKlogWithCobra()

	if err := rootCmd.Execute(); err != nil {
		// exitCode := cmd.GetExitCode(err)
		// os.Exit(exitCode)
	}
}
