package storage

import (
	"database/sql"
	"time"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type StorageClient interface {
	Close() error
	DB() *sql.DB
}

type Data struct {
	UsersEmail        string
	UsersName         string
	UsersId           int
	UsersPassword     string
	ProjectId         int
	ProjectName       string
	ProjectCreateTime time.Time
	RoleId            int
	RoleName          string
	PolicyId          int
	Policy            string
}

func NewClient(component util.Components) (StorageClient, error) {
	dbType := env.Get(env.GetEnvPrefix(component) + "_MANAGER_STORAGE_TYPE")

	switch dbType {
	case "postgresql":
		client := &PostgresClient{
			Client: newPostgresDB(component),
			Table:  defaultTableName,
		}
		return client, nil
	default:
		return nil, nil
	}
}
