package client

import (
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"google.golang.org/grpc"
)

type GatewayClientSet struct {
	GrpcClient    map[util.Components]client.GrpcClient
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
}

func EmptyClientSet() *GatewayClientSet {
	return &GatewayClientSet{}
}

func (cs *GatewayClientSet) ApplyGrpcClient(conn client.GrpcClient) {
	cs.GrpcClient[util.InstallManager] = conn
}

func (cs *GatewayClientSet) New(conn *grpc.ClientConn) client.GrpcClient {
	return &GatewayGrpcClient{
		Client: conn,
	}
}
