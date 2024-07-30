package storage

import (
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

func NewClient(component util.Components) (storage.StorageClient, error) {
	client, _ := storage.NewClient(component)

	if err := initPostgresTables(client.DB()); err != nil {
		return nil, err
	}
	if err := initPostgresFunctions(client.DB()); err != nil {
		return nil, err
	}

	client.SetDeleteQuery(getDeleteQuery())
	client.SetGetQuery(getGetQuery())
	client.SetPutQuery(getPutQuery())
	client.SetListQuery(getListQuery())
	return client, nil
}
