package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/gateway/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
)

type Gateway struct {
	Config *GatewayConfig `yaml:"gateway,omitempty"`
}

type GatewayConfig struct {
	Name     string
	Enabled  bool                      `yaml:"enabled,omitempty"`
	Service  *service.GatewaySvcConfig `yaml:"service,omitempty"`
	Replicas int                       `yaml:"replicas,omitempty"`
}

func New() *GatewayConfig {
	cfg := &GatewayConfig{
		Name: "gateway",
	}

	log.Printf("Start reading home-idp configuration file...")
	file.ParseConfigFile(cfg, config.DefaultConfigFilePath)
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *GatewayConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("GATEWAY_SERVICE_PORT", strconv.Itoa(cfg.Service.Port))
	env.Set("GATEWAY_SERVICE_TYPE", cfg.Service.Type)
	env.Set("GATEWAY_ENABLED", strconv.FormatBool(cfg.Enabled))
}

func (cfg *GatewayConfig) GetName() string {
	return cfg.Name
}
