package config

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/mail"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/service"
	"gopkg.in/yaml.v2"
)

type RbacManagerConfig struct {
	Name     string
	Enabled  bool                          `yaml:"enabled,omitempty"`
	Service  *service.RbacManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                           `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig         `yaml:"storage,omitempty"`
	Smtp     *mail.SmtpClient              `yaml:"smtp,omitempty"`
}

func New() *RbacManagerConfig {
	cfg := &RbacManagerConfig{Name: "rbac-manager"}

	log.Printf("Start reading home-idp configuration file...")
	parseConfigFile(cfg, config.DefaultConfigFilePath)

	return cfg
}

func parseConfigFile(cfg *RbacManagerConfig, filePath string) error {
	bytes, err := file.ReadFile(filePath)
	if err != nil {
		return err
	}

	tmp := &struct {
		Config *RbacManagerConfig `yaml:"install-manager,omitempty"`
	}{
		Config: cfg,
	}

	if err := yaml.Unmarshal(bytes, tmp); err != nil {
		log.Fatalf("Invalid config file format")
		return err
	}

	return nil
}
