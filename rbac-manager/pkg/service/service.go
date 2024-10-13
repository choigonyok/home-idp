package service

import (
	"fmt"
	"time"

	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/client"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
)

type RbacManager struct {
	ClientSet *client.RbacManagerClientSet
	Server    *grpc.RbacManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *RbacManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &RbacManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *RbacManager) Stop() {
	if err := svc.Server.Stop(); err != nil {
		return
	}
	return
}

func (svc *RbacManager) Start() {
	for {
		if svc.ClientSet.StorageClient.IsHealthy() {
			break
		}
		fmt.Println("RBAC MANAGER POSTGRESQL DATABASE IS NOT READY YET!")
		time.Sleep(time.Second * 1)
	}
	svc.Server.Run()
}
