package config

import (
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/pkg/util"
	"gopkg.in/yaml.v2"
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

type Storage struct {
	Persistence *Persistence `yaml:"persistence"`
	Username    string       `yaml:"username"`
	Password    string       `yaml:"password"`
	Database    string       `yaml:"database"`
}

type Persistence struct {
	StorageClass string `yaml:"storageClass"`
	Size         string `yaml:"size"`
}

type GlobalConfig struct {
	Namespace     string               `yaml:"namespace,omitempty"`
	Port          int                  `yaml:"port,omitempty"`
	Ingress       *GlobalConfigIngress `yaml:"ingress,omitempty"`
	Git           *GlobalConfigGit     `yaml:"git,omitempty"`
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
	Username string                `yaml:"username,omitempty"`
	Token    string                `yaml:"token,omitempty"`
	Repo     string                `yaml:"repository,omitempty"`
	Email    string                `yaml:"email,omitempty"`
	Oauth    *GlobalConfigGitOauth `yaml:"oauth,omitempty"`
}

type GlobalConfigGitOauth struct {
	ClientID     string `yaml:"clientId,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}

const (
	DefaultConfigFilePath = "/etc/home-idp/config.yaml"
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

func ParseFromFile(cfg Config, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return err
	}

	return nil
}
