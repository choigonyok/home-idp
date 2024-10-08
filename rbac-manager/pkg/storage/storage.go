package storage

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

func NewClient(component util.Components) (storage.StorageClient, error) {
	client, _ := storage.NewClient(component)

	if err := initPostgresTables(client.DB()); err != nil {
		fmt.Println(err)
		return client, err
	}
	if err := initPostgresFunctions(client.DB()); err != nil {
		fmt.Println(err)
		return client, err
	}

	return client, nil
}
