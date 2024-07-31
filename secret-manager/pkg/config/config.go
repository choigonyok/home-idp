package config

import (
	"errors"
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
	Enabled  bool                        `yaml:"enabled,omitempty"`
	Service  *svc.SecretManagerSvcConfig `yaml:"service,omitempty"`
	Replicas int                         `yaml:"replicas,omitempty"`
	Storage  *config.StorageConfig       `yaml:"storage,omitempty"`

	KubeConfig *rest.Config
}

func New() *SecretManager {
	return &SecretManager{}
}

func (c *SecretManager) Init(filepath string) error {
	log.Printf("Start reading home-idp configuration file...")
	c.parseManagerConfigFile(filepath)
	c.setEnvFromConfig()
	return nil
}

func (c *SecretManager) setEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("SECRET_MANAGER_PORT", strconv.Itoa(c.Config.Service.Port))
	env.Set("SECRET_MANAGER_STORAGE_TYPE", c.Config.Storage.Type)
	env.Set("SECRET_MANAGER_STORAGE_HOST", c.Config.Storage.Host)
	env.Set("SECRET_MANAGER_STORAGE_USERNAME", c.Config.Storage.Username)
	env.Set("SECRET_MANAGER_STORAGE_PASSWORD", c.Config.Storage.Password)
	env.Set("SECRET_MANAGER_STORAGE_DATABASE", c.Config.Storage.Database)
	env.Set("SECRET_MANAGER_STORAGE_PORT", strconv.Itoa(c.Config.Storage.Port))
}

func (c *SecretManager) parseManagerConfigFile(filepath string) error {
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

// func (c *SecretManager) Get(key string) (any, bool, error) {
// 	fmt.Println("START FINDING", key)
// 	v := reflect.ValueOf(c)

// 	if v.Kind() != reflect.Pointer {
// 		return nil, false, fmt.Errorf("%s", "IS NOT POINTER TYPE")
// 	}

// 	v = v.Elem()

// 	sf, ok := v.Type().FieldByName(key)

// 	if !ok {
// 		return nil, false, nil
// 	}
// 	fmt.Println("ELEM:", v.FieldByName(sf.Name))
// 	return v.FieldByName(sf.Name), true, nil
// }
