package main

import (
	"os"

	"github.com/choigonyok/home-idp/install-manager/pkg/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
