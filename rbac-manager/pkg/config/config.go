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
	Storage  *config.StorageConfig     `yaml:"storage,omitempty"`
	Smtp     *mail.SmtpClient          `yaml:"smtp,omitempty"`
}

type RbacManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type Config struct {
	Service *RbacManagerConfig   `yaml:"rbac-manager,omitempty"`
	Global  *config.GlobalConfig `yaml:"global,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global: &config.GlobalConfig{},
		Service: &RbacManagerConfig{
			Service: &RbacManagerServiceConfig{},
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
	env.Set("HOME_IDP_STORAGE_USERNAME", cfg.Global.Storage.Username)
	env.Set("HOME_IDP_STORAGE_PASSWORD", cfg.Global.Storage.Password)
	env.Set("HOME_IDP_STORAGE_DATABASE", cfg.Global.Storage.Database)
	env.Set("RBAC_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
}
