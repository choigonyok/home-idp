package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
)

type Config struct {
	Service *InstallManagerConfig `yaml:"install-manager,omitempty"`
	Storage *config.Storage       `yaml:"storage,omitempty"`
	Global  *config.GlobalConfig  `yaml:"global,omitempty"`
}

type InstallManagerConfig struct {
	Enabled  bool                         `yaml:"enabled,omitempty"`
	Service  *InstallManagerServiceConfig `yaml:"service,omitempty"`
	Replicas int                          `yaml:"replicas,omitempty"`
}

type InstallManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type StorageClassPersistentConfig struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Size    string `yaml:"size,omitempty"`
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
		Service: &InstallManagerConfig{
			Service: &InstallManagerServiceConfig{},
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
	env.Set("HOME_IDP_ADMIN_PASSWORD", cfg.Global.AdminPassword)
	env.Set("HOME_IDP_STORAGE_CLASS_NAME", cfg.Storage.Persistence.StorageClass)
	env.Set("HOME_IDP_API_HOST", cfg.Global.Ingress.Host)
	env.Set("HOME_IDP_API_PORT", strconv.Itoa(cfg.Global.Ingress.Port))
	env.Set("HOME_IDP_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_GIT_TOKEN", cfg.Global.Git.Token)
	env.Set("HOME_IDP_GIT_EMAIL", cfg.Global.Git.Email)
	env.Set("HOME_IDP_GIT_REPO", cfg.Global.Git.Repo)
	env.Set("HOME_IDP_NAMESPACE", cfg.Global.Namespace)
	env.Set("HOME_IDP_HARBOR_HOST", cfg.Global.Harbor.Host)
	env.Set("HOME_IDP_HARBOR_PORT", strconv.Itoa(cfg.Global.Harbor.Port))
	env.Set("HOME_IDP_HARBOR_TLS_ENABLED", strconv.FormatBool(cfg.Global.Harbor.TLS))
	env.Set("HOME_IDP_STORAGE_USERNAME", cfg.Storage.Username)
	env.Set("HOME_IDP_STORAGE_PASSWORD", cfg.Storage.Password)
	env.Set("HOME_IDP_STORAGE_DB", cfg.Storage.Database)
	env.Set("INSTALL_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
	env.Set("POSTGRES_STORAGECLASS", cfg.Storage.Persistence.StorageClass)
	env.Set("POSTGRES_SIZE", cfg.Storage.Persistence.Size)
}
