package config

import (
	"errors"
	"log"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
)

type RbacManager struct {
	Config *RbacManagerConfig `yaml:"rbac-manager,omitempty"`
}

type RbacManagerConfig struct {
	Enabled  bool `yaml:"enabled,omitempty"`
	Replicas int  `yaml:"replicas,omitempty"`

	KubeConfig *rest.Config
}

func New() *RbacManager {
	return &RbacManager{}
}

func (c *RbacManager) Init(filepath string) error {
	log.Printf("Start reading home-idp configuration file...")
	c.parseSecretManagerConfigFile(filepath)
	c.setEnvFromConfig()
	return nil
}

func (c *RbacManager) setEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("KUBECONFIG", "$HOME/.kube/config")
}

func (c *RbacManager) parseSecretManagerConfigFile(filepath string) error {
	if !file.Exist(filepath) {
		log.Fatalf("Cannot find config file in %s", filepath)
		return errors.New("Cannot find config file")

	}
	bytes, _ := file.ReadFile(filepath)

	log.Printf("Start parsing home-idp configuration file...")
	if err := yaml.Unmarshal(bytes, c); err != nil {
		log.Fatalf("Invalid config file format")
		return errors.New("Invalid config file format")
	}
	return nil
}
