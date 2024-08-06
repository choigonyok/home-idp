package config

import (
	"log"

	"github.com/choigonyok/home-idp/deploy-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/file"
	"gopkg.in/yaml.v3"
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
	parseConfigFile(cfg, config.DefaultConfigFilePath)
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *DeployManagerConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
}

func parseConfigFile(cfg *DeployManagerConfig, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	tmp := &struct {
		Config *DeployManagerConfig `yaml:"gateway,omitempty"`
	}{
		Config: cfg,
	}

	if err := yaml.Unmarshal(bytes, tmp); err != nil {
		log.Fatalf("Invalid config file format")
		return err
	}
	return nil
}
