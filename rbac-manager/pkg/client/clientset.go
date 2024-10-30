package client

import (
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/trace"
	"github.com/choigonyok/home-idp/pkg/util"
)

type RbacManagerClientSet struct {
	StorageClient storage.StorageClient
	TraceClient   *trace.TraceClient
	MailClient    mail.MailClient
}

func EmptyClientSet() *RbacManagerClientSet {
	return &RbacManagerClientSet{}
}

func (cs *RbacManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.StoragePostgresClient:
		tmp := &storage.PostgresClient{}
		tmp.Set(i)
		cs.StorageClient = tmp
		return
	case util.TraceClient:
		tmp := &trace.TraceClient{}
		tmp.Set(i)
		cs.TraceClient = tmp
		return
	default:
		return
	}
}
