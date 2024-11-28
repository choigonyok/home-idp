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
}

type DeployManagerServiceConfig struct {
	Port int `yaml:"port,omitempty"`
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
		Service: &DeployManagerConfig{
			Service: &DeployManagerServiceConfig{},
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
	env.Set("HOME_IDP_NAMESPACE", cfg.Global.Namespace)
	env.Set("HOME_IDP_HARBOR_HOST", cfg.Global.Harbor.Host)
	env.Set("HOME_IDP_HARBOR_PORT", strconv.Itoa(cfg.Global.Harbor.Port))
	env.Set("HOME_IDP_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_GIT_EMAIL", cfg.Global.Git.Email)
	env.Set("HOME_IDP_GIT_TOKEN", cfg.Global.Git.Token)
	env.Set("HOME_IDP_GIT_REPO", cfg.Global.Git.Repo)
	env.Set("DEPLOY_MANAGER_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
	env.Set("DEPLOY_MANAGER_REGISTRY_PASSWORD", cfg.Global.AdminPassword)
}
