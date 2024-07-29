package storage

import (
	"database/sql"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type StorageClient interface {
	Close() error
	DB() *sql.DB
}

func NewClient(component util.Components) (StorageClient, error) {
	client := &PostgresClient{
		Client: newDB(env.Get(env.GetEnvPrefix(component)+"_MANAGER_STORAGE_TYPE"), component),
		Table:  defaultTableName,
	}
	return client, nil
}

func newDB(storageType string, component util.Components) *sql.DB {
	switch storageType {
	case "postgresql":
		return newPostgresDB(component)
	default:
		return nil
	}
}
