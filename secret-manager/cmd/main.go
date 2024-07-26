package main

import (
	"os"

	"github.com/choigonyok/home-idp/idpctl/pkg/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
