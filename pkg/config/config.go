package config

import (
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type StorageConfig struct {
	Type     string `yaml:"type,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type Config interface {
	SetEnvFromConfig()
	GetName() string
}

const (
	DefaultConfigFilePath = "./.idtctl/config.yaml"
)

func Enabled(component util.Components, client string) bool {
	prefix := env.GetEnvPrefix(component)
	switch client {
	case "mail":
		if env.Get(prefix+"_MANAGER_SMTP_ENABLED") == "true" {
			return true
		}
	case "storage":
		if env.Get(prefix+"_MANAGER_SMTP_ENABLED") == "true" {
			return true
		}
	}
	return false
}
