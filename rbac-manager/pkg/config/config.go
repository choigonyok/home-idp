package config

import (
	"errors"
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/mail"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/service"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
)

type RbacManager struct {
	Config *RbacManagerConfig `yaml:"rbac-manager,omitempty"`
}

type RbacManagerConfig struct {
	Enabled  bool                          `yaml:"enabled,omitempty"`
	Service  *service.RbacManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                           `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig         `yaml:"storage,omitempty"`
	Smtp     *mail.SmtpClient              `yaml:"smtp,omitempty"`

	KubeConfig *rest.Config
}

func New() *RbacManager {
	return &RbacManager{}
}

func (c *RbacManager) Init(filepath string) error {
	log.Printf("Start reading home-idp configuration file...")
	c.parseManagerConfigFile(filepath)
	c.setEnvFromConfig()
	return nil
}

func (c *RbacManager) setEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("RBAC_MANAGER_PORT", strconv.Itoa(c.Config.Service.Port))
	env.Set("RBAC_MANAGER_STORAGE_TYPE", c.Config.Storage.Type)
	env.Set("RBAC_MANAGER_STORAGE_HOST", c.Config.Storage.Host)
	env.Set("RBAC_MANAGER_STORAGE_USERNAME", c.Config.Storage.Username)
	env.Set("RBAC_MANAGER_STORAGE_PASSWORD", c.Config.Storage.Password)
	env.Set("RBAC_MANAGER_STORAGE_DATABASE", c.Config.Storage.Database)
	env.Set("RBAC_MANAGER_STORAGE_PORT", strconv.Itoa(c.Config.Storage.Port))
	if c.Config.Smtp.Enabled == true {
		env.Set("RBAC_MANAGER_SMTP_HOST", c.Config.Smtp.Config.Host)
		env.Set("RBAC_MANAGER_SMTP_PORT", c.Config.Smtp.Config.Port)
		env.Set("RBAC_MANAGER_SMTP_USER", c.Config.Smtp.Config.User)
		env.Set("RBAC_MANAGER_SMTP_PASSWORD", c.Config.Smtp.Config.Password)
		env.Set("RBAC_MANAGER_SMTP_DOMAIN", c.Config.Smtp.Config.Domain)
		env.Set("RBAC_MANAGER_SMTP_ENABLED", strconv.FormatBool(c.Config.Smtp.Enabled))
	}
}

func (c *RbacManager) parseManagerConfigFile(filepath string) error {
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
