package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
)

type TraceManagerConfig struct {
	Name     string
	Enabled  bool                       `yaml:"enabled,omitempty"`
	Service  *TraceManagerServiceConfig `yaml:"service,omitempty"`
	Replicas int                        `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig      `yaml:"storage,omitempty"`
}

type TraceManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type Config struct {
	Service *TraceManagerConfig  `yaml:"trace-manager,omitempty"`
	Global  *config.GlobalConfig `yaml:"global,omitempty"`
}

func New() *Config {
	cfg := &Config{
		Global: &config.GlobalConfig{},
		Service: &TraceManagerConfig{
			Service: &TraceManagerServiceConfig{},
		},
	}

	config.ParseFromFile(cfg, config.DefaultConfigFilePath)
	return cfg
}

func (cfg *Config) SetEnvVars() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("HOME_IDP_ADMIN_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_ADMIN_PASSWORD", cfg.Global.AdminPassword)
	env.Set("HOME_IDP_STORAGE_USERNAME", cfg.Global.Storage.Username)
	env.Set("HOME_IDP_STORAGE_PASSWORD", cfg.Global.Storage.Password)
	env.Set("HOME_IDP_STORAGE_DATABASE", cfg.Global.Storage.Database)
	env.Set("TRACE_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
}
