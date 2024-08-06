package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/gateway/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"gopkg.in/yaml.v2"
)

type GatewayConfig struct {
	Name     string
	Enabled  bool                      `yaml:"enabled,omitempty"`
	Service  *service.GatewaySvcConfig `yaml:"service,omitempty"`
	Replicas int                       `yaml:"replicas,omitempty"`
}

func New() *GatewayConfig {
	cfg := &GatewayConfig{
		Name:    "gateway",
		Service: &service.GatewaySvcConfig{},
	}

	log.Printf("Start reading home-idp configuration file...")
	parseConfigFile(cfg, config.DefaultConfigFilePath)

	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *GatewayConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("GATEWAY_SERVICE_PORT", strconv.Itoa(cfg.Service.Port))
	env.Set("GATEWAY_SERVICE_TYPE", cfg.Service.Type)
	env.Set("GATEWAY_ENABLED", strconv.FormatBool(cfg.Enabled))
}

func parseConfigFile(cfg *GatewayConfig, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	tmp := &struct {
		Config *GatewayConfig `yaml:"gateway,omitempty"`
	}{
		Config: cfg,
	}

	if err := yaml.Unmarshal(bytes, tmp); err != nil {
		log.Fatalf("Invalid config file format")
		return err
	}
	return nil
}
