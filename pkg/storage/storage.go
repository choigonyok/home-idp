package storage

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
)

type StorageClient interface {
	Close() error
	Init(config.Components) error
}

func NewClient(component config.Components) (StorageClient, error) {
	switch component {
	case 0: // Secret-Manager
		client := newClient(env.Get("SECRET_MANAGER_STORAGE_TYPE"), component)
		client.Init(component)
		return client, nil
	case 1: // Deploy-Manager
		return nil, nil
	case 2: // Rbac-Manager
		client := newClient(env.Get("RBAC_MANAGER_STORAGE_TYPE"), component)
		client.Init(component)
		return client, nil
	}
	return nil, fmt.Errorf("%s", "INVALID STORAGE TYPE")
}

func newClient(storageType string, component config.Components) StorageClient {
	switch storageType {
	case "postgresql":
		return newPostgreSQLClient(component)
	}
	return nil
}
