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
	env.Set("INSTALL_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Service.Service.Port))
	env.Set("DEFAULT_REGISTRY_ENABLED", strconv.FormatBool(cfg.Service.DefaultRegistry.Enabled))
	env.Set("DEFAULT_REGISTRY_ADMIN_PASSWORD", cfg.Service.DefaultRegistry.AdminPassword)
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
	env.Set("GLOBAL_STORAGE_CLASS_NAME", cfg.Global.StorageClass)
	env.Set("GLOBAL_NAMESPACE", cfg.Global.Namespace)
}
