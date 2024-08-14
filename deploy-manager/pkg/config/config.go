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
	Service *DeployManagerConfig `yaml:"deploy-manager,omitempty"`
	Global  *config.GlobalConfig `yaml:"global,omitempty"`
}

type DeployManagerConfig struct {
	Name     string
	Enabled  bool                        `yaml:"enabled,omitempty"`
	Replicas int                         `yaml:"replicas,omitempty"`
	Service  *DeployManagerServiceConfig `yaml:"service,omitempty"`
	Docker   *DeployManagerDockerConfig  `yaml:"docker,omitempty"`
}

type DeployManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
}

type DeployManagerDockerConfig struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// KubeConfig *rest.Config
func New() *Config {
	cfg := &Config{
		Global: &config.GlobalConfig{},
		Service: &DeployManagerConfig{
			Service: &DeployManagerServiceConfig{},
			Docker:  &DeployManagerDockerConfig{},
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
	env.Set("DEPLOY_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Service.Service.Port))
	env.Set("DEPLOY_MANAGER_DOCKER_USERNAME", cfg.Service.Docker.Username)
	env.Set("DEPLOY_MANAGER_DOCKER_PASSWORD", cfg.Service.Docker.Password)
	env.Set("GLOBAL_NAMESPACE", cfg.Global.Namespace)
}
