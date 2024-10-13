package model

import (
	"github.com/choigonyok/home-idp/pkg/env"
)

type User struct {
	Name     string
	Roles    []*Role
	Email    string
	Password string
}

func GetDefaultUser() *User {
	u := &User{
		Name:     "admin",
		Password: env.Get("HOME_IDP_ADMIN_PASSWORD"),
		Email:    env.Get("HOME_IDP_GIT_EMAIL"),
		Roles:    []*Role{GetDefaultRole()},
	}

	StoreUser(u)
	return u
}

func StoreUser(u *User) error {
	// STORE TO STORAGE
	return nil
}
