package role

import "github.com/choigonyok/home-idp/rbac-manager/pkg/policy"

type Role struct {
	Name     string
	Policies []*policy.Policy
}

func GetDefaultRole() *Role {
	r := &Role{
		Name:     "admin",
		Policies: []*policy.Policy{policy.GetDefaultPolicy()},
	}

	Store(r)
	return r
}

func Store(r *Role) error {
	// STORE TO STORAGE
	return nil
}
