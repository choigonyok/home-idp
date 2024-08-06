package config

import (
	"log"

	"github.com/choigonyok/home-idp/install-manager/pkg/service"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/file"
	"gopkg.in/yaml.v3"
)

type InstallManagerConfig struct {
	Name     string
	Enabled  bool                             `yaml:"enabled,omitempty"`
	Service  *service.InstallManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                              `yaml:"replicas,omitempty"`
}

func New() *InstallManagerConfig {
	cfg := &InstallManagerConfig{Name: "install-manager"}

	log.Printf("Start reading home-idp configuration file...")
	parseConfigFile(cfg, config.DefaultConfigFilePath)

	return cfg
}

func parseConfigFile(cfg *InstallManagerConfig, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	tmp := &struct {
		Config *InstallManagerConfig `yaml:"install-manager,omitempty"`
	}{
		Config: cfg,
	}

	if err := yaml.Unmarshal(bytes, tmp); err != nil {
		log.Fatalf("Invalid config file format")
		return err
	}

	return nil
}
