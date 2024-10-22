package repository

// import (
// 	"github.com/choigonyok/home-idp/pkg/model"
// 	"github.com/choigonyok/home-idp/pkg/storage"
// )

// type Repository struct {
// 	client storage.StorageClient
// }

// type ProjectsTable struct {
// 	client storage.StorageClient
// }

// type UsersTable struct {
// 	client storage.StorageClient
// }

// type Table interface {
// 	Get(...string) (interface{}, error)
// }

// func New(c storage.StorageClient) *Repository {
// 	return &Repository{
// 		client: c,
// 	}
// }

// func (repo *Repository) Table(table string) Table {
// 	switch table {
// 	case "projects":
// 		return &ProjectsTable{}
// 		// case "users":
// 		// 	return &UsersTable{}
// 	}

// 	return nil
// }

// func (p *ProjectsTable) Get(cols ...string) (interface{}, error) {
// 	r, err := p.client.GetQueryFromProjects(cols...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer r.Close()

// 	projs := []*model.Project{}

// 	for r.Next() {
// 		proj := model.Project{}
// 		r.Scan(&proj.ID, &proj.Name, &proj.Creator)
// 		projs = append(projs, &proj)
// 	}

// 	return projs, nil
// }

// // func (p *UsersTable) Get(cols ...string) (interface{}, error) {
// // 	r, err := p.client.GetQueryFromUsers(cols...)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	projs := []*model.User{}

// // 	for r.Next() {
// // 		usr := model.User{}
// // 		r.Scan(&usr.ID, &usr.Name, &usr.Creator)
// // 		projs = append(projs, &proj)
// // 	}
// // }
