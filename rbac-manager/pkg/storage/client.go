package storage

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/storage"
)

type RbacManagerStorageClient struct {
	Client storage.StorageClient
}

func (c *RbacManagerStorageClient) Set(i interface{}) {
	c.Client = parseStorageClientFromInterface(i)
}

func parseStorageClientFromInterface(i interface{}) storage.StorageClient {
	client := i.(storage.StorageClient)
	return client
}

func (c *RbacManagerStorageClient) IsHealthy() bool {
	err := c.Client.DB().Ping()
	if err != nil {
		fmt.Println("TEST POSTGRESQL HEALTHY ERR: ", err)
		return false
	}

	return true
}
