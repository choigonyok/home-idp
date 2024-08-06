package config

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/file"
	svc "github.com/choigonyok/home-idp/secret-manager/pkg/service"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
)

type SecretManagerConfig struct {
	Name     string
	Enabled  bool                        `yaml:"enabled,omitempty"`
	Service  *svc.SecretManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                         `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig       `yaml:"storage,omitempty"`

	KubeConfig *rest.Config
}

func New() *SecretManagerConfig {
	cfg := &SecretManagerConfig{Name: "install-manager"}

	log.Printf("Start reading home-idp configuration file...")
	parseConfigFile(cfg, config.DefaultConfigFilePath)

	return cfg
}

func parseConfigFile(cfg *SecretManagerConfig, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	tmp := &struct {
		Config *SecretManagerConfig `yaml:"secret-manager,omitempty"`
	}{
		Config: cfg,
	}

	if err := yaml.Unmarshal(bytes, tmp); err != nil {
		log.Fatalf("Invalid config file format")
		return err
	}

	return nil
}
