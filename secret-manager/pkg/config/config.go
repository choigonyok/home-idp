package config

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	svc "github.com/choigonyok/home-idp/secret-manager/pkg/service"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
)

type SecretManager struct {
	Config *SecretManagerConfig `yaml:"secret-manager,omitempty"`
}

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
	cfg.SetEnvFromConfig()

	return cfg
}

func (cfg *SecretManagerConfig) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("SECRET_MANAGER_PORT", strconv.Itoa(cfg.Service.Port))
	env.Set("SECRET_MANAGER_STORAGE_TYPE", cfg.Storage.Type)
	env.Set("SECRET_MANAGER_STORAGE_HOST", cfg.Storage.Host)
	env.Set("SECRET_MANAGER_STORAGE_USERNAME", cfg.Storage.Username)
	env.Set("SECRET_MANAGER_STORAGE_PASSWORD", cfg.Storage.Password)
	env.Set("SECRET_MANAGER_STORAGE_DATABASE", cfg.Storage.Database)
	env.Set("SECRET_MANAGER_STORAGE_PORT", strconv.Itoa(cfg.Storage.Port))
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
