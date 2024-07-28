package main

import (
	"os"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
