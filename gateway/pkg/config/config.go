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
		Global: &config.GlobalConfig{
			Ingress: &config.GlobalConfigIngress{},
			Git: &config.GlobalConfigGit{
				Oauth: &config.GlobalConfigGitOauth{},
			},
			Harbor: &config.GlobalConfigHarbor{},
			UI:     &config.GlobalConfigUI{},
		},
		Service: &GatewayConfig{
			Service: &GatewayServiceConfig{},
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
	env.Set("HOME_IDP_API_HOST", cfg.Global.Ingress.Host)
	env.Set("HOME_IDP_API_PORT", strconv.Itoa(cfg.Global.Ingress.Port))
	env.Set("HOME_IDP_API_TLS_ENABLED", strconv.FormatBool(cfg.Global.Ingress.TLS))
	env.Set("HOME_IDP_HARBOR_HOST", cfg.Global.Harbor.Host)
	env.Set("HOME_IDP_HARBOR_PORT", strconv.Itoa(cfg.Global.Harbor.Port))
	env.Set("HOME_IDP_HARBOR_TLS_ENABLED", strconv.FormatBool(cfg.Global.Harbor.TLS))
	env.Set("HOME_IDP_GIT_USERNAME", cfg.Global.Git.Username)
	env.Set("HOME_IDP_GIT_EMAIL", cfg.Global.Git.Email)
	env.Set("HOME_IDP_GIT_REPO", cfg.Global.Git.Repo)

	env.Set("HOME_IDP_GIT_OAUTH_CLIENT_ID", cfg.Global.Git.Oauth.ClientID)
	env.Set("HOME_IDP_GIT_OAUTH_CLIENT_SECRET", cfg.Global.Git.Oauth.ClientSecret)
	env.Set("HOME_IDP_ADMIN_PASSWORD", cfg.Global.AdminPassword)
	env.Set("HOME_IDP_NAMESPACE", cfg.Global.Namespace)
	env.Set("HOME_IDP_UI_HOST", cfg.Global.UI.Host)
	env.Set("HOME_IDP_UI_TLS_ENABLED", strconv.FormatBool(cfg.Global.UI.TLS))
	env.Set("HOME_IDP_UI_PORT", strconv.Itoa(cfg.Global.UI.Port))
	env.Set("HOME_IDP_GIT_TOKEN", cfg.Global.Git.Token)
	env.Set("GATEWAY_SERVICE_PORT", strconv.Itoa(cfg.Global.Port))
	env.Set("GATEWAY_SERVICE_TYPE", cfg.Service.Service.Type)
	env.Set("GATEWAY_ENABLED", strconv.FormatBool(cfg.Service.Enabled))
}
