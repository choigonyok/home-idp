package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/gateway/pkg/service"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"k8s.io/client-go/rest"
)

type Gateway struct {
	Config *GatewayConfig `yaml:"gateway,omitempty"`
}

type GatewayConfig struct {
	Enabled  bool                      `yaml:"enabled,omitempty"`
	Service  *service.GatewaySvcConfig `yaml:"service,omitempty"`
	Replicas int                       `yaml:"replicas,omitempty"`

	KubeConfig *rest.Config
}

func New() *Gateway {
	return &Gateway{
		Config: &GatewayConfig{
			Service: &service.GatewaySvcConfig{},
		},
	}
}

func (c *Gateway) Init(filepath string) error {
	log.Printf("Start reading home-idp configuration file...")
	file.ParseConfigFile(c, filepath)
	setEnvFromConfig(c.Config)
	return nil
}

func setEnvFromConfig(cfg *GatewayConfig) {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("GATEWAY_SERVICE_PORT", strconv.Itoa(cfg.Service.Port))
	env.Set("GATEWAY_SERVICE_TYPE", cfg.Service.Type)
	env.Set("GATEWAY_ENABLED", strconv.FormatBool(cfg.Enabled))
}
