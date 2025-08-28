package service

import (
	"github.com/choigonyok/idp/pkg/model"
	"github.com/choigonyok/idp/pkg/store"
)

type Service struct {
	db *store.SQLiteStore
}

func NewService(db *store.SQLiteStore) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create(svc *model.Service) error {
	if err := s.db.CreateService(svc); err != nil {
		return err
	}
	return nil
}

func (s *Service) List() (*[]model.Service, error) {
	return s.db.ListServices()
}
