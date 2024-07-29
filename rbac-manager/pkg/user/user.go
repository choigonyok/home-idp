package user

import (
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/role"
)

type User struct {
	Name     string
	Roles    []*role.Role
	Email    string
	Password string
}

func GetDefaultUser() *User {
	u := &User{
		Name:     env.Get("RBAC_MANAGER_ADMIN_USERNAME"),
		Password: env.Get("RBAC_MANAGER_ADMIN_PASSWORD"),
		Email:    env.Get("RBAC_MANAGER_ADMIN_EMAIL"),
		Roles:    []*role.Role{role.GetDefaultRole()},
	}

	Store(u)
	return u
}

func Store(u *User) error {
	// STORE TO STORAGE
	return nil
}
