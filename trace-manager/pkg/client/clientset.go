package client

import (
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

type TraceManagerClientSet struct {
	StorageClient storage.StorageClient
}

func EmptyClientSet() *TraceManagerClientSet {
	return &TraceManagerClientSet{}
}

func (cs *TraceManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.StoragePostgresClient:
		tmp := &storage.PostgresClient{}
		tmp.Set(i)
		cs.StorageClient = tmp
		return
	default:
		return
	}
}
