package infra

import (
	k8sclient "github.com/choigonyok/idp/pkg/client/k8s"
	"github.com/choigonyok/idp/pkg/model"
	"github.com/choigonyok/idp/pkg/store"
)

type Infra struct {
	db  *store.SQLiteStore
	k8s *k8sclient.K8s
}

func NewInfra(db *store.SQLiteStore, k8s *k8sclient.K8s) *Infra {
	return &Infra{
		db:  db,
		k8s: k8s,
	}
}

func (i *Infra) Create(infra *model.Infra) error {
	switch infra.Platform.Kind {
	case "K8s":
		if err := i.k8s.SetContext(infra.Platform.K8sCluster); err != nil {
			return err
		}
		i.k8s.AddClient(infra.Platform.K8sCluster.Context)
		go i.k8s.StartEventCheck(infra.Platform.K8sCluster.Context)
		v, err := i.k8s.Version(infra.Platform.K8sCluster.Context)
		if err != nil {
			return err
		}
		infra.Version = v
	}
	if err := i.db.CreateInfra(infra); err != nil {
		return err
	}
	return nil
}

func (s *Infra) List() (*[]model.Infra, error) {
	return s.db.ListInfra()
}
