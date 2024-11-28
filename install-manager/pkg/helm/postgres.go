package helm

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type Postgres struct {
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
}

func NewPostgresClient(namespace, releaseName string) *Postgres {
	return &Postgres{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *Postgres) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("bitnami/postgresql:16.0.0", c.Namespace, c.ReleaseName, postgresOverrideValues())
	return nil
}

func (c *Postgres) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

func postgresOverrideValues() map[string]interface{} {
	return map[string]interface{}{
		"global": map[string]interface{}{
			"defaultStorageClass": env.Get("POSTGRES_STORAGECLASS"),
			"volumeName":          "pg",
		},
		"auth": map[string]interface{}{
			"enablePostgresUser": false,
			"username":           env.Get("HOME_IDP_STORAGE_USERNAME"),
			"password":           env.Get("HOME_IDP_STORAGE_PASSWORD"),
			"database":           env.Get("HOME_IDP_STORAGE_DB"),
		},
		"primary": map[string]interface{}{
			"initdb": map[string]interface{}{
				"scriptsConfigMap": "home-idp-postgres-initdb",
				"user":             env.Get("HOME_IDP_STORAGE_USERNAME"),
				"password":         env.Get("HOME_IDP_STORAGE_PASSWORD"),
			},
			"persistence": map[string]interface{}{
				"size": env.Get("POSTGRES_SIZE"),
			},
		},
	}
}
