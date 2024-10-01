package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service *InstallManagerConfig `yaml:"install-manager,omitempty"`
	Global  *config.GlobalConfig  `yaml:"global,omitempty"`
}

type InstallManagerConfig struct {
	Enabled         bool                                 `yaml:"enabled,omitempty"`
	Service         *InstallManagerServiceConfig         `yaml:"service,omitempty"`
	Replicas        int                                  `yaml:"replicas,omitempty"`
	DefaultRegistry *InstallManagerConfigDefaultRegistry `yaml:"defaultRegistry,omitempty"`
	DefaultCD       *InstallManagerConfigDefaultCD       `yaml:"defaultCD,omitempty"`
	DefaultCI       *InstallManagerConfigDefaultCI       `yaml:"defaultCI,omitempty"`
}

type InstallManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type InstallManagerConfigDefaultCI struct {
	Enabled       bool                          `yaml:"enabled,omitempty"`
	Type          string                        `yaml:"type,omitempty"`
	AdminPassword string                        `yaml:"adminPassword,omitempty"`
	Persistent    *StorageClassPersistentConfig `yaml:"persistent,omitempty"`
}

type InstallManagerConfigDefaultCD struct {
	Enabled       bool                          `yaml:"enabled,omitempty"`
	Type          string                        `yaml:"type,omitempty"`
	AdminPassword string                        `yaml:"adminPassword,omitempty"`
	Persistent    *StorageClassPersistentConfig `yaml:"persistent,omitempty"`
}

type InstallManagerConfigDefaultRegistry struct {
	Enabled       bool                          `yaml:"enabled,omitempty"`
	Type          string                        `yaml:"type,omitempty"`
	AdminPassword string                        `yaml:"adminPassword,omitempty"`
	Persistent    *StorageClassPersistentConfig `yaml:"persistent,omitempty"`
}

type StorageClassPersistentConfig struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Size    string `yaml:"size,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global: &config.GlobalConfig{},
		Service: &InstallManagerConfig{
			Service: &InstallManagerServiceConfig{},
			DefaultRegistry: &InstallManagerConfigDefaultRegistry{
				Persistent: &StorageClassPersistentConfig{},
			},
			DefaultCD: &InstallManagerConfigDefaultCD{
				Persistent: &StorageClassPersistentConfig{},
			},
			DefaultCI: &InstallManagerConfigDefaultCI{
				Persistent: &StorageClassPersistentConfig{},
			},
		},
	}

	parseFromFile(cfg, config.DefaultConfigFilePath)
	return cfg
}

func parseFromFile(cfg config.Config, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetEnvVars() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("HOME_IDP_ADMIN_PASSWORD", cfg.Global.AdminPassword)
	env.Set("HOME_IDP_STORAGE_CLASS_NAME", cfg.Global.StorageClass)
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
	env.Set("INSTALL_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Service.Service.Port))
	env.Set("DEFAULT_REGISTRY_ENABLED", strconv.FormatBool(cfg.Service.DefaultRegistry.Enabled))
	env.Set("DEFAULT_REGISTRY_STORAGE_CLASS_ENABLED", strconv.FormatBool(cfg.Service.DefaultRegistry.Persistent.Enabled))
	env.Set("DEFAULT_REGISTRY_STORAGE_CLASS_SIZE", cfg.Service.DefaultRegistry.Persistent.Size)
	env.Set("DEFAULT_CD_ENABLED", strconv.FormatBool(cfg.Service.DefaultCD.Enabled))
	env.Set("DEFAULT_CD_ADMIN_PASSWORD", cfg.Service.DefaultCD.AdminPassword)
	env.Set("DEFAULT_CD_STORAGE_CLASS_ENABLED", strconv.FormatBool(cfg.Service.DefaultCD.Persistent.Enabled))
	env.Set("DEFAULT_CD_STORAGE_CLASS_SIZE", cfg.Service.DefaultCD.Persistent.Size)
	env.Set("DEFAULT_CI_ENABLED", strconv.FormatBool(cfg.Service.DefaultCD.Enabled))
	env.Set("DEFAULT_CI_ADMIN_PASSWORD", cfg.Service.DefaultCD.AdminPassword)
	env.Set("DEFAULT_CI_STORAGE_CLASS_ENABLED", strconv.FormatBool(cfg.Service.DefaultCD.Persistent.Enabled))
	env.Set("DEFAULT_CI_STORAGE_CLASS_SIZE", cfg.Service.DefaultCD.Persistent.Size)
}
