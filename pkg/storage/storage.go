package storage

import (
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type StorageClient interface {
	Close() error
	Init(util.Components) error
}

func NewClient(component util.Components) (StorageClient, error) {
	client := newClient(env.Get(env.GetEnvPrefix(component)+"_MANAGER_STORAGE_TYPE"), component)
	client.Init(component)
	return client, nil
}

func newClient(storageType string, component util.Components) StorageClient {
	switch storageType {
	case "postgresql":
		return newPostgreSQLClient(component)
	default:
		return nil
	}
}
