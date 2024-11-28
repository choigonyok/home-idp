package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/mail"
)

type RbacManagerConfig struct {
	Name     string
	Enabled  bool                      `yaml:"enabled,omitempty"`
	Service  *RbacManagerServiceConfig `yaml:"service,omitempty"`
	Replicas int                       `yaml:"replicas,omitempty"`
	Smtp     *mail.SmtpClient          `yaml:"smtp,omitempty"`
}

type RbacManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type Config struct {
	Service *RbacManagerConfig   `yaml:"rbac-manager,omitempty"`
	Global  *config.GlobalConfig `yaml:"global,omitempty"`
	Storage *config.Storage      `yaml:"storage,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global: &config.GlobalConfig{
			Ingress: &config.GlobalConfigIngress{},
			Git: &config.GlobalConfigGit{
				Oauth: &config.GlobalConfigGitOauth{},
			},
			Harbor: &config.GlobalConfigHarbor{},
			UI:     &config.GlobalConfigUI{},
		},
		Service: &RbacManagerConfig{
			Service: &RbacManagerServiceConfig{},
		},
		Storage: &config.Storage{
			Persistence: &config.Persistence{},
		},
	}

	config.ParseFromFile(cfg, config.DefaultConfigFilePath)
	return cfg
}

func (cfg *Config) SetEnvVars() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("HOME_IDP_ADMIN_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_ADMIN_PASSWORD", cfg.Global.AdminPassword)
	env.Set("HOME_IDP_GIT_EMAIL", cfg.Global.Git.Email)
	env.Set("HOME_IDP_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_STORAGE_USERNAME", cfg.Storage.Username)
	env.Set("HOME_IDP_STORAGE_PASSWORD", cfg.Storage.Password)
	env.Set("HOME_IDP_STORAGE_DATABASE", cfg.Storage.Database)
	env.Set("RBAC_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
	env.Set("HOME_IDP_GIT_TOKEN", cfg.Global.Git.Token)
}
