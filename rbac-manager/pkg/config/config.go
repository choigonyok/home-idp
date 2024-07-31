package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/mail"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/service"
)

type RbacManagerConfig struct {
	Name     string
	Enabled  bool                          `yaml:"enabled,omitempty"`
	Service  *service.RbacManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                           `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig         `yaml:"storage,omitempty"`
	Smtp     *mail.SmtpClient              `yaml:"smtp,omitempty"`
}

func New() *RbacManagerConfig {
	cfg := &RbacManagerConfig{Name: "rbac-manager"}

	log.Printf("Start reading home-idp configuration file...")
	file.ParseConfigFile(cfg, config.DefaultConfigFilePath)
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *RbacManagerConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("RBAC_MANAGER_PORT", strconv.Itoa(cfg.Service.Port))
	env.Set("RBAC_MANAGER_STORAGE_TYPE", cfg.Storage.Type)
	env.Set("RBAC_MANAGER_STORAGE_HOST", cfg.Storage.Host)
	env.Set("RBAC_MANAGER_STORAGE_USERNAME", cfg.Storage.Username)
	env.Set("RBAC_MANAGER_STORAGE_PASSWORD", cfg.Storage.Password)
	env.Set("RBAC_MANAGER_STORAGE_DATABASE", cfg.Storage.Database)
	env.Set("RBAC_MANAGER_STORAGE_PORT", strconv.Itoa(cfg.Storage.Port))
	if cfg.Smtp.Enabled == true {
		env.Set("RBAC_MANAGER_SMTP_HOST", cfg.Smtp.Config.Host)
		env.Set("RBAC_MANAGER_SMTP_PORT", cfg.Smtp.Config.Port)
		env.Set("RBAC_MANAGER_SMTP_USER", cfg.Smtp.Config.User)
		env.Set("RBAC_MANAGER_SMTP_PASSWORD", cfg.Smtp.Config.Password)
		env.Set("RBAC_MANAGER_SMTP_DOMAIN", cfg.Smtp.Config.Domain)
		env.Set("RBAC_MANAGER_SMTP_ENABLED", strconv.FormatBool(cfg.Smtp.Enabled))
	}
}

func (cfg *RbacManagerConfig) GetName() string {
	return cfg.Name
}
