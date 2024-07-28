package storage

import (
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
)

type StorageClient interface {
	Close() error
	Init(config.Components) error
}

func NewClient(component config.Components) (StorageClient, error) {
	client := newClient(env.Get(env.GetEnvPrefix(component)+"_MANAGER_STORAGE_TYPE"), component)
	client.Init(component)
	return client, nil
}

func newClient(storageType string, component config.Components) StorageClient {
	switch storageType {
	case "postgresql":
		return newPostgreSQLClient(component)
	default:
		return nil
	}
}
