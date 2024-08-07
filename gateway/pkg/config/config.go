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
	Service *GatewayConfig       `yaml:"gateway,omitempty"`
	Global  *config.GlobalConfig `yaml:"global,omitempty"`
}

type GatewayConfig struct {
	Name     string
	Enabled  bool                  `yaml:"enabled,omitempty"`
	Service  *GatewayServiceConfig `yaml:"service,omitempty"`
	Replicas int                   `yaml:"replicas,omitempty"`
}

type GatewayServiceConfig struct {
	Port int    `yaml:"port,omitempty"`
	Type string `yaml:"type,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global:  &config.GlobalConfig{},
		Service: &GatewayConfig{},
	}
	return cfg
}

func (cfg *Config) ParseFromFile(filePath string) error {
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
	env.Set("GATEWAY_SERVICE_PORT", strconv.Itoa(cfg.Service.Service.Port))
	env.Set("GATEWAY_SERVICE_TYPE", cfg.Service.Service.Type)
	env.Set("GATEWAY_ENABLED", strconv.FormatBool(cfg.Service.Enabled))
}
