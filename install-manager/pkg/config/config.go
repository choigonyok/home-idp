package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/install-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/pkg/mail"
)

type InstallManagerConfig struct {
	Name     string
	Enabled  bool                             `yaml:"enabled,omitempty"`
	Service  *service.InstallManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                              `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig            `yaml:"storage,omitempty"`
	Smtp     *mail.SmtpClient                 `yaml:"smtp,omitempty"`
}

func New() *InstallManagerConfig {
	cfg := &InstallManagerConfig{Name: "install-manager"}

	log.Printf("Start reading home-idp configuration file...")
	file.ParseConfigFile(cfg, config.DefaultConfigFilePath)
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *InstallManagerConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("RBAC_MANAGER_PORT", strconv.Itoa(cfg.Service.Port))
}

func (cfg *InstallManagerConfig) GetName() string {
	return cfg.Name
}
