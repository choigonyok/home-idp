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
	// SetEnvVars registers environment variables from configuration file.
	SetEnvVars()
}

type GlobalConfig struct {
	Namespace     string               `yaml:"namespace,omitempty"`
	StorageClass  string               `yaml:"storageClass,omitempty"`
	Git           *GlobalConfigGit     `yaml:"git,omitempty"`
	Ingress       *GlobalConfigIngress `yaml:"ingress,omitempty"`
	AdminPassword string               `yaml:"adminPassword,omitempty"`
	Harbor        *GlobalConfigHarbor  `yaml:"harbor,omitempty"`
	UI            *GlobalConfigUI      `yaml:"ui,omitempty"`
}

type GlobalConfigUI struct {
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
	TLS  bool   `yaml:"tls,omitempty"`
}

type GlobalConfigHarbor struct {
	Host string `yaml:"host,omitempty"`
	TLS  bool   `yaml:"tls,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

type GlobalConfigIngress struct {
	Host string `yaml:"host,omitempty"`
	TLS  bool   `yaml:"tls,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

type GlobalConfigGit struct {
	Username string `yaml:"username,omitempty"`
	Token    string `yaml:"token,omitempty"`
	Repo     string `yaml:"repository,omitempty"`
	Email    string `yaml:"email,omitempty"`
}

const (
	DefaultConfigFilePath = "./.idpctl/config.yaml"
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
