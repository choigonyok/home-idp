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
	Enabled  bool                         `yaml:"enabled,omitempty"`
	Service  *InstallManagerServiceConfig `yaml:"service,omitempty"`
	Replicas int                          `yaml:"replicas,omitempty"`
}

type InstallManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global:  &config.GlobalConfig{},
		Service: &InstallManagerConfig{},
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

func (cfg *Config) GetPort() int {
	return cfg.Service.Service.Port
}

func (cfg *Config) GetNamespace() string {
	return cfg.Global.Namespace
}

func (cfg *Config) SetEnvVars() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("INSTALL_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Service.Service.Port))
	env.Set("INSTALL_MANAGER_ENABLED", strconv.FormatBool(cfg.Service.Enabled))
}
