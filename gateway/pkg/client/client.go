package client

import (
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"google.golang.org/grpc"
)

const (
	rbacManagerPort    = "5105"
	rbacManagerHost    = "localhost"
	installManagerPort = "5107"
	installManagerHost = "localhost"
)

type GatewayClientSet struct {
	GrpcClient    map[util.Components]client.GrpcClient
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
}

type GatewayGrpcClient struct {
	Client *grpc.ClientConn
}

func (gc *GatewayGrpcClient) Close() {
	return
}

func (gc *GatewayGrpcClient) GetConnection() *grpc.ClientConn {
	return gc.Client
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

// ////
// ////
// ////
// ////
// ////
// func New() *GatewayGrpcClient {
// 	client := &GatewayGrpcClient{
// 		InstallManager: globalgrpc.NewClient(installManagerHost, installManagerPort),
// 	}

// 	return client
// }

// func (c *GatewayGrpcClient) GetConnection() *grpc.ClientConn {
// 	return c.InstallManager
// }
// func (c *GatewayGrpcClient) Handler(name string) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
// 		state := hc.getState()
// 		template := hc.responses[state.status]

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(template.statusCode)

// 		w.Write(hc.createRespBody(state, template))
// 	})
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {}
