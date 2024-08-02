package config

import (
	"log"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/file"
	"k8s.io/client-go/rest"
)

type DeployManagerConfig struct {
	Name     string
	Enabled  bool                            `yaml:"enabled,omitempty"`
	Replicas int                             `yaml:"replicas,omitempty"`
	Service  *service.DeployManagerSvcConfig `yaml:"service,omitempty"`

	KubeConfig *rest.Config
}

func New() *DeployManagerConfig {
	cfg := &DeployManagerConfig{Name: "deploy-manager"}

	log.Printf("Start reading home-idp configuration file...")
	file.ParseConfigFile(cfg, config.DefaultConfigFilePath)
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *DeployManagerConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
}

func (cfg *DeployManagerConfig) GetName() string {
	return cfg.Name
}
