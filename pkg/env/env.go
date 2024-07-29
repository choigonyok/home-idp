package env

import (
	"os"

	"github.com/choigonyok/home-idp/pkg/util"
)

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func GetEnvPrefix(c util.Components) string {
	switch c {
	case 0: // Secret-Manager
		return "SECRET"
	case 1: // Deploy-Manager
		return "DEPLOY"
	case 2: // Rbac-Manager
		return "RBAC"
	}
	return ""
}
